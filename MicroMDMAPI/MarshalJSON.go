package MicroMDMAPI

//https://github.com/micromdm/micromdm/blob/e2acfb977e8c42976edd0a0b9b0e8fe49898d91d/mdm/mdm/marshal_json.go

import (
	"encoding/json"
	"fmt"
)

func (cr *CommandRequest) MarshalJSON() ([]byte, error) {
	c := cr.Command
	switch c.RequestType {
	case "ProfileList",
		"ProvisioningProfileList",
		"CertificateList",
		"SecurityInfo",
		"RestartDevice",
		"ShutDownDevice",
		"StopMirroring",
		"ClearRestrictionsPassword",
		"UserList",
		"LogOutUser",
		"PlayLostModeSound",
		"DisableLostMode",
		"DeviceLocation",
		"ManagedMediaList",
		"DeviceConfigured",
		"AvailableOSUpdates",
		"NSExtensionMappings",
		"OSUpdateStatus",
		"EnableRemoteDesktop",
		"DisableRemoteDesktop",
		"ActivationLockBypassCode":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
		}{
			UDID:        cr.UDID,
			RequestType: c.RequestType,
		}
		return json.Marshal(&x)
	case "InstallProfile":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*InstallProfile
		}{
			UDID:           cr.UDID,
			RequestType:    c.RequestType,
			InstallProfile: c.InstallProfile,
		}
		return json.Marshal(&x)
	case "RemoveProfile":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*RemoveProfile
		}{
			UDID:          cr.UDID,
			RequestType:   c.RequestType,
			RemoveProfile: c.RemoveProfile,
		}
		return json.Marshal(&x)
	case "InstallProvisioningProfile":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*InstallProvisioningProfile
		}{
			UDID:                       cr.UDID,
			RequestType:                c.RequestType,
			InstallProvisioningProfile: c.InstallProvisioningProfile,
		}
		return json.Marshal(&x)
	case "RemoveProvisioningProfile":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*RemoveProvisioningProfile
		}{
			UDID:                      cr.UDID,
			RequestType:               c.RequestType,
			RemoveProvisioningProfile: c.RemoveProvisioningProfile,
		}
		return json.Marshal(&x)
	case "InstalledApplicationList":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*InstalledApplicationList
		}{
			UDID:                     cr.UDID,
			RequestType:              c.RequestType,
			InstalledApplicationList: c.InstalledApplicationList,
		}
		return json.Marshal(&x)
	case "DeviceInformation":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*DeviceInformation
		}{
			UDID:              cr.UDID,
			RequestType:       c.RequestType,
			DeviceInformation: c.DeviceInformation,
		}
		return json.Marshal(&x)
	case "DeviceLock":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*DeviceLock
		}{
			UDID:        cr.UDID,
			RequestType: c.RequestType,
			DeviceLock:  c.DeviceLock,
		}
		return json.Marshal(&x)
	case "ClearPasscode":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*ClearPasscode
		}{
			UDID:          cr.UDID,
			RequestType:   c.RequestType,
			ClearPasscode: c.ClearPasscode,
		}
		return json.Marshal(&x)
	case "EraseDevice":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*EraseDevice
		}{
			UDID:        cr.UDID,
			RequestType: c.RequestType,
			EraseDevice: c.EraseDevice,
		}
		return json.Marshal(&x)
	case "RequestMirroring":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*RequestMirroring
		}{
			UDID:             cr.UDID,
			RequestType:      c.RequestType,
			RequestMirroring: c.RequestMirroring,
		}
		return json.Marshal(&x)
	case "Restrictions":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*Restrictions
		}{
			UDID:         cr.UDID,
			RequestType:  c.RequestType,
			Restrictions: c.Restrictions,
		}
		return json.Marshal(&x)
	case "UnlockUserAccount":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*UnlockUserAccount
		}{
			UDID:              cr.UDID,
			RequestType:       c.RequestType,
			UnlockUserAccount: c.UnlockUserAccount,
		}
		return json.Marshal(&x)
	case "DeleteUser":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*DeleteUser
		}{
			UDID:        cr.UDID,
			RequestType: c.RequestType,
			DeleteUser:  c.DeleteUser,
		}
		return json.Marshal(&x)
	case "EnableLostMode":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*EnableLostMode
		}{
			UDID:           cr.UDID,
			RequestType:    c.RequestType,
			EnableLostMode: c.EnableLostMode,
		}
		return json.Marshal(&x)
	case "InstallApplication":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*InstallApplication
		}{
			UDID:               cr.UDID,
			RequestType:        c.RequestType,
			InstallApplication: c.InstallApplication,
		}
		return json.Marshal(&x)
	case "InstallEnterpriseApplication":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*InstallEnterpriseApplication
		}{
			UDID:                         cr.UDID,
			RequestType:                  c.RequestType,
			InstallEnterpriseApplication: c.InstallEnterpriseApplication,
		}
		return json.Marshal(&x)
	case "AccountConfiguration":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*AccountConfiguration
		}{
			UDID:                 cr.UDID,
			RequestType:          c.RequestType,
			AccountConfiguration: c.AccountConfiguration,
		}
		return json.Marshal(&x)
	case "ApplyRedemptionCode":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*ApplyRedemptionCode
		}{
			UDID:                cr.UDID,
			RequestType:         c.RequestType,
			ApplyRedemptionCode: c.ApplyRedemptionCode,
		}
		return json.Marshal(&x)
	case "ManagedApplicationList":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*ManagedApplicationList
		}{
			UDID:                   cr.UDID,
			RequestType:            c.RequestType,
			ManagedApplicationList: c.ManagedApplicationList,
		}
		return json.Marshal(&x)
	case "RemoveApplication":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*RemoveApplication
		}{
			UDID:              cr.UDID,
			RequestType:       c.RequestType,
			RemoveApplication: c.RemoveApplication,
		}
		return json.Marshal(&x)
	case "InviteToProgram":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*InviteToProgram
		}{
			UDID:            cr.UDID,
			RequestType:     c.RequestType,
			InviteToProgram: c.InviteToProgram,
		}
		return json.Marshal(&x)
	case "ValidateApplications":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*ValidateApplications
		}{
			UDID:                 cr.UDID,
			RequestType:          c.RequestType,
			ValidateApplications: c.ValidateApplications,
		}
		return json.Marshal(&x)
	case "InstallMedia":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*InstallMedia
		}{
			UDID:         cr.UDID,
			RequestType:  c.RequestType,
			InstallMedia: c.InstallMedia,
		}
		return json.Marshal(&x)
	case "RemoveMedia":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*RemoveMedia
		}{
			UDID:        cr.UDID,
			RequestType: c.RequestType,
			RemoveMedia: c.RemoveMedia,
		}
		return json.Marshal(&x)
	case "Settings":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*Settings
		}{
			UDID:        cr.UDID,
			RequestType: c.RequestType,
			Settings:    c.Settings,
		}
		return json.Marshal(&x)
	case "ManagedApplicationConfiguration":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*ManagedApplicationConfiguration
		}{
			UDID:                            cr.UDID,
			RequestType:                     c.RequestType,
			ManagedApplicationConfiguration: c.ManagedApplicationConfiguration,
		}
		return json.Marshal(&x)
	case "ManagedApplicationAttributes":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*ManagedApplicationAttributes
		}{
			UDID:                         cr.UDID,
			RequestType:                  c.RequestType,
			ManagedApplicationAttributes: c.ManagedApplicationAttributes,
		}
		return json.Marshal(&x)
	case "ManagedApplicationFeedback":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*ManagedApplicationFeedback
		}{
			UDID:                       cr.UDID,
			RequestType:                c.RequestType,
			ManagedApplicationFeedback: c.ManagedApplicationFeedback,
		}
		return json.Marshal(&x)
	case "SetFirmwarePassword":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*SetFirmwarePassword
		}{
			UDID:                cr.UDID,
			RequestType:         c.RequestType,
			SetFirmwarePassword: c.SetFirmwarePassword,
		}
		return json.Marshal(&x)
	case "VerifyFirmwarePassword":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*VerifyFirmwarePassword
		}{
			UDID:                   cr.UDID,
			RequestType:            c.RequestType,
			VerifyFirmwarePassword: c.VerifyFirmwarePassword,
		}
		return json.Marshal(&x)
	case "SetAutoAdminPassword":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*SetAutoAdminPassword
		}{
			UDID:                 cr.UDID,
			RequestType:          c.RequestType,
			SetAutoAdminPassword: c.SetAutoAdminPassword,
		}
		return json.Marshal(&x)
	case "ScheduleOSUpdate":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*ScheduleOSUpdate
		}{
			UDID:             cr.UDID,
			RequestType:      c.RequestType,
			ScheduleOSUpdate: c.ScheduleOSUpdate,
		}
		return json.Marshal(&x)
	case "ScheduleOSUpdateScan":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*ScheduleOSUpdateScan
		}{
			UDID:                 cr.UDID,
			RequestType:          c.RequestType,
			ScheduleOSUpdateScan: c.ScheduleOSUpdateScan,
		}
		return json.Marshal(&x)
	case "ActiveNSExtensions":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*ActiveNSExtensions
		}{
			UDID:               cr.UDID,
			RequestType:        c.RequestType,
			ActiveNSExtensions: c.ActiveNSExtensions,
		}
		return json.Marshal(&x)
	case "RotateFileVaultKey":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*RotateFileVaultKey
		}{
			UDID:               cr.UDID,
			RequestType:        c.RequestType,
			RotateFileVaultKey: c.RotateFileVaultKey,
		}
		return json.Marshal(&x)
	case "SetBootstrapToken":
		var x = struct {
			UDID        string `json:"udid"`
			RequestType string `json:"request_type"`
			*SetBootstrapToken
		}{
			UDID:              cr.UDID,
			RequestType:       c.RequestType,
			SetBootstrapToken: c.SetBootstrapToken,
		}
		return json.Marshal(&x)
	default:
		return nil, fmt.Errorf("mdm: unknown RequestType: %s", c.RequestType)
	}
}
