package outgoing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/config"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/logger"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/model/events"
	"log"
	"net/http"
	"time"
)

import "github.com/gofrs/uuid"

type EventForwarder struct {
	eventPublishURL *string
	client          *http.Client
}

func NewEventForwarder() *EventForwarder {
	client := &http.Client{
		Transport: &http.Transport{
			DisableCompression:  false,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 0,
			MaxConnsPerHost:     0,
			IdleConnTimeout:     30 * time.Second,
		},
	}

	return &EventForwarder{
		eventPublishURL: config.GlobalConfig.EventPublishURL,
		client:          client,
	}
}

func (e *EventForwarder) Forward(event *events.CloudEvent) (int, error) {
	eventBytes, err := json.Marshal(event.Data)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, *e.eventPublishURL, bytes.NewReader(eventBytes))
	if err != nil {
		return 0, err
	}


	if err != nil {
		log.Fatal(err)
	}


	req = e.enrichRequest(event, req)

	for name, values := range req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			logger.Logger.Infow("event request header: ", string(name), string(value))
		}
	}

	resp, err := e.client.Do(req)
	if err != nil {
		logger.Logger.Infow("error received: ", err)
		return 0, err
	}

	logger.Logger.Infow("Received response without error");

	defer resp.Body.Close()

	infoMsg := fmt.Sprintf("Received response from event gateway with http status %d (%s) ", resp.StatusCode, resp.Status)
	logger.Logger.Infow(infoMsg)

	logger.Logger.Infow("Start decoding response");


	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("unexpected response when publishing event %d (%s)", resp.StatusCode, resp.Status)
		logger.Logger.Error(errMsg)
		return resp.StatusCode, fmt.Errorf(errMsg)
	}

	return resp.StatusCode, nil

}

func (e *EventForwarder) enrichRequest(event *events.CloudEvent, request *http.Request) *http.Request {

	eventId, _ := generateEventID()
	sourceID := *config.GlobalConfig.AppName

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("ce-specversion", "1.0")
	request.Header.Set("ce-type", event.EventType)
	request.Header.Set("ce-eventtypeversion", "v1")
	request.Header.Set("ce-id", eventId)
	request.Header.Set("ce-source", sourceID)
	return request
}

func generateEventID() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
