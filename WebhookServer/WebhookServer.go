package WebhookServer

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"strings"
)

type WebhookServer struct {
	log                   *log.Logger
	router                *httprouter.Router
	expectedSourceAddress string
}

var activeServer *WebhookServer = nil

func New(expectedSourceAddress string, loggerFlags int) *WebhookServer {
	server := new(WebhookServer)

	server.expectedSourceAddress = expectedSourceAddress

	server.log = log.New(os.Stderr, "[Webhooks] ", loggerFlags)

	server.router = httprouter.New()
	server.router.POST("/ara-hook", mdmWebhook)

	//POST INIT -- all fields must be created

	return server
}

/*
Starts the server in the background
*/
func (server *WebhookServer) Launch(port int) {
	if activeServer != nil {
		server.log.Fatal("Attempted to start duplicate webhook server instance!")
	}

	activeServer = server

	go func() {
		server.log.Printf("Starting on port %v", port)
		server.log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), server.router))
	}()
}

func mdmWebhook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !activeServer.isWebhookRequesterValid(r.RemoteAddr) {
		w.WriteHeader(401)
		write(w, "Get lost, scoundrel: you are not microMDM! If you ACTUALLY are microMDM, ensure Configuration has the correct requesting address!\nRequest came from "+r.RemoteAddr+"\n")
	}

	decoder := json.NewDecoder(r.Body)
	var event Event
	err := decoder.Decode(&event)
	if err != nil {
		activeServer.log.Printf("Event decode error: " + err.Error())
		w.WriteHeader(500)
		write(w, "Decode error. See logs.")
		return
	}

	// Release the server
	write(w, "OK")

	//pass it on
	go activeServer.handleEvent(event)
}

func (server *WebhookServer) isWebhookRequesterValid(remoteAddr string) bool {
	requestParts := strings.Split(remoteAddr, ":")
	if len(requestParts) != 2 {
		server.log.Printf("Unable to parse remote addr " + remoteAddr)
		return false
	}

	return server.expectedSourceAddress == requestParts[0]
}

func write(w http.ResponseWriter, value string) {
	_, err := fmt.Fprint(w, value)
	if err != nil {
		log.Println("Write error:" + err.Error())
	}
}
