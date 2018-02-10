// +build windows

package runner

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-errors/errors"
	"github.com/itchio/butler/cmd/winsandbox"
	"github.com/itchio/butler/runner/execas"
	"github.com/itchio/butler/runner/syscallex"
	"github.com/itchio/butler/runner/winutil"
	"github.com/itchio/wharf/state"
)

type winsandboxRunner struct {
	params *RunnerParams

	playerData *winsandbox.PlayerData
}

var _ Runner = (*winsandboxRunner)(nil)

func newWinSandboxRunner(params *RunnerParams) (Runner, error) {
	wr := &winsandboxRunner{
		params: params,
	}
	return wr, nil
}

func (wr *winsandboxRunner) Prepare() error {
	// TODO: create user if it doesn't exist
	consumer := wr.params.Consumer

	nullConsumer := &state.Consumer{}
	err := winsandbox.Check(nullConsumer)
	if err != nil {
		consumer.Warnf("Sandbox isn't setup properly: %s", err.Error())
		return errors.New("TODO: ask user for permission to set up")
	}

	playerData, err := winsandbox.GetPlayerData()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	wr.playerData = playerData

	consumer.Infof("Successfully retrieved login details for sandbox user")
	return nil
}

func (wr *winsandboxRunner) Run() error {
	var err error
	params := wr.params
	consumer := params.Consumer
	pd := wr.playerData

	consumer.Infof("Running as user (%s)", pd.Username)

	env := params.Env
	setEnv := func(key string, value string) {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	setEnv("username", pd.Username)
	// we're not setting `userdomain` or `userdomain_roaming_profile`,
	// since we expect those to be the same for the regular user
	// and the sandbox user

	// TODO: check, and trigger setup if needed

	err = winutil.Impersonate(pd.Username, ".", pd.Password, func() error {
		profileDir, err := winutil.GetFolderPath(winutil.FolderTypeProfile)
		if err != nil {
			return errors.Wrap(err, 0)
		}
		// environment variables are case-insensitive on windows,
		// and exec{,as}.Command do case-insensitive deduplication properly
		setEnv("userprofile", profileDir)

		// when %userprofile% is `C:\Users\terry`,
		// %homepath% is usually `\Users\terry`.
		homePath := strings.TrimPrefix(profileDir, filepath.VolumeName(profileDir))
		setEnv("homepath", homePath)

		appDataDir, err := winutil.GetFolderPath(winutil.FolderTypeAppData)
		if err != nil {
			return errors.Wrap(err, 0)
		}
		setEnv("appdata", appDataDir)

		localAppDataDir, err := winutil.GetFolderPath(winutil.FolderTypeLocalAppData)
		if err != nil {
			return errors.Wrap(err, 0)
		}
		setEnv("localappdata", localAppDataDir)

		return nil
	})

	sp := &winutil.SharingPolicy{
		Trustee: pd.Username,
	}
	sp.Entries = append(sp.Entries, &winutil.ShareEntry{
		Path:        params.InstallFolder,
		Inheritance: winutil.InheritanceModeFull,
		Rights:      winutil.RightsFull,
	})
	consumer.Infof("Sharing policy: %s", sp)

	err = sp.Grant(consumer)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	defer sp.Revoke(consumer)

	err = SetupJobObject(consumer)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	cmd := execas.CommandContext(params.Ctx, params.FullTargetPath, params.Args...)
	cmd.Username = pd.Username
	cmd.Domain = "."
	cmd.Password = pd.Password
	cmd.Dir = params.Dir
	cmd.Env = env
	cmd.Stdout = params.Stdout
	cmd.Stderr = params.Stderr
	cmd.SysProcAttr = &syscallex.SysProcAttr{
		LogonFlags: syscallex.LOGON_WITH_PROFILE,
	}

	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	err = WaitJobObject(consumer)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}
