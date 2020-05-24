package WebhookServer

import "time"

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
