package main

import (
	"AraMDM/MicroMDMAPI"
	"AraMDM/WebhookServer"
	_ "AraMDM/WebhookServer"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "[Main] ", LOGGER_FLAGS)

func main() {
	logger.Printf("----Launching AraMDM----")
	ws := WebhookServer.New(MICROMDM_EXPECTED_SRC, LOGGER_FLAGS)
	ws.Launch(MICROMDM_WEBHOOK_PORT)

	mAPI := MicroMDMAPI.New(MICROMDM_API_BASE_URL, MICROMDM_API_KEY, LOGGER_FLAGS)
	_, err := mAPI.ExecuteMDMCommand(MicroMDMAPI.CommandRequest{
		UDID: "00008027-000638663C0A002E",
		Command: &MicroMDMAPI.Command{
			RequestType: "DeviceInformation",
			DeviceInformation: &MicroMDMAPI.DeviceInformation{
				Queries: []string{"DeviceName", "BatteryLevel"},
			},
		},
	})

	if err != nil {
		println(err.Error())
	}

	runLoop()
}

func runLoop() {
	logger.Printf("Entering runloop")
	for {
		select {}
	}
}
