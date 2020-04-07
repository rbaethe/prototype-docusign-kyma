package events

import (
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/config"
	"strings"
	"time"
)
import "github.com/gofrs/uuid"

/*
type Envelope struct {
	EnvelopeId string `xml:"EnvelopeStatus>EnvelopeID", json:"EnvelopeID"`
	Status     string `xml:"EnvelopeStatus>Status", json:"status"`
	StatusTime string `xml:"EnvelopeStatus>TimeGenerated", json:"time"`
	Recipients [] Recipient `xml:"EnvelopeStatus>RecipientStatuses->RecipientStatus", json:"RecipientStatus"`
}

*/


type Recipient struct {
	email string `xml:"Email", json:"email"`
	recipienttype     string `xml:"Type", json:"type"`
}

type KymaEvent struct {
	SourceID         *string     `json:"source-id"`
	EventType        string      `json:"event-type"`
	EventTypeVersion string      `json:"event-type-version"`
	EventID          string      `json:"event-id"`
	EventTime        string      `json:"event-time"`
	Data             interface{} `json:"data"`
}

type DocuSignEnvelopeInformation struct {
	EnvelopeStatus struct {
		Status           string `xml:"Status"`
		StatusTime       string `xml:"TimeGenerated"`
		EnvelopeID       string `xml:"EnvelopeID"`
		RecipientStatuses struct {
			RecipientStatus struct {
				Type  string `xml:"Type"`
				Email string `xml:"Email"`
				} `xml:"RecipientStatus"`
			} `xml:"RecipientStatuses"`
			DocumentStatuses struct {
				DocumentStatus struct {
					ID           string `xml:"ID"`
					Name         string `xml:"Name"`
					TemplateName string `xml:"TemplateName"`
				} `xml:"DocumentStatus"`
			} `xml:"DocumentStatuses"`
		} `xml:"EnvelopeStatus"`
}



func Map(envelope *DocuSignEnvelopeInformation) *KymaEvent {
	eventId, _ := generateEventID()
	eventType := *config.GlobalConfig.BaseTopic + "." + "envelope" + "." + strings.ToLower(envelope.EnvelopeStatus.Status)

	return &KymaEvent{
		SourceID:         config.GlobalConfig.AppName,
		EventType:        eventType,
		EventTypeVersion: "v1",
		EventID:          eventId,
		EventTime:        time.Now().Format(time.RFC3339),
		Data:             envelope.EnvelopeStatus,
	}
}


func generateEventID() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
