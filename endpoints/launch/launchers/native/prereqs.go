package native

import (
	"fmt"
	"strings"

	"github.com/itchio/butler/buse/messages"

	"github.com/go-errors/errors"
	"github.com/itchio/butler/buse"
	"github.com/itchio/butler/cmd/prereqs"
	"github.com/itchio/butler/endpoints/launch"
)

func handlePrereqs(params *launch.LauncherParams) error {
	consumer := params.RequestContext.Consumer

	var listed []string

	// add manifest prereqs
	if params.AppManifest == nil {
		consumer.Infof("No manifest, no prereqs")
	} else {
		if len(params.AppManifest.Prereqs) == 0 {
			consumer.Infof("Got manifest but no prereqs requested")
		} else {
			for _, p := range params.AppManifest.Prereqs {
				listed = append(listed, p.Name)
			}
		}
	}

	// append built-in params if we need some
	runtime := params.Runtime
	if runtime.Platform == buse.ItchPlatformLinux && params.Sandbox {
		firejailName := fmt.Sprintf("firejail-%s", runtime.Arch())
		listed = append(listed, firejailName)
	}

	if params.PrereqsDir == "" {
		return errors.New("PrereqsDir cannot be empty")
	}

	pc := &prereqs.PrereqsContext{
		RequestContext: params.RequestContext,
		Credentials:    params.Credentials,
		Runtime:        params.Runtime,
		Consumer:       params.RequestContext.Consumer,
		PrereqsDir:     params.PrereqsDir,
	}

	var pending []string
	for _, name := range listed {
		if pc.HasInstallMarker(name) {
			continue
		}

		pending = append(pending, name)
	}

	var err error
	pending, err = pc.FilterPrereqs(pending)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	if len(pending) == 0 {
		consumer.Infof("✓ %d Prereqs already installed or irrelevant: %s", len(listed), strings.Join(listed, ", "))
		return nil
	}

	pa, err := pc.AssessPrereqs(pending)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	if len(pa.Done) > 0 {
		consumer.Infof("✓ %d Prereqs already done: %s", len(pa.Done), strings.Join(pa.Done, ", "))
	}

	if len(pa.Todo) == 0 {
		consumer.Infof("Everything done!")
		return nil
	}
	consumer.Infof("→ %d Prereqs to install: %s", len(pa.Todo), strings.Join(pa.Todo, ", "))

	{
		psn := &buse.PrereqsStartedNotification{
			Tasks: make(map[string]*buse.PrereqTask),
		}
		for i, name := range pa.Todo {
			entry, err := pc.GetEntry(name)
			if err != nil {
				return errors.Wrap(err, 0)
			}

			psn.Tasks[name] = &buse.PrereqTask{
				FullName: entry.FullName,
				Order:    i,
			}
		}

		err = messages.PrereqsStarted.Notify(params.RequestContext, psn)
		if err != nil {
			consumer.Warnf(err.Error())
		}
	}

	tsc := &prereqs.TaskStateConsumer{
		OnState: func(state *buse.PrereqsTaskStateNotification) {
			err = messages.PrereqsTaskState.Notify(params.RequestContext, state)
			if err != nil {
				consumer.Warnf(err.Error())
			}
		},
	}

	err = pc.FetchPrereqs(tsc, pa.Todo)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	plan, err := pc.BuildPlan(pa.Todo)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	err = pc.InstallPrereqs(tsc, plan)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	for _, name := range pa.Todo {
		err = pc.MarkInstalled(name)
		if err != nil {
			return errors.Wrap(err, 0)
		}
	}

	err = messages.PrereqsEnded.Notify(params.RequestContext, &buse.PrereqsEndedNotification{})
	if err != nil {
		consumer.Warnf(err.Error())
	}

	return nil
}
