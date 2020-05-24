package WebhookServer

import (
	"bytes"
	"howett.net/plist"
)

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
	s := string(event.AcknowledgeEvent.RawPayload)
	println(s)
	//data := decodePLIST(event.AcknowledgeEvent.RawPayload)
	//fmt.Printf("%+v", data)
}

/* Utils. */
func decodePLIST(data []byte) interface{} {
	buf := bytes.NewReader(data)
	var obj interface{}
	err := plist.NewDecoder(buf).Decode(obj)
	if err != nil {
		return nil
	}

	return obj
}
