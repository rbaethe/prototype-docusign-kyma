package events

import (
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/config"
	"strings"
)


type Recipient struct {
	email string `xml:"Email", json:"email"`
	recipienttype     string `xml:"Type", json:"type"`
}

type CloudEvent struct {
	EventType        string      `json:"type"`
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
	eventType := *config.GlobalConfig.BaseTopic + "." + "envelope" + "." + strings.ToLower(envelope.EnvelopeStatus.Status)

	return &CloudEvent{
		EventType:        eventType,
		Data:             envelope.EnvelopeStatus,
	}
}
