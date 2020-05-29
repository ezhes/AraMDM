package MicroMDMAPI

import "encoding/json"

type removeDevicesRequest struct{ Opts RemoveDevicesOptions }

type removeDevicesResponse struct {
	Err error `json:"err,omitempty"`
}

type RemoveDevicesOptions struct {
	UDIDs   []string `json:"udids"`
	Serials []string `json:"serials"`
}

//Removes devices by their UDID or serials from the MicroMDM server. Note that this does not remove a device from management but instead removes the server records
func (client *Client) RemoveDevices(options RemoveDevicesOptions) error {
	request := removeDevicesRequest{Opts: options}
	data, err := client.doAPIRequest("v1/devices", "DELETE", request)
	if err != nil {
		return err
	}

	response := new(removeDevicesResponse)
	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	return response.Err
}
