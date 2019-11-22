package handlers

import (
	"crypto/subtle"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/config"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/model/errors"
	"net/http"
)

func AuthenticationHandler(fn http.HandlerFunc) http.HandlerFunc {
	passwordString := *config.GlobalConfig.Passwd
	usernameString := *config.GlobalConfig.UserName

	return func(writer http.ResponseWriter, request *http.Request) {
		user, pass, ok := request.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(usernameString)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(passwordString)) != 1 {
			errors.HandleError(writer, nil, errors.UnAuthorized)
			return
		}
		fn(writer, request)
	}
}
