// Code generated by busegen; DO NOT EDIT.

package messages

import (
	"encoding/json"
	"errors"

	"github.com/itchio/butler/buse"
	"github.com/sourcegraph/jsonrpc2"
)


//==============================
// Utilities
//==============================

// Version.Get (Request)

type VersionGetType struct {}

var _ RequestMessage = (*VersionGetType)(nil)

func (r *VersionGetType) Method() string {
  return "Version.Get"
}

func (r *VersionGetType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.VersionGetParams) (*buse.VersionGetResult, error)) {
  router.Register("Version.Get", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.VersionGetParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Version.Get")
    }
    return res, nil
  })
}

var VersionGet *VersionGetType


//==============================
// Miscellaneous
//==============================

// Log (Notification)

type LogType struct {}

var _ NotificationMessage = (*LogType)(nil)

func (r *LogType) Method() string {
  return "Log"
}

func (r *LogType) Notify(rc *buse.RequestContext, params *buse.LogNotification) (error) {
  return rc.Notify("Log", params)
}

var Log *LogType


//==============================
// Profile
//==============================

// Profile.List (Request)

type ProfileListType struct {}

var _ RequestMessage = (*ProfileListType)(nil)

func (r *ProfileListType) Method() string {
  return "Profile.List"
}

func (r *ProfileListType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.ProfileListParams) (*buse.ProfileListResult, error)) {
  router.Register("Profile.List", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.ProfileListParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Profile.List")
    }
    return res, nil
  })
}

var ProfileList *ProfileListType

// Profile.LoginWithPassword (Request)

type ProfileLoginWithPasswordType struct {}

var _ RequestMessage = (*ProfileLoginWithPasswordType)(nil)

func (r *ProfileLoginWithPasswordType) Method() string {
  return "Profile.LoginWithPassword"
}

func (r *ProfileLoginWithPasswordType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.ProfileLoginWithPasswordParams) (*buse.ProfileLoginWithPasswordResult, error)) {
  router.Register("Profile.LoginWithPassword", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.ProfileLoginWithPasswordParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Profile.LoginWithPassword")
    }
    return res, nil
  })
}

var ProfileLoginWithPassword *ProfileLoginWithPasswordType

// Profile.RequestCaptcha (Request)

type ProfileRequestCaptchaType struct {}

var _ RequestMessage = (*ProfileRequestCaptchaType)(nil)

func (r *ProfileRequestCaptchaType) Method() string {
  return "Profile.RequestCaptcha"
}

func (r *ProfileRequestCaptchaType) Call(rc *buse.RequestContext, params *buse.ProfileRequestCaptchaParams) (*buse.ProfileRequestCaptchaResult, error) {
  var result buse.ProfileRequestCaptchaResult
  err := rc.Call("Profile.RequestCaptcha", params, &result)
  return &result, err
}

var ProfileRequestCaptcha *ProfileRequestCaptchaType

// Profile.RequestTOTP (Request)

type ProfileRequestTOTPType struct {}

var _ RequestMessage = (*ProfileRequestTOTPType)(nil)

func (r *ProfileRequestTOTPType) Method() string {
  return "Profile.RequestTOTP"
}

func (r *ProfileRequestTOTPType) Call(rc *buse.RequestContext, params *buse.ProfileRequestTOTPParams) (*buse.ProfileRequestTOTPResult, error) {
  var result buse.ProfileRequestTOTPResult
  err := rc.Call("Profile.RequestTOTP", params, &result)
  return &result, err
}

var ProfileRequestTOTP *ProfileRequestTOTPType

// Profile.UseSavedLogin (Request)

type ProfileUseSavedLoginType struct {}

var _ RequestMessage = (*ProfileUseSavedLoginType)(nil)

func (r *ProfileUseSavedLoginType) Method() string {
  return "Profile.UseSavedLogin"
}

func (r *ProfileUseSavedLoginType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.ProfileUseSavedLoginParams) (*buse.ProfileUseSavedLoginResult, error)) {
  router.Register("Profile.UseSavedLogin", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.ProfileUseSavedLoginParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Profile.UseSavedLogin")
    }
    return res, nil
  })
}

var ProfileUseSavedLogin *ProfileUseSavedLoginType

// Profile.Forget (Request)

type ProfileForgetType struct {}

var _ RequestMessage = (*ProfileForgetType)(nil)

func (r *ProfileForgetType) Method() string {
  return "Profile.Forget"
}

func (r *ProfileForgetType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.ProfileForgetParams) (*buse.ProfileForgetResult, error)) {
  router.Register("Profile.Forget", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.ProfileForgetParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Profile.Forget")
    }
    return res, nil
  })
}

var ProfileForget *ProfileForgetType


//==============================
// Fetch
//==============================

// Fetch.Game (Request)

type FetchGameType struct {}

var _ RequestMessage = (*FetchGameType)(nil)

func (r *FetchGameType) Method() string {
  return "Fetch.Game"
}

func (r *FetchGameType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.FetchGameParams) (*buse.FetchGameResult, error)) {
  router.Register("Fetch.Game", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.FetchGameParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Fetch.Game")
    }
    return res, nil
  })
}

var FetchGame *FetchGameType

// Fetch.Game.Yield (Notification)

type FetchGameYieldType struct {}

var _ NotificationMessage = (*FetchGameYieldType)(nil)

func (r *FetchGameYieldType) Method() string {
  return "Fetch.Game.Yield"
}

func (r *FetchGameYieldType) Notify(rc *buse.RequestContext, params *buse.FetchGameYieldNotification) (error) {
  return rc.Notify("Fetch.Game.Yield", params)
}

var FetchGameYield *FetchGameYieldType

// Fetch.Collection (Request)

type FetchCollectionType struct {}

var _ RequestMessage = (*FetchCollectionType)(nil)

func (r *FetchCollectionType) Method() string {
  return "Fetch.Collection"
}

func (r *FetchCollectionType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.FetchCollectionParams) (*buse.FetchCollectionResult, error)) {
  router.Register("Fetch.Collection", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.FetchCollectionParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Fetch.Collection")
    }
    return res, nil
  })
}

var FetchCollection *FetchCollectionType

// Fetch.Collection.Yield (Notification)

type FetchCollectionYieldType struct {}

var _ NotificationMessage = (*FetchCollectionYieldType)(nil)

func (r *FetchCollectionYieldType) Method() string {
  return "Fetch.Collection.Yield"
}

func (r *FetchCollectionYieldType) Notify(rc *buse.RequestContext, params *buse.FetchCollectionYieldNotification) (error) {
  return rc.Notify("Fetch.Collection.Yield", params)
}

var FetchCollectionYield *FetchCollectionYieldType

// Fetch.ProfileCollections (Request)

type FetchProfileCollectionsType struct {}

var _ RequestMessage = (*FetchProfileCollectionsType)(nil)

func (r *FetchProfileCollectionsType) Method() string {
  return "Fetch.ProfileCollections"
}

func (r *FetchProfileCollectionsType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.FetchProfileCollectionsParams) (*buse.FetchProfileCollectionsResult, error)) {
  router.Register("Fetch.ProfileCollections", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.FetchProfileCollectionsParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Fetch.ProfileCollections")
    }
    return res, nil
  })
}

var FetchProfileCollections *FetchProfileCollectionsType

// Fetch.ProfileCollections.Yield (Notification)

type FetchProfileCollectionsYieldType struct {}

var _ NotificationMessage = (*FetchProfileCollectionsYieldType)(nil)

func (r *FetchProfileCollectionsYieldType) Method() string {
  return "Fetch.ProfileCollections.Yield"
}

func (r *FetchProfileCollectionsYieldType) Notify(rc *buse.RequestContext, params *buse.FetchProfileCollectionsYieldNotification) (error) {
  return rc.Notify("Fetch.ProfileCollections.Yield", params)
}

var FetchProfileCollectionsYield *FetchProfileCollectionsYieldType

// Fetch.ProfileGames (Request)

type FetchProfileGamesType struct {}

var _ RequestMessage = (*FetchProfileGamesType)(nil)

func (r *FetchProfileGamesType) Method() string {
  return "Fetch.ProfileGames"
}

func (r *FetchProfileGamesType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.FetchProfileGamesParams) (*buse.FetchProfileGamesResult, error)) {
  router.Register("Fetch.ProfileGames", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.FetchProfileGamesParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Fetch.ProfileGames")
    }
    return res, nil
  })
}

var FetchProfileGames *FetchProfileGamesType

// Fetch.ProfileGames.Yield (Notification)

type FetchProfileGamesYieldType struct {}

var _ NotificationMessage = (*FetchProfileGamesYieldType)(nil)

func (r *FetchProfileGamesYieldType) Method() string {
  return "Fetch.ProfileGames.Yield"
}

func (r *FetchProfileGamesYieldType) Notify(rc *buse.RequestContext, params *buse.FetchProfileGamesYieldNotification) (error) {
  return rc.Notify("Fetch.ProfileGames.Yield", params)
}

var FetchProfileGamesYield *FetchProfileGamesYieldType

// Fetch.ProfileOwnedKeys (Request)

type FetchProfileOwnedKeysType struct {}

var _ RequestMessage = (*FetchProfileOwnedKeysType)(nil)

func (r *FetchProfileOwnedKeysType) Method() string {
  return "Fetch.ProfileOwnedKeys"
}

func (r *FetchProfileOwnedKeysType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.FetchProfileOwnedKeysParams) (*buse.FetchProfileOwnedKeysResult, error)) {
  router.Register("Fetch.ProfileOwnedKeys", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.FetchProfileOwnedKeysParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Fetch.ProfileOwnedKeys")
    }
    return res, nil
  })
}

var FetchProfileOwnedKeys *FetchProfileOwnedKeysType

// Fetch.ProfileOwnedKeys.Yield (Notification)

type FetchProfileOwnedKeysYieldType struct {}

var _ NotificationMessage = (*FetchProfileOwnedKeysYieldType)(nil)

func (r *FetchProfileOwnedKeysYieldType) Method() string {
  return "Fetch.ProfileOwnedKeys.Yield"
}

func (r *FetchProfileOwnedKeysYieldType) Notify(rc *buse.RequestContext, params *buse.FetchProfileOwnedKeysYieldNotification) (error) {
  return rc.Notify("Fetch.ProfileOwnedKeys.Yield", params)
}

var FetchProfileOwnedKeysYield *FetchProfileOwnedKeysYieldType


//==============================
// Install
//==============================

// Game.FindUploads (Request)

type GameFindUploadsType struct {}

var _ RequestMessage = (*GameFindUploadsType)(nil)

func (r *GameFindUploadsType) Method() string {
  return "Game.FindUploads"
}

func (r *GameFindUploadsType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.GameFindUploadsParams) (*buse.GameFindUploadsResult, error)) {
  router.Register("Game.FindUploads", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.GameFindUploadsParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Game.FindUploads")
    }
    return res, nil
  })
}

var GameFindUploads *GameFindUploadsType

// Operation.Start (Request)

type OperationStartType struct {}

var _ RequestMessage = (*OperationStartType)(nil)

func (r *OperationStartType) Method() string {
  return "Operation.Start"
}

func (r *OperationStartType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.OperationStartParams) (*buse.OperationStartResult, error)) {
  router.Register("Operation.Start", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.OperationStartParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Operation.Start")
    }
    return res, nil
  })
}

var OperationStart *OperationStartType

// Operation.Cancel (Request)

type OperationCancelType struct {}

var _ RequestMessage = (*OperationCancelType)(nil)

func (r *OperationCancelType) Method() string {
  return "Operation.Cancel"
}

func (r *OperationCancelType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.OperationCancelParams) (*buse.OperationCancelResult, error)) {
  router.Register("Operation.Cancel", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.OperationCancelParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Operation.Cancel")
    }
    return res, nil
  })
}

var OperationCancel *OperationCancelType

// PickUpload (Request)

type PickUploadType struct {}

var _ RequestMessage = (*PickUploadType)(nil)

func (r *PickUploadType) Method() string {
  return "PickUpload"
}

func (r *PickUploadType) Call(rc *buse.RequestContext, params *buse.PickUploadParams) (*buse.PickUploadResult, error) {
  var result buse.PickUploadResult
  err := rc.Call("PickUpload", params, &result)
  return &result, err
}

var PickUpload *PickUploadType

// GetReceipt (Request)

type GetReceiptType struct {}

var _ RequestMessage = (*GetReceiptType)(nil)

func (r *GetReceiptType) Method() string {
  return "GetReceipt"
}

func (r *GetReceiptType) Call(rc *buse.RequestContext, params *buse.GetReceiptParams) (*buse.GetReceiptResult, error) {
  var result buse.GetReceiptResult
  err := rc.Call("GetReceipt", params, &result)
  return &result, err
}

var GetReceipt *GetReceiptType

// Operation.Progress (Notification)

type OperationProgressType struct {}

var _ NotificationMessage = (*OperationProgressType)(nil)

func (r *OperationProgressType) Method() string {
  return "Operation.Progress"
}

func (r *OperationProgressType) Notify(rc *buse.RequestContext, params *buse.OperationProgressNotification) (error) {
  return rc.Notify("Operation.Progress", params)
}

var OperationProgress *OperationProgressType

// TaskStarted (Notification)

type TaskStartedType struct {}

var _ NotificationMessage = (*TaskStartedType)(nil)

func (r *TaskStartedType) Method() string {
  return "TaskStarted"
}

func (r *TaskStartedType) Notify(rc *buse.RequestContext, params *buse.TaskStartedNotification) (error) {
  return rc.Notify("TaskStarted", params)
}

var TaskStarted *TaskStartedType

// TaskSucceeded (Notification)

type TaskSucceededType struct {}

var _ NotificationMessage = (*TaskSucceededType)(nil)

func (r *TaskSucceededType) Method() string {
  return "TaskSucceeded"
}

func (r *TaskSucceededType) Notify(rc *buse.RequestContext, params *buse.TaskSucceededNotification) (error) {
  return rc.Notify("TaskSucceeded", params)
}

var TaskSucceeded *TaskSucceededType


//==============================
// Update
//==============================

// CheckUpdate (Request)

type CheckUpdateType struct {}

var _ RequestMessage = (*CheckUpdateType)(nil)

func (r *CheckUpdateType) Method() string {
  return "CheckUpdate"
}

func (r *CheckUpdateType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.CheckUpdateParams) (*buse.CheckUpdateResult, error)) {
  router.Register("CheckUpdate", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.CheckUpdateParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for CheckUpdate")
    }
    return res, nil
  })
}

var CheckUpdate *CheckUpdateType

// GameUpdateAvailable (Notification)

type GameUpdateAvailableType struct {}

var _ NotificationMessage = (*GameUpdateAvailableType)(nil)

func (r *GameUpdateAvailableType) Method() string {
  return "GameUpdateAvailable"
}

func (r *GameUpdateAvailableType) Notify(rc *buse.RequestContext, params *buse.GameUpdateAvailableNotification) (error) {
  return rc.Notify("GameUpdateAvailable", params)
}

var GameUpdateAvailable *GameUpdateAvailableType


//==============================
// Launch
//==============================

// Launch (Request)

type LaunchType struct {}

var _ RequestMessage = (*LaunchType)(nil)

func (r *LaunchType) Method() string {
  return "Launch"
}

func (r *LaunchType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.LaunchParams) (*buse.LaunchResult, error)) {
  router.Register("Launch", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.LaunchParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Launch")
    }
    return res, nil
  })
}

var Launch *LaunchType

// LaunchRunning (Notification)

type LaunchRunningType struct {}

var _ NotificationMessage = (*LaunchRunningType)(nil)

func (r *LaunchRunningType) Method() string {
  return "LaunchRunning"
}

func (r *LaunchRunningType) Notify(rc *buse.RequestContext, params *buse.LaunchRunningNotification) (error) {
  return rc.Notify("LaunchRunning", params)
}

var LaunchRunning *LaunchRunningType

// LaunchExited (Notification)

type LaunchExitedType struct {}

var _ NotificationMessage = (*LaunchExitedType)(nil)

func (r *LaunchExitedType) Method() string {
  return "LaunchExited"
}

func (r *LaunchExitedType) Notify(rc *buse.RequestContext, params *buse.LaunchExitedNotification) (error) {
  return rc.Notify("LaunchExited", params)
}

var LaunchExited *LaunchExitedType

// PickManifestAction (Request)

type PickManifestActionType struct {}

var _ RequestMessage = (*PickManifestActionType)(nil)

func (r *PickManifestActionType) Method() string {
  return "PickManifestAction"
}

func (r *PickManifestActionType) Call(rc *buse.RequestContext, params *buse.PickManifestActionParams) (*buse.PickManifestActionResult, error) {
  var result buse.PickManifestActionResult
  err := rc.Call("PickManifestAction", params, &result)
  return &result, err
}

var PickManifestAction *PickManifestActionType

// ShellLaunch (Request)

type ShellLaunchType struct {}

var _ RequestMessage = (*ShellLaunchType)(nil)

func (r *ShellLaunchType) Method() string {
  return "ShellLaunch"
}

func (r *ShellLaunchType) Call(rc *buse.RequestContext, params *buse.ShellLaunchParams) (*buse.ShellLaunchResult, error) {
  var result buse.ShellLaunchResult
  err := rc.Call("ShellLaunch", params, &result)
  return &result, err
}

var ShellLaunch *ShellLaunchType

// HTMLLaunch (Request)

type HTMLLaunchType struct {}

var _ RequestMessage = (*HTMLLaunchType)(nil)

func (r *HTMLLaunchType) Method() string {
  return "HTMLLaunch"
}

func (r *HTMLLaunchType) Call(rc *buse.RequestContext, params *buse.HTMLLaunchParams) (*buse.HTMLLaunchResult, error) {
  var result buse.HTMLLaunchResult
  err := rc.Call("HTMLLaunch", params, &result)
  return &result, err
}

var HTMLLaunch *HTMLLaunchType

// URLLaunch (Request)

type URLLaunchType struct {}

var _ RequestMessage = (*URLLaunchType)(nil)

func (r *URLLaunchType) Method() string {
  return "URLLaunch"
}

func (r *URLLaunchType) Call(rc *buse.RequestContext, params *buse.URLLaunchParams) (*buse.URLLaunchResult, error) {
  var result buse.URLLaunchResult
  err := rc.Call("URLLaunch", params, &result)
  return &result, err
}

var URLLaunch *URLLaunchType

// SaveVerdict (Request)

type SaveVerdictType struct {}

var _ RequestMessage = (*SaveVerdictType)(nil)

func (r *SaveVerdictType) Method() string {
  return "SaveVerdict"
}

func (r *SaveVerdictType) Call(rc *buse.RequestContext, params *buse.SaveVerdictParams) (*buse.SaveVerdictResult, error) {
  var result buse.SaveVerdictResult
  err := rc.Call("SaveVerdict", params, &result)
  return &result, err
}

var SaveVerdict *SaveVerdictType

// AllowSandboxSetup (Request)

type AllowSandboxSetupType struct {}

var _ RequestMessage = (*AllowSandboxSetupType)(nil)

func (r *AllowSandboxSetupType) Method() string {
  return "AllowSandboxSetup"
}

func (r *AllowSandboxSetupType) Call(rc *buse.RequestContext, params *buse.AllowSandboxSetupParams) (*buse.AllowSandboxSetupResult, error) {
  var result buse.AllowSandboxSetupResult
  err := rc.Call("AllowSandboxSetup", params, &result)
  return &result, err
}

var AllowSandboxSetup *AllowSandboxSetupType

// PrereqsStarted (Notification)

type PrereqsStartedType struct {}

var _ NotificationMessage = (*PrereqsStartedType)(nil)

func (r *PrereqsStartedType) Method() string {
  return "PrereqsStarted"
}

func (r *PrereqsStartedType) Notify(rc *buse.RequestContext, params *buse.PrereqsStartedNotification) (error) {
  return rc.Notify("PrereqsStarted", params)
}

var PrereqsStarted *PrereqsStartedType

// PrereqsTaskState (Notification)

type PrereqsTaskStateType struct {}

var _ NotificationMessage = (*PrereqsTaskStateType)(nil)

func (r *PrereqsTaskStateType) Method() string {
  return "PrereqsTaskState"
}

func (r *PrereqsTaskStateType) Notify(rc *buse.RequestContext, params *buse.PrereqsTaskStateNotification) (error) {
  return rc.Notify("PrereqsTaskState", params)
}

var PrereqsTaskState *PrereqsTaskStateType

// PrereqsEnded (Notification)

type PrereqsEndedType struct {}

var _ NotificationMessage = (*PrereqsEndedType)(nil)

func (r *PrereqsEndedType) Method() string {
  return "PrereqsEnded"
}

func (r *PrereqsEndedType) Notify(rc *buse.RequestContext, params *buse.PrereqsEndedNotification) (error) {
  return rc.Notify("PrereqsEnded", params)
}

var PrereqsEnded *PrereqsEndedType

// PrereqsFailed (Request)

type PrereqsFailedType struct {}

var _ RequestMessage = (*PrereqsFailedType)(nil)

func (r *PrereqsFailedType) Method() string {
  return "PrereqsFailed"
}

func (r *PrereqsFailedType) Call(rc *buse.RequestContext, params *buse.PrereqsFailedParams) (*buse.PrereqsFailedResult, error) {
  var result buse.PrereqsFailedResult
  err := rc.Call("PrereqsFailed", params, &result)
  return &result, err
}

var PrereqsFailed *PrereqsFailedType


//==============================
// Clean Downloads
//==============================

// CleanDownloads.Search (Request)

type CleanDownloadsSearchType struct {}

var _ RequestMessage = (*CleanDownloadsSearchType)(nil)

func (r *CleanDownloadsSearchType) Method() string {
  return "CleanDownloads.Search"
}

func (r *CleanDownloadsSearchType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.CleanDownloadsSearchParams) (*buse.CleanDownloadsSearchResult, error)) {
  router.Register("CleanDownloads.Search", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.CleanDownloadsSearchParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for CleanDownloads.Search")
    }
    return res, nil
  })
}

var CleanDownloadsSearch *CleanDownloadsSearchType

// CleanDownloads.Apply (Request)

type CleanDownloadsApplyType struct {}

var _ RequestMessage = (*CleanDownloadsApplyType)(nil)

func (r *CleanDownloadsApplyType) Method() string {
  return "CleanDownloads.Apply"
}

func (r *CleanDownloadsApplyType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.CleanDownloadsApplyParams) (*buse.CleanDownloadsApplyResult, error)) {
  router.Register("CleanDownloads.Apply", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.CleanDownloadsApplyParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for CleanDownloads.Apply")
    }
    return res, nil
  })
}

var CleanDownloadsApply *CleanDownloadsApplyType


//==============================
// Test
//==============================

// Test.DoubleTwice (Request)

type TestDoubleTwiceType struct {}

var _ RequestMessage = (*TestDoubleTwiceType)(nil)

func (r *TestDoubleTwiceType) Method() string {
  return "Test.DoubleTwice"
}

func (r *TestDoubleTwiceType) Register(router *buse.Router, f func(*buse.RequestContext, *buse.TestDoubleTwiceParams) (*buse.TestDoubleTwiceResult, error)) {
  router.Register("Test.DoubleTwice", func (rc *buse.RequestContext) (interface{}, error) {
    var params buse.TestDoubleTwiceParams
    err := json.Unmarshal(*rc.Params, &params)
    if err != nil {
    	return nil, &buse.RpcError{Code: jsonrpc2.CodeParseError, Message: err.Error()}
    }
    res, err := f(rc, &params)
    if err != nil {
    	return nil, err
    }
    if res == nil {
    	return nil, errors.New("internal error: nil result for Test.DoubleTwice")
    }
    return res, nil
  })
}

var TestDoubleTwice *TestDoubleTwiceType

// Test.Double (Request)

type TestDoubleType struct {}

var _ RequestMessage = (*TestDoubleType)(nil)

func (r *TestDoubleType) Method() string {
  return "Test.Double"
}

func (r *TestDoubleType) Call(rc *buse.RequestContext, params *buse.TestDoubleParams) (*buse.TestDoubleResult, error) {
  var result buse.TestDoubleResult
  err := rc.Call("Test.Double", params, &result)
  return &result, err
}

var TestDouble *TestDoubleType


func EnsureAllRequests(router *buse.Router) {
  if _, ok := router.Handlers["Version.Get"]; !ok { panic("missing request handler for (Version.Get)") }
  if _, ok := router.Handlers["Profile.List"]; !ok { panic("missing request handler for (Profile.List)") }
  if _, ok := router.Handlers["Profile.LoginWithPassword"]; !ok { panic("missing request handler for (Profile.LoginWithPassword)") }
  if _, ok := router.Handlers["Profile.UseSavedLogin"]; !ok { panic("missing request handler for (Profile.UseSavedLogin)") }
  if _, ok := router.Handlers["Profile.Forget"]; !ok { panic("missing request handler for (Profile.Forget)") }
  if _, ok := router.Handlers["Fetch.Game"]; !ok { panic("missing request handler for (Fetch.Game)") }
  if _, ok := router.Handlers["Fetch.Collection"]; !ok { panic("missing request handler for (Fetch.Collection)") }
  if _, ok := router.Handlers["Fetch.ProfileCollections"]; !ok { panic("missing request handler for (Fetch.ProfileCollections)") }
  if _, ok := router.Handlers["Fetch.ProfileGames"]; !ok { panic("missing request handler for (Fetch.ProfileGames)") }
  if _, ok := router.Handlers["Fetch.ProfileOwnedKeys"]; !ok { panic("missing request handler for (Fetch.ProfileOwnedKeys)") }
  if _, ok := router.Handlers["Game.FindUploads"]; !ok { panic("missing request handler for (Game.FindUploads)") }
  if _, ok := router.Handlers["Operation.Start"]; !ok { panic("missing request handler for (Operation.Start)") }
  if _, ok := router.Handlers["Operation.Cancel"]; !ok { panic("missing request handler for (Operation.Cancel)") }
  if _, ok := router.Handlers["CheckUpdate"]; !ok { panic("missing request handler for (CheckUpdate)") }
  if _, ok := router.Handlers["Launch"]; !ok { panic("missing request handler for (Launch)") }
  if _, ok := router.Handlers["CleanDownloads.Search"]; !ok { panic("missing request handler for (CleanDownloads.Search)") }
  if _, ok := router.Handlers["CleanDownloads.Apply"]; !ok { panic("missing request handler for (CleanDownloads.Apply)") }
  if _, ok := router.Handlers["Test.DoubleTwice"]; !ok { panic("missing request handler for (Test.DoubleTwice)") }
}

