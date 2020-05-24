package WebhookServer

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
