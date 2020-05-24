package WebhookServer

import "fmt"

func (server *WebhookServer) handleEvent(event Event) {
	switch event.Topic {
	case "mdm.Authenticate":
		server.handleAuthenticate(event)
		break
	case "mdm.TokenUpdate":
		server.handleTokenUpdate(event)
		break
	case "mdm.CheckOut":
		server.handleCheckOut(event)
		break
	case "mdm.Connect":
		server.handleConnect(event)
		break
	default:
		server.log.Printf("Unknown event: " + event.Topic)
		break
	}
}

/* Checkin Events. */
func (server *WebhookServer) handleAuthenticate(event Event) {
	server.log.Fatal("NOT IMPL")
}

func (server *WebhookServer) handleTokenUpdate(event Event) {
	server.log.Fatal("NOT IMPL")
}

func (server *WebhookServer) handleCheckOut(event Event) {
	server.log.Fatal("NOT IMPL")
}

/* ACK events. This includes true ACKs and command responses */
func (server *WebhookServer) handleConnect(event Event) {
	//s := string(event.AcknowledgeEvent.RawPayload)
	//println(s)
	sparse := parseAcknowledge(event.AcknowledgeEvent.RawPayload)
	if sparse == nil {
		return
	}

	switch sparse.Status {
	case ACKNOWLEDGED:
		server.handleCommandAcknowledgement(*sparse)
		break
	case IDLE, NOT_NOW:
		break
	default:
		server.log.Printf("Not handling state " + sparse.Status.String())
	}
}

func (server *WebhookServer) handleCommandAcknowledgement(sparse AcknowledgeSparse) {
	fmt.Printf("%+v", sparse.AllKeys)

}

/* Utils. */
