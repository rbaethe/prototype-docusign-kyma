package handlers

import (
	"fmt"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/incoming"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/logger"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/model/errors"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/outgoing"
	"io/ioutil"
	"net/http"
)

type EventPublisher struct {
	eventForwarder *outgoing.EventForwarder
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{eventForwarder: outgoing.NewEventForwarder()}
}

func (ep *EventPublisher) EventHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			logger.Logger.Errorw("error when parsing request", "error", err)
		}

		kymaEvent, err := incoming.Process(body)
		if err != nil {
			errors.HandleError(w, err, errors.InternalError)
			return
		}

		logger.Logger.Infow("kyma event request body", "event",  (kymaEvent.Data))

		respCode, err := ep.eventForwarder.Forward(kymaEvent)
		if err != nil {
			errors.HandleError(w, err, errors.InternalError)
			return
		}

		infoMsg := fmt.Sprintf("Received response for event publishing: %d ", respCode)
		logger.Logger.Infow(infoMsg)

		w.WriteHeader(http.StatusOK)
	})
}
