package MicroMDMAPI

import (
	"encoding/json"
	"time"
)

type getDevicesResponse struct {
	Devices []DeviceDTO `json:"devices"`
	Err     error       `json:"err,omitempty"`
}
type DeviceDTO struct {
	SerialNumber     string           `json:"serial_number"`
	UDID             string           `json:"udid"`
	EnrollmentStatus bool             `json:"enrollment_status"`
	LastSeen         time.Time        `json:"last_seen"`
	DEPProfileStatus DEPProfileStatus `json:"dep_profile_status"`
}

// DEPProfileStatus is the status of the DEP Profile
// can be either "empty", "assigned", "pushed", or "removed"
type DEPProfileStatus string

// DEPProfileStatus values
const (
	EMPTY    DEPProfileStatus = "empty"
	ASSIGNED                  = "assigned"
	PUSHED                    = "pushed"
	REMOVED                   = "removed"
)

type getDevicesRequest struct{ Opts ListDevicesOption }

// As of May 29, 2020 these options are ignored by the backend and all results are returned
type ListDevicesOption struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`

	FilterSerial []string `json:"filter_serial"`
	FilterUDID   []string `json:"filter_udid"`
}

// Gets a list of devices enrolled with the MDM server from MicroMDM
// Currently, optionsUNUSED is unused by MicroMDM and so does nothing. Pass an empty struct.
func (client *Client) GetDevices(optionsUNUSED ListDevicesOption) (*[]DeviceDTO, error) {
	data, err := client.doAPIRequest("v1/devices", "POST", optionsUNUSED)
	if err != nil {
		return nil, err
	}

	response := new(getDevicesResponse)
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response.Devices, response.Err
}
