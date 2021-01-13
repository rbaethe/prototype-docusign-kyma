package outgoing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/config"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/logger"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/model/events"
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

func (e *EventForwarder) Forward(event *events.CloudEvent) (map[string]interface{}, error) {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, *e.eventPublishURL, bytes.NewReader(eventBytes))
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var respMap map[string]interface{}
	err = dec.Decode(&respMap)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("unexpected response when publishing event %d (%s)", resp.StatusCode, resp.Status)
		logger.Logger.Error(errMsg)
		return respMap, fmt.Errorf(errMsg)
	}

	return respMap, nil

}

func (e *EventForwarder) enrichRequest(event *events.CloudEvent, request *http.Request) *http.Request {

	eventId, _ := generateEventID()
	//sourceID := (config.GlobalConfig.AppName).String()

	request.Header.Set("Content-Type", "application/cloudevents+json; charset=utf-8")
	request.Header.Set("ce-specversion", "1.0")
	request.Header.Set("ce-type", event.EventType)
	request.Header.Set("ce-eventtypeversion", "v1")
	request.Header.Set("ce-id", eventId)
	request.Header.Set("ce-source", "docusign")
	return request
}

func generateEventID() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
