package operate

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"

	humanize "github.com/dustin/go-humanize"
	"github.com/itchio/butler/buse"
	"github.com/itchio/butler/installer/bfs"
	"github.com/itchio/wharf/eos"

	"github.com/itchio/butler/installer"

	"github.com/go-errors/errors"
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

	verb := ""
	switch params.Fresh {
	case false:
		verb = "performing re-install "
	default:
		verb = "performing fresh install "
	}

	consumer.Infof("%s for %s", verb, gameToString(params.Game))
	consumer.Infof("...into directory %s", params.InstallFolder)
	consumer.Infof("...using stage directory %s", oc.StageFolder())

	client, err := clientFromCredentials(params.Credentials)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	if params.Upload == nil {
		consumer.Infof("No upload specified, looking for compatible ones...")
		uploadsFilterResult, err := getFilteredUploads(client, params.Game, params.Credentials, consumer)
		if err != nil {
			return nil, errors.Wrap(err, 0)
		}

		if len(uploadsFilterResult.Uploads) == 0 {
			consumer.Errorf("Didn't find a compatible upload.")
			consumer.Errorf("The initial uploads were:", len(uploadsFilterResult.InitialUploads))
			for _, upload := range uploadsFilterResult.InitialUploads {
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

	var archiveURLPath string
	if params.Build == nil {
		archiveURLPath = fmt.Sprintf("/upload/%d/download", params.Upload.ID)
	} else {
		archiveURLPath = fmt.Sprintf("/upload/%d/download/builds/%d/archive", params.Upload.ID, params.Build.ID)
	}
	values := make(url.Values)
	values.Set("api_key", params.Credentials.APIKey)
	if params.Credentials.DownloadKey != 0 {
		values.Set("download_key_id", fmt.Sprintf("%d", params.Credentials.DownloadKey))
	}
	var archiveURL = fmt.Sprintf("itchfs://%s?%s", archiveURLPath, values.Encode())

	// TODO: support http servers that don't have range request
	// (just copy it first). see DownloadInstallSource later on.
	file, err := eos.Open(archiveURL)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	istate := &InstallSubcontextState{}

	isub := &InstallSubcontext{
		data: istate,
	}

	oc.Load(isub)

	if istate.InstallerInfo == nil {
		consumer.Infof("Probing %s (%s)", stats.Name(), humanize.IBytes(uint64(stats.Size())))

		installerInfo, err := installer.GetInstallerInfo(consumer, file)
		if err != nil {
			return nil, errors.Wrap(err, 0)
		}

		// sniffing may have read parts of the file, so seek back to beginning
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return nil, errors.Wrap(err, 0)
		}

		istate.InstallerInfo = installerInfo
		oc.Save(isub)
	} else {
		consumer.Infof("Using cached installer info")
	}

	installerInfo := istate.InstallerInfo
	consumer.Infof("Will use installer %s", installerInfo.Type)
	manager := installer.GetManager(string(installerInfo.Type))
	if manager == nil {
		msg := fmt.Sprintf("No manager for installer %s", installerInfo.Type)
		return nil, errors.New(msg)
	}

	receiptIn, err := bfs.ReadReceipt(params.InstallFolder)
	if err != nil {
		receiptIn = nil
		consumer.Errorf("Could not read existing receipt: %s", err.Error())
	}

	managerInstallParams := &installer.InstallParams{
		Consumer: consumer,
		Fresh:    params.Fresh,

		File:              file,
		InstallerInfo:     istate.InstallerInfo,
		StageFolderPath:   oc.StageFolder(),
		InstallFolderPath: params.InstallFolder,

		ReceiptIn: receiptIn,

		Context: oc.ctx,
	}

	tryInstall := func() (*installer.InstallResult, error) {
		defer managerInstallParams.File.Close()

		select {
		case <-oc.ctx.Done():
			return nil, ErrCancelled
		default:
			// keep going!
		}

		err = oc.conn.Notify(oc.ctx, "TaskStarted", &buse.TaskStartedNotification{
			Reason:    buse.TaskReasonInstall,
			Type:      buse.TaskTypeInstall,
			Game:      params.Game,
			Upload:    params.Upload,
			Build:     params.Build,
			TotalSize: stats.Size(),
		})
		if err != nil {
			return nil, errors.Wrap(err, 0)
		}

		oc.StartProgress()
		res, installErr := manager.Install(managerInstallParams)
		oc.EndProgress()

		err = oc.conn.Notify(oc.ctx, "TaskEnded", &buse.TaskEndedNotification{})
		if err != nil {
			return nil, errors.Wrap(err, 0)
		}

		if installErr != nil {
			return nil, errors.Wrap(installErr, 0)
		}
		return res, nil
	}

	res, err := tryInstall()
	if err != nil && errors.Is(err, installer.ErrNeedLocal) {
		consumer.Infof("Install source needs to be available locally, copying to disk...")

		dlErr := func() error {
			err = oc.conn.Notify(oc.ctx, "TaskStarted", &buse.TaskStartedNotification{
				Reason:    buse.TaskReasonInstall,
				Type:      buse.TaskTypeDownload,
				Game:      params.Game,
				Upload:    params.Upload,
				Build:     params.Build,
				TotalSize: stats.Size(),
			})
			if err != nil {
				return errors.Wrap(err, 0)
			}

			// TODO: add missing TaskStarted/TaskEnded notifications
			// TODO: if download is already done, don't redo it (save that in context?)
			destName := filepath.Base(stats.Name())
			destPath := filepath.Join(oc.StageFolder(), "install-source", destName)

			oc.StartProgress()
			err := DownloadInstallSource(oc, file, destPath)
			oc.EndProgress()
			oc.consumer.Progress(0)
			if err != nil {
				return errors.Wrap(err, 0)
			}

			err = oc.conn.Notify(oc.ctx, "TaskEnded", &buse.TaskEndedNotification{})
			if err != nil {
				return errors.Wrap(err, 0)
			}

			// fixme: this must be closed
			lf, err := os.Open(destPath)
			if err != nil {
				return errors.Wrap(err, 0)
			}

			managerInstallParams.File = lf
			return nil
		}()

		if dlErr != nil {
			return nil, errors.Wrap(dlErr, 0)
		}

		consumer.Infof("Re-invoking manager with local file...")
		res, err = tryInstall()
	}

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	consumer.Infof("Install successful, writing receipt")
	receipt := &bfs.Receipt{
		InstallerName: string(installerInfo.Type),
		Game:          params.Game,
		Upload:        params.Upload,
		Build:         params.Build,

		Files: res.Files,

		// optionals:
		MSIProductID: res.MSIProductCode,
	}

	err = receipt.WriteReceipt(params.InstallFolder)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return res, nil
}

type InstallSubcontextState struct {
	InstallerInfo *installer.InstallerInfo `json:"installerInfo,omitempty"`
}

type InstallSubcontext struct {
	data *InstallSubcontextState
}

var _ Subcontext = (*InstallSubcontext)(nil)

func (mt *InstallSubcontext) Key() string {
	return "install"
}

func (mt *InstallSubcontext) Data() interface{} {
	return &mt.data
}
