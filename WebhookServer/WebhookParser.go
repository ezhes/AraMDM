package WebhookServer

import (
	"fmt"
	"howett.net/plist"
	"time"
)

type StringAnyMap map[string]interface{}

type Event struct {
	Topic     string    `json:"topic"`
	EventID   string    `json:"event_id"`
	CreatedAt time.Time `json:"created_at"`

	AcknowledgeEvent *AcknowledgeEvent `json:"acknowledge_event,omitempty"`
	CheckinEvent     *CheckinEvent     `json:"checkin_event,omitempty"`
}

type AcknowledgeEvent struct {
	UDID        string            `json:"udid"`
	Status      string            `json:"status"`
	CommandUUID string            `json:"command_uuid"`
	Params      map[string]string `json:"url_params"`
	RawPayload  []byte            `json:"raw_payload"`
}

type CheckinEvent struct {
	UDID       string            `json:"udid"`
	Params     map[string]string `json:"url_params"`
	RawPayload []byte            `json:"raw_payload"`
}

type Command struct {
	UDID        string
	CommandUUID string
	Status      string /* Use acknowledgeStateStringToEnum */

	/* Payloads -- most will be nil, meaning they are not this message type */
	QueryResponses *QueryResponsesT
}

type QueryResponsesT struct {
	AppAnalyticsEnabled             *bool
	AutomaticAppInstallationEnabled *bool
	AutomaticCheckEnabled           *bool
	AutomaticOSInstallationEnabled  *bool
	AutomaticSecurityUpdatesEnabled *bool
	AvailableDeviceCapacity         *float64
	AwaitingConfiguration           *bool
	BackgroundDownloadEnabled       *bool
	BatteryLevel                    *float64
	BluetoothMAC                    *string
	BuildVersion                    *string
	CarrierSettingsVersion          *string
	CatalogURL                      *string
	CellularTechnology              *int
	CurrentCarrierNetwork           *string
	CurrentMCC                      *string
	CurrentMNC                      *string
	DataRoamingEnabled              *bool
	DeviceCapacity                  *float64
	DeviceID                        *string
	DeviceName                      *string
	DiagnosticSubmissionEnabled     *bool
	EASDeviceIdentifier             *string
	EthernetMACs                    *[]string
	HostName                        *string
	ICCID                           *string
	IMEI                            *string
	IsActivationLockEnabled         *bool
	IsCloudBackupEnabled            *bool
	IsDefaultCatalog                *bool
	IsDeviceLocatorServiceEnabled   *bool
	IsDoNotDisturbInEffect          *bool
	IsMDMLostModeEnabled            *bool
	IsMultiUser                     *bool
	IsNetworkTethered               *bool
	IsRoaming                       *bool
	IsSupervised                    *bool
	Languages                       *[]string
	LastCloudBackupDate             *time.Time
	LocalHostName                   *string
	Locales                         *[]string
	MDMOptions                      *MDMOptionsT
	MEID                            *string
	MaximumResidentUsers            *int
	Model                           *string
	ModelName                       *string
	ModemFirmwareVersion            *string
	OSUpdateSettings                *OSUpdateSettingsT
	OSVersion                       *string
	OrganizationInfo                *string
	PerformPeriodicCheck            *bool
	PersonalHotspotEnabled          *bool
	PhoneNumber                     *string
	PreviousScanDate                *time.Time
	PreviousScanResult              *int
	ProductName                     *string
	SIMMCC                          *string
	SerialNumber                    *string
	ServiceSubscriptions            *[]ServiceSubscriptionsT
	SubscriberCarrierNetwork        *string
	SubscriberMCC                   *string
	SubscriberMNC                   *string
	UDID                            *string
	VoiceRoamingEnabled             *bool
	WiFiMAC                         *string
	iTunesStoreAccountHash          *string
	iTunesStoreAccountIsActive      *bool
}
type MDMOptionsT struct {
	ActivationLockAllowedWhileSupervised bool
}
type OSUpdateSettingsT struct {
	CatalogURL                      string
	IsDefaultCatalog                bool
	PreviousScanDate                time.Time
	PreviousScanResult              int
	PerformPeriodicCheck            bool
	AutomaticCheckEnabled           bool
	BackgroundDownloadEnabled       bool
	AutomaticAppInstallationEnabled bool
	AutomaticOSInstallationEnabled  bool
	AutomaticSecurityUpdatesEnabled bool
}
type ServiceSubscriptionsT struct {
	CarrierSettingsVersion string
	CurrentCarrierNetwork  string
	CurrentMCC             string
	CurrentMNC             string
	ICCID                  string
	IMEI                   string
	IsDataPreferred        bool
	IsRoaming              bool
	IsVoicePreferred       bool
	Label                  string
	LabelID                string
	PhoneNumber            string
	Slot                   string
}
type AcknowledgeEventState int

const (
	ACKNOWLEDGED AcknowledgeEventState = iota
	ERROR
	COMMAND_FORMAT_ERROR
	IDLE
	NOT_NOW
	UNKNOWN
)

func (s AcknowledgeEventState) String() string {
	return [...]string{"ACKNOWLEDGED", "ERROR", "COMMAND_FORMAT_ERROR", "IDLE", "NOT_NOW", "UNKNOWN"}[s]
}

func acknowledgeStateStringToEnum(state string) AcknowledgeEventState {
	switch state {
	case "Acknowledged":
		return ACKNOWLEDGED
	case "CommandFormatError":
		return COMMAND_FORMAT_ERROR
	case "Error":
		return ERROR
	case "NotNow":
		return NOT_NOW
	case "Idle":
		return IDLE
	default:
		activeServer.log.Printf("Found known state: " + state)
		return UNKNOWN
	}
}

func parseAcknowledge(data []byte) *Command {
	fmt.Println(string(data))
	var obj Command
	_, err := plist.Unmarshal(data, &obj)
	if err != nil {
		activeServer.log.Printf("Unable to decode PLIST: " + err.Error())
		return nil
	}

	return &obj
}
