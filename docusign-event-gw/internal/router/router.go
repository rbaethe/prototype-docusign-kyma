package router

import (
	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/config"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/handlers"
	"net/http"
	"os"
)

type Rtr struct {
	http.Handler
	*handlers.EventPublisher
}

func New() http.Handler {
	r := mux.NewRouter()
	ep := handlers.NewEventPublisher()

	eventsHandler := applyBasicAuth(ep.EventHandler())
	r.HandleFunc("/events", eventsHandler)

	return &Rtr{
		Handler:        applyLogging(r),
		EventPublisher: ep,
	}
}

func applyLogging(r *mux.Router) http.Handler {
	if !*config.GlobalConfig.LogRequest {
		return r
	} else {
		return gh.LoggingHandler(os.Stdout, r)
	}
}

func applyBasicAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	if *config.GlobalConfig.BasicAuthEnabled {
		return handlers.AuthenticationHandler(handlerFunc)
	} else {
		return handlerFunc
	}
}
