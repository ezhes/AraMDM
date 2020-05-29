package WebhookServer

const (
	QUERY_RESPONSES = "QueryResponses"
)

/* ACK events. This includes true ACKs and command responses */
func (server *WebhookServer) handleConnect(event Event) {
	sparse := parseAcknowledge(event.AcknowledgeEvent.RawPayload)
	if sparse == nil {
		return
	}

	statusEnum := acknowledgeStateStringToEnum(sparse.Status)

	switch statusEnum {
	case ACKNOWLEDGED:
		if !server.handleCommandAcknowledgement(*sparse) {
			server.log.Printf("ACK was not handled. Plist:\n%s\n", string(event.AcknowledgeEvent.RawPayload))
		}
		break
	case IDLE, NOT_NOW:
		break
	case COMMAND_FORMAT_ERROR, ERROR:
		server.log.Printf("Command UUID %s gave error %+v", sparse.CommandUUID, sparse.ErrorChain)
	default:
		server.log.Printf("Not handling state " + sparse.Status)
	}
}

func (server *WebhookServer) handleCommandAcknowledgement(command Command) bool {
	if command.QueryResponses != nil {
		server.handleQueryResponse(command)
	} else if command.CommandUUID != "" {
		//KEEP THIS CASE LAST
		//This is to catch empty acknowledges which are given when the device executes a command and just informs us that it did so
		server.handleCommandCompletedOK(command)
	} else {
		return false
	}

	return true
}

func (server *WebhookServer) handleQueryResponse(command Command) {
	server.log.Println("Got QueryResponse!")
	if command.QueryResponses.BatteryLevel != nil {
		server.log.Printf("Device name is %s and is charged to %.2f%%\n", *command.QueryResponses.DeviceName,
			*command.QueryResponses.BatteryLevel*100)
	} else {
		server.log.Printf("Device name is %s\n", *command.QueryResponses.DeviceName)
	}
	server.log.Printf("%+v", command.QueryResponses)
}

func (server *WebhookServer) handleCommandCompletedOK(command Command) {
	server.log.Printf("Command %s completed without errors!\n", command.CommandUUID)
}

/* Utils. */
