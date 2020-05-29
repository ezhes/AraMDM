package WebhookServer

import "fmt"

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
	} else {
		return false
	}

	return true
}

func (server *WebhookServer) handleQueryResponse(command Command) {
	server.log.Println("Got QueryResponse!")
	server.log.Printf("Device name is %s\n", *command.QueryResponses.DeviceName)

	//server.log.Printf("%s is at %.2f%%", *command.QueryResponses.DeviceName, (*command.QueryResponses.BatteryLevel)*100)

	fmt.Printf("%+v", command.QueryResponses)
}

/* Utils. */
