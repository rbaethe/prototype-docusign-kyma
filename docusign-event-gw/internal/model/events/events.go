package events

import (
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/config"
	"strings"
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

type CloudEvent struct {
	EventType        string      `json:"event-type"`
	EventSpecVersion string      `json:"specversion"`
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



func Map(envelope *DocuSignEnvelopeInformation) *CloudEvent {
	//eventId, _ := generateEventID()
	eventType := *config.GlobalConfig.BaseTopic + "." + "envelope" + "." + strings.ToLower(envelope.EnvelopeStatus.Status)

	return &CloudEvent{
		EventType:        eventType,
		EventSpecVersion: "1.0",
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
