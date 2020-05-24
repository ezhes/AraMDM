package main

import "log"

const (
	//Sets the global logger flags
	LOGGER_FLAGS = log.Ldate | log.Ltime | log.Lshortfile

	/*
		MicroMDM configuration
	*/
	//The base address for the API server
	MICROMDM_API_BASE_URL = ""

	//The expected source for webhooks. This is used to prevent spoofing.
	MICROMDM_EXPECTED_SRC = "192.168.1.30"
	//Port at which the webhook server should communicate with (results in a webhook URL  http://0.0.0.0:<port>/ara-hook)
	MICROMDM_WEBHOOK_PORT = 52801
)
