package main

import (
	"AraMDM/MicroMDMAPI"
	"AraMDM/WebhookServer"
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "[Main] ", LOGGER_FLAGS)

func main() {
	logger.Printf("----Launching AraMDM----")
	ws := WebhookServer.New(MICROMDM_EXPECTED_SRC, LOGGER_FLAGS)
	ws.Launch(MICROMDM_WEBHOOK_PORT)

	mAPI := MicroMDMAPI.New(MICROMDM_API_BASE_URL, MICROMDM_API_KEY, LOGGER_FLAGS)

	resp, err :=
		mAPI.ExecuteMDMCommand(MicroMDMAPI.CommandRequest{
			UDID: "797566FD-5E51-5519-ACE9-7C299284BE11",
			Command: &MicroMDMAPI.Command{
				RequestType: "DeviceInformation",
				DeviceInformation: &MicroMDMAPI.DeviceInformation{
					Queries: []string{"DeviceName", "BatteryLevel"},
				},
			},
		})
	//	mAPI.ExecuteMDMCommand(MicroMDMAPI.CommandRequest{
	//	UDID:    "797566FD-5E51-5519-ACE9-7C299284BE11",
	//	Command: &MicroMDMAPI.Command{
	//		RequestType:   "RemoveProfile",
	//		RemoveProfile: &MicroMDMAPI.RemoveProfile{Identifier:"com.unwiredmdm.mobileconfig.profile-service"},
	//	},
	//})

	if err != nil {
		println(err.Error())
	} else if resp != nil {
		fmt.Printf("Sent command %s\n", resp.Payload.CommandUUID)
	}

	runLoop()
}

func runLoop() {
	logger.Printf("Entering runloop")
	for {
		select {}
	}
}
