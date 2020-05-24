package WebhookServer

import (
	"howett.net/plist"
	"time"
)

type Event struct {
	Topic     string    `json:"topic"`
	EventID   string    `json:"event_id"`
	CreatedAt time.Time `json:"created_at"`

	AcknowledgeEvent *AcknowledgeEvent `json:"acknowledge_event,omitempty"`
	CheckinEvent     *CheckinEvent     `json:"checkin_event,omitempty"`
}

type AcknowledgeEvent struct {
	UDID        string            `json:"udid"`
	Status      string            `json:"status"`
	CommandUUID string            `json:"command_uuid"`
	Params      map[string]string `json:"url_params"`
	RawPayload  []byte            `json:"raw_payload"`
}

type CheckinEvent struct {
	UDID       string            `json:"udid"`
	Params     map[string]string `json:"url_params"`
	RawPayload []byte            `json:"raw_payload"`
}

type AcknowledgeEventState int

const (
	ACKNOWLEDGED AcknowledgeEventState = iota
	ERROR
	COMMAND_FORMAT_ERROR
	IDLE
	NOT_NOW
	UNKNOWN
)

func (s AcknowledgeEventState) String() string {
	return [...]string{"ACKNOWLEDGED", "ERROR", "COMMAND_FORMAT_ERROR", "IDLE", "NOT_NOW", "UNKNOWN"}[s]
}

func acknowledgeStateStringToEnum(state string) AcknowledgeEventState {
	switch state {
	case "Acknowledged":
		return ACKNOWLEDGED
	case "CommandFormatError":
		return COMMAND_FORMAT_ERROR
	case "Error":
		return ERROR
	case "NotNow":
		return NOT_NOW
	case "Idle":
		return IDLE
	default:
		activeServer.log.Printf("Found known state: " + state)
		return UNKNOWN
	}
}

type AcknowledgeSparse struct {
	UDID        string
	CommandUUID string
	Status      AcknowledgeEventState
	AllKeys     map[string]interface{}
}

func parseAcknowledge(data []byte) *AcknowledgeSparse {
	mapData := decodePLIST(data)
	if mapData == nil {
		return nil
	}

	udid := mapStringSafeGet(mapData, "UDID")
	if udid == nil {
		activeServer.log.Printf("Could not extract UDID")
		return nil
	}

	var commandUUID string
	commandUUIDPtr := mapStringSafeGet(mapData, "CommandUUID")
	if commandUUIDPtr == nil {
		//certain commands will not have a UUID, stuff it with DNE uuid
		commandUUID = "AraMDM__INTERNAL__COMMAND_NOT_STAPLED"
	} else {
		commandUUID = *commandUUIDPtr
	}

	status := mapStringSafeGet(mapData, "Status")
	if status == nil {
		activeServer.log.Printf("Could not extract Status")
		return nil
	}

	return &AcknowledgeSparse{
		UDID:        *udid,
		CommandUUID: commandUUID,
		Status:      acknowledgeStateStringToEnum(*status),
		AllKeys:     mapData,
	}
}

/* Attempts to lookup a key for a string value. Returns nil if not found or not a string. */
func mapStringSafeGet(source map[string]interface{}, key string) *string {
	buf, ok := source[key]
	if !ok {
		return nil
	}

	str, ok := buf.(string)
	if !ok {
		return nil
	}

	return &str
}

/* Generically decode a PLIST. Returns nil on failure. */
func decodePLIST(data []byte) map[string]interface{} {
	var obj map[string]interface{}
	_, err := plist.Unmarshal(data, &obj)
	if err != nil {
		activeServer.log.Printf("Unable to decode PLIST: " + err.Error())
		return nil
	}

	return obj
}
