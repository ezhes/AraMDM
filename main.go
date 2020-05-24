package main

import (
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

	runLoop()
}

func runLoop() {
	logger.Printf("Entering runloop")
	for {
		select {}
	}
}
