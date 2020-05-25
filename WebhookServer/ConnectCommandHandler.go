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
		server.handleCommandAcknowledgement(*sparse)
		break
	case IDLE, NOT_NOW:
		break
	case COMMAND_FORMAT_ERROR:
		server.log.Printf("COMMAND_FORMAT_ERROR! " + fmt.Sprint(sparse.ErrorChain))
	default:
		server.log.Printf("Not handling state " + sparse.Status)
	}
}

func (server *WebhookServer) handleCommandAcknowledgement(command Command) {
	if command.QueryResponses != nil {
		server.handleQueryResponse(command)
	} else {
		server.log.Printf("Not handling command ACK")
		fmt.Printf("%+v", command)
	}
}

func (server *WebhookServer) handleQueryResponse(command Command) {
	server.log.Println("Got QueryResponse!")

	server.log.Printf("%s is at %.2f%%", *command.QueryResponses.DeviceName, (*command.QueryResponses.BatteryLevel)*100)

	//fmt.Printf("%+v", command.QueryResponses)
}

/* Utils. */
