package MicroMDMAPI

import "encoding/json"

//FROM https://github.com/micromdm/micromdm/blob/e2acfb977e8c42976edd0a0b9b0e8fe49898d91d/mdm/mdm/command.go

type CommandRequest struct {
	UDID string `json:"udid"`
	*Command
}

type CommandResponse struct {
	Payload *CommandPayload `json:"payload,omitempty"`
	Err     *error          `json:"error,omitempty"`
}

type CommandPayload struct {
	CommandUUID string `json:"command_uuid"`
}

type Command struct {
	RequestType                      string `json:"request_type"`
	*InstallProfile                  `json:",omitempty"`
	*RemoveProfile                   `json:",omitempty"`
	*InstallProvisioningProfile      `json:",omitempty"`
	*RemoveProvisioningProfile       `json:",omitempty"`
	*InstalledApplicationList        `json:",omitempty"`
	*DeviceInformation               `json:",omitempty"`
	*DeviceLock                      `json:",omitempty"`
	*ClearPasscode                   `json:",omitempty"`
	*EraseDevice                     `json:",omitempty"`
	*RequestMirroring                `json:",omitempty"`
	*Restrictions                    `json:",omitempty"`
	*UnlockUserAccount               `json:",omitempty"`
	*DeleteUser                      `json:",omitempty"`
	*EnableLostMode                  `json:",omitempty"`
	*InstallApplication              `json:",omitempty"`
	*InstallEnterpriseApplication    `json:",omitempty"`
	*AccountConfiguration            `json:",omitempty"`
	*ApplyRedemptionCode             `json:",omitempty"`
	*ManagedApplicationList          `json:",omitempty"`
	*RemoveApplication               `json:",omitempty"`
	*InviteToProgram                 `json:",omitempty"`
	*ValidateApplications            `json:",omitempty"`
	*InstallMedia                    `json:",omitempty"`
	*RemoveMedia                     `json:",omitempty"`
	*Settings                        `json:",omitempty"`
	*ManagedApplicationConfiguration `json:",omitempty"`
	*ManagedApplicationAttributes    `json:",omitempty"`
	*ManagedApplicationFeedback      `json:",omitempty"`
	*SetFirmwarePassword             `json:",omitempty"`
	*VerifyFirmwarePassword          `json:",omitempty"`
	*SetAutoAdminPassword            `json:",omitempty"`
	*ScheduleOSUpdate                `json:",omitempty"`
	*ScheduleOSUpdateScan            `json:",omitempty"`
	*ActiveNSExtensions              `json:",omitempty"`
	*RotateFileVaultKey              `json:",omitempty"`
	*SetBootstrapToken               `json:",omitempty"`
}

// InstallProfile is an InstallProfile MDM Command
type InstallProfile struct {
	Payload []byte `json:"payload,omitempty"`
}

type RemoveProfile struct {
	Identifier string `json:"identifier,omitempty"`
}

type InstallProvisioningProfile struct {
	ProvisioningProfile []byte `plist:",omitempty" json:"provisioning_profile,omitempty"`
}

type RemoveProvisioningProfile struct {
	UUID string `json:"uuid"`
}

type InstalledApplicationList struct {
	Identifiers     []string `plist:",omitempty" json:"identifiers,omitempty"`
	ManagedAppsOnly bool     `plist:",omitempty" json:"managed_appd_only,omitempty"`
}

type DeviceInformation struct {
	Queries []string `plist:",omitempty" json:"queries,omitempty"`
}

type DeviceLock struct {
	PIN         string `plist:",omitempty" json:"pin"`
	Message     string `plist:",omitempty" json:"message,omitempty"`
	PhoneNumber string `plist:",omitempty" json:"phone_number,omitempty"`
}

type ClearPasscode struct {
	UnlockToken []byte `json:"unlock_token"`
}

type EraseDevice struct {
	PIN                    string `json:"pin"`
	PreserveDataPlan       bool   `plist:",omitempty" json:"preserve_data_plan,omitempty"`
	DisallowProximitySetup bool   `plist:",omitempty" json:"disallow_proximity_setup,omitempty"`
}

type RequestMirroring struct {
	DestinationName     string `plist:",omitempty" json:"destination_name,omitempty"`
	DestinationDeviceID string `plist:",omitempty" json:"destination_device_id,omitempty"`
	ScanTime            string `plist:",omitempty" json:"scan_time,omitempty"`
	Password            string `plist:",omitempty" json:"password,omitempty"`
}

type Restrictions struct {
	ProfileRestrictions bool `json:"profile_restrictions"`
}

type UnlockUserAccount struct {
	UserName string `json:"username"`
}

type DeleteUser struct {
	UserName      string `plist:",omitempty" json:"username,omitempty"`
	ForceDeletion bool   `plist:",omitempty" json:"force_deletion,omitempty"`
}

type EnableLostMode struct {
	Message     string `plist:",omitempty" json:"message,omitempty"`
	PhoneNumber string `plist:",omitempty" json:"phone_number,omitempty"`
	Footnote    string `plist:",omitempty" json:"footnote,omitempty"`
}

type InstallEnterpriseApplication struct {
	//Manifest                       *appmanifest.Manifest `plist:",omitempty" json:"manifest,omitempty"`
	ManifestURL *string `plist:",omitempty" json:"manifest_url,omitempty"`
	//ManifestURLPinningCerts        [][]byte              `plist:",omitempty" json:"manifest_url_pinning_certs,omitempty"`
	//PinningRevocationCheckRequired *bool                 `plist:",omitempty" json:"pinning_revocation_check_required,omitempty"`
}

type InstallApplication struct {
	ITunesStoreID         *int64                           `plist:"iTunesStoreID,omitempty" json:"itunes_store_id,omitempty"`
	Identifier            *string                          `plist:",omitempty" json:"identifier,omitempty"`
	ManagementFlags       *int                             `plist:",omitempty" json:"management_flags,omitempty"`
	ChangeManagementState *string                          `plist:",omitempty" json:"change_management_state,omitempty"`
	ManifestURL           *string                          `plist:",omitempty" json:"manifest_url,omitempty"`
	Options               *InstallApplicationOptions       `plist:"Options,omitempty" json:"options,omitempty"`
	Configuration         *InstallApplicationConfiguration `plist:",omitempty" json:"configuration,omitempty"`
	Attributes            *InstallApplicationAttributes    `plist:",omitempty" json:"attributes,omitempty"`
}

type InstallApplicationOptions struct {
	PurchaseMethod *int64 `plist:"PurchaseMethod,omitempty" json:"purchase_method,omitempty"`
}

type InstallApplicationConfiguration struct{}
type InstallApplicationAttributes struct{}

type AccountConfiguration struct {
	SkipPrimarySetupAccountCreation     bool           `plist:",omitempty" json:"skip_primary_setup_account_creation,omitempty"`
	SetPrimarySetupAccountAsRegularUser bool           `plist:",omitempty" json:"set_primary_setup_account_as_regular_user,omitempty"`
	DontAutoPopulatePrimaryAccountInfo  bool           `plist:",omitempty" json:"dont_auto_populate_primary_account_info,omitempty"`
	LockPrimaryAccountInfo              bool           `plist:",omitempty" json:"lock_primary_account_info,omitempty"`
	PrimaryAccountFullName              string         `plist:",omitempty" json:"primary_account_full_name,omitempty"`
	PrimaryAccountUserName              string         `plist:",omitempty" json:"primary_account_user_name,omitempty"`
	AutoSetupAdminAccounts              []AdminAccount `plist:",omitempty" json:"auto_setup_admin_accounts,omitempty"`
}

type AdminAccount struct {
	ShortName    string `plist:"shortName" json:"short_name"`
	FullName     string `plist:"fullName,omitempty" json:"full_name,omitempty"`
	PasswordHash []byte `plist:"passwordHash" json:"password_hash"`
	Hidden       bool   `plist:"hidden,omitempty" json:"hidden,omitempty"`
}

type ApplyRedemptionCode struct {
	Identifier     string `plist:",omitempty" json:"identifier,omitempty"`
	RedemptionCode string `plist:",omitempty" json:"redemption_code,omitempty"`
}

type ManagedApplicationList struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type RemoveApplication struct {
	Identifier string `plist:",omitempty" json:"identifier,omitempty"`
}

type InviteToProgram struct {
	ProgramID     string `plist:",omitempty" json:"program_id,omitempty"`
	InvitationURL string `plist:",omitempty" json:"invitation_url,omitempty"`
}

type ValidateApplications struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type InstallMedia struct {
	ITunesStoreID *int64 `plist:"iTunesStoreID,omitempty" json:"itunes_store_id,omitempty"`
	MediaURL      string `plist:",omitempty" json:"media_url,omitempty"`
	MediaType     string `plist:",omitempty" json:"media_type,omitempty"`
}

type RemoveMedia struct {
	ITunesStoreID *int64 `plist:"iTunesStoreID,omitempty" json:"itunes_store_id,omitempty"`
	MediaType     string `plist:",omitempty" json:"media_type,omitempty"`
	PersistentID  string `plist:",omitempty" json:"persistent_id,omitempty"`
}

type Settings struct {
	Settings []Setting `plist:",omitempty" json:"settings,omitempty"`
}

type Setting struct {
	Item                    string                 `json:"item"`
	Enabled                 *bool                  `plist:",omitempty" json:"enabled,omitempty"`
	DeviceName              *string                `plist:",omitempty" json:"device_name,omitempty"`
	HostName                *string                `plist:",omitempty" json:"hostname,omitempty"`
	Identifier              *string                `plist:",omitempty" json:"identifier"`
	Attributes              map[string]string      `plist:",omitempty" json:"attributes,omitempty"`
	Image                   []byte                 `plist:",omitempty" json:"image,omitempty"`
	Where                   *int                   `plist:",omitempty" json:"where,omitempty"`
	MDMOptions              map[string]interface{} `plist:",omitempty" json:"mdm_options,omitempty"`
	PasscodeLockGracePeriod *int                   `plist:",omitempty" json:"passcode_lock_grace_period,omitempty"`
	MaximumResidentUsers    *int                   `plist:",omitempty" json:"maximum_resident_users,omitempty"`
	Configuration           map[string]interface{} `plist:",omitempty" json:"-"`
	ConfigurationData       []byte                 `plist:"-" json:"configuration"` // used to build the dictionary
}

type ManagedApplicationConfiguration struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type ManagedApplicationAttributes struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type ManagedApplicationFeedback struct {
	Identifiers    []string `plist:",omitempty" json:"identifiers,omitempty"`
	DeleteFeedback bool     `plist:",omitempty" json:"delete_feedback,omitempty"`
}

type SetFirmwarePassword struct {
	CurrentPassword string `plist:",omitempty" json:"current_password,omitempty"`
	NewPassword     string `plist:",omitempty" json:"new_password,omitempty"`
	AllowOroms      bool   `plist:",omitempty" json:"allow_oroms,omitempty"`
}

type VerifyFirmwarePassword struct {
	Password string `plist:",omitempty" json:"password,omitempty"`
}

type SetAutoAdminPassword struct {
	GUID         string `plist:",omitempty" json:"guid,omitempty"`
	PasswordHash []byte `plist:"passwordHash" json:"password_hash"`
}

type SetBootstrapToken struct {
	BootstrapToken string `plist:",omitempty" json:"bootstrap_token,omitempty"`
}

type OSUpdate struct {
	ProductKey    string `json:"product_key"`
	InstallAction string `json:"install_action"`
}

type ScheduleOSUpdate struct {
	Updates []OSUpdate `plist:",omitempty" json:"updates,omitempty"`
}

type ScheduleOSUpdateScan struct {
	Force bool `plist:",omitempty" json:"force,omitempty"`
}

type ActiveNSExtensions struct {
	FilterExtensionPoints []string `plist:",omitempty" json:"filter_extensions_points,omitempty"`
}

type RotateFileVaultKey struct {
	KeyType                    string          `plist:",omitempty" json:"key_type,omitempty"`
	FileVaultUnlock            FileVaultUnlock `plist:",omitempty" json:"filevault_unlock,omitempty"`
	NewCertificate             []byte          `plist:",omitempty" json:"new_certificate,omitempty"`
	ReplyEncryptionCertificate []byte          `plist:",omitempty" json:"reply_encryption_certificate,omitempty"`
}

type FileVaultUnlock struct {
	Password                 string `plist:",omitempty" json:"password,omitempty"`
	PrivateKeyExport         []byte `plist:",omitempty" json:"private_key_export,omitempty"`
	PrivateKeyExportPassword string `plist:",omitempty" json:"private_key_export_password,omitempty"`
}

func (client *Client) ExecuteMDMCommand(command CommandRequest) (*CommandResponse, error) {
	data, err := client.doAPIRequest("v1/commands", "POST", command)
	if err != nil {
		return nil, err
	}

	response := new(CommandResponse)
	err = json.Unmarshal(data, &response)
	return response, err
}
