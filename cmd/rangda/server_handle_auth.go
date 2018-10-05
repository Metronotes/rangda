package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/openware/rails5session-go"
	"github.com/openware/rangda/pkg/barong"
	"github.com/openware/rangda/pkg/log"
	"github.com/reconquest/karma-go"
)

func (server *Server) HandleAuth(
	writer http.ResponseWriter, request *http.Request,
) {
	context := karma.
		Describe("remote", request.RemoteAddr).
		Describe("request_id", middleware.GetReqID(request.Context()))

	log.Infof(context, "handling auth request")

	cookie, err := request.Cookie(CookieSession)
	if err != nil && err != http.ErrNoCookie {
		log.Errorf(err, "unable to get session cookie: %s", CookieSession)
		return
	}

	if err == http.ErrNoCookie {
		// handle case when there is no cookie
		return
	}

	log.Debugf(nil, "user's cookie: %s", cookie.Value)

	sessionData, err := rails5session.VerifyAndDecryptCookieSession(
		server.encryption,
		cookie.Value,
	)
	if err != nil {
		log.Warningf(err, "unable to verify/decrypt cookie session")
		return
	}

	var session barong.Session
	err = json.Unmarshal(sessionData, &session)
	if err != nil {
		log.Errorf(
			err,
			"unable to decode JSON session data: %s",
			string(sessionData),
		)

		return
	}

	writer.WriteHeader(http.StatusOK)
}
