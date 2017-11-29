package operate

import (
	"fmt"
	"io"
	"net/url"

	humanize "github.com/dustin/go-humanize"
	"github.com/itchio/butler/buse"
	"github.com/itchio/butler/installer/bfs"
	"github.com/itchio/wharf/eos"

	"github.com/itchio/butler/installer"

	"github.com/go-errors/errors"
	"github.com/itchio/butler/manager"
	itchio "github.com/itchio/go-itchio"
)

func install(oc *OperationContext, meta *MetaSubcontext) (*installer.InstallResult, error) {
	consumer := oc.Consumer()

	params := meta.data.InstallParams

	if params == nil {
		return nil, errors.New("Missing install params")
	}

	if params.Game == nil {
		return nil, errors.New("Missing game in install")
	}

	if params.InstallFolder == "" {
		return nil, errors.New("Missing install folder in install")
	}

	consumer.Infof("Installing game %s", params.Game.Title)
	consumer.Infof("...into directory %s", params.InstallFolder)
	consumer.Infof("...using stage directory %s", oc.StageFolder())

	client, err := clientFromCredentials(params.Credentials)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	if params.Upload == nil {
		consumer.Infof("No upload specified, looking for compatible ones...")
		uploads, err := client.ListGameUploads(&itchio.ListGameUploadsParams{
			GameID:        params.Game.ID,
			DownloadKeyID: params.Credentials.DownloadKey,
		})
		if err != nil {
			return nil, errors.Wrap(err, 0)
		}

		consumer.Infof("Filtering %d uploads", len(uploads.Uploads))

		uploadsFilterResult := manager.NarrowDownUploads(uploads.Uploads, params.Game, manager.CurrentRuntime())
		consumer.Infof("After filter, got %d uploads, they are: ", len(uploadsFilterResult.Uploads))
		for _, upload := range uploadsFilterResult.Uploads {
			consumer.Infof("- %#v", upload)
		}

		if len(uploadsFilterResult.Uploads) == 0 {
			consumer.Warnf("Didn't find a compatible upload. The initial uploads were:", len(uploads.Uploads))
			for _, upload := range uploads.Uploads {
				consumer.Infof("- %#v", upload)
			}

			return nil, (&OperationError{
				Code:      "noCompatibleUploads",
				Message:   "No compatible uploads",
				Operation: "install",
			}).Throw()
		}

		if len(uploadsFilterResult.Uploads) == 1 {
			params.Upload = uploadsFilterResult.Uploads[0]
		} else {
			var r buse.PickUploadResult
			err := oc.conn.Call(oc.ctx, "PickUpload", &buse.PickUploadParams{
				Uploads: uploadsFilterResult.Uploads,
			}, &r)
			if err != nil {
				return nil, errors.Wrap(err, 0)
			}

			params.Upload = uploadsFilterResult.Uploads[r.Index]
		}

		if params.Upload.Build != nil {
			// if we reach this point, we *just now* queried for an upload,
			// so we know the build object is the latest
			params.Build = params.Upload.Build
		}

		oc.Save(meta)
	}

	// TODO: if upload is wharf-enabled, retrieve build & include it in context/receipt/result etc.

	var archiveUrlPath string
	if params.Build == nil {
		archiveUrlPath = fmt.Sprintf("/upload/%d/download", params.Upload.ID)
	} else {
		archiveUrlPath = fmt.Sprintf("/upload/%d/download/builds/%d/archive", params.Upload.ID, params.Build.ID)
	}
	values := make(url.Values)
	values.Set("api_key", params.Credentials.APIKey)
	if params.Credentials.DownloadKey != 0 {
		values.Set("download_key_id", fmt.Sprintf("%d", params.Credentials.DownloadKey))
	}
	var archiveUrl = fmt.Sprintf("itchfs://%s?%s", archiveUrlPath, values.Encode())

	// TODO: support http servers that don't have range request
	// (just copy it first)
	file, err := eos.Open(archiveUrl)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	consumer.Infof("Probing %s (%s)", stats.Name(), humanize.IBytes(uint64(stats.Size())))

	installerInfo, err := getInstallerInfo(consumer, file)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	// TODO: cache get installer info result in context
	consumer.Infof("Will use installer %s", installerInfo.Type)
	manager := installer.GetManager(string(installerInfo.Type))
	if manager == nil {
		msg := fmt.Sprintf("No manager for installer %s", installerInfo.Type)
		return nil, errors.New(msg)
	}

	receiptIn, err := bfs.ReadReceipt(params.InstallFolder)
	if err != nil {
		receiptIn = nil
		consumer.Warnf("Could not read existing receipt: %s", err.Error())
	}

	// sniffing may have read parts of the file, so seek back to beginning
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	err = oc.conn.Notify(oc.ctx, "TaskStarted", &buse.TaskStartedNotification{
		Reason: buse.TaskReasonInstall,
		Type:   buse.TaskTypeInstall,
		Game:   params.Game,
		Upload: params.Upload,
		Build:  params.Build,
	})
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	oc.StartProgressWithTotalBytes(stats.Size())
	res, err := manager.Install(&installer.InstallParams{
		Consumer: consumer,

		File:              file,
		StageFolderPath:   oc.StageFolder(),
		InstallFolderPath: params.InstallFolder,

		ReceiptIn: receiptIn,
	})
	oc.EndProgress()
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	err = oc.conn.Notify(oc.ctx, "TaskEnded", &buse.TaskEndedNotification{})
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	consumer.Infof("Install successful, writing receipt")
	receipt := &bfs.Receipt{
		Files:         res.Files,
		InstallerName: string(installerInfo.Type),
		Game:          params.Game,
		Upload:        params.Upload,
		Build:         params.Build,
	}

	err = receipt.WriteReceipt(params.InstallFolder)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return res, nil
}
