package events

import (
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/config"
	"strings"
	"time"
)
import "github.com/gofrs/uuid"

type Envelope struct {
	EnvelopeId string `xml:"EnvelopeStatus>EnvelopeID", json:"EnvelopeID"`
	Status     string `xml:"EnvelopeStatus>Status", json:"status"`
	StatusTime string `xml:"EnvelopeStatus>TimeGenerated", json:"time"`
}

type KymaEvent struct {
	SourceID         *string     `json:"source-id"`
	EventType        string      `json:"event-type"`
	EventTypeVersion string      `json:"event-type-version"`
	EventID          string      `json:"event-id"`
	EventTime        string      `json:"event-time"`
	Data             interface{} `json:"data"`
}

func Map(envelope *Envelope) *KymaEvent {
	eventId, _ := generateEventID()
	eventType := *config.GlobalConfig.BaseTopic + "." + "envelope" + "." + strings.ToLower(envelope.Status)

	return &KymaEvent{
		SourceID:         config.GlobalConfig.AppName,
		EventType:        eventType,
		EventTypeVersion: "v1",
		EventID:          eventId,
		EventTime:        time.Now().Format(time.RFC3339),
		Data:             envelope,
	}
}

func generateEventID() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
