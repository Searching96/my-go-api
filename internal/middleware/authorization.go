package middleware

import (
	"errors"
	"net/http"

	"github.com/Searching96/my-go-api.git/api"
	"github.com/Searching96/my-go-api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid user name or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Reuqest) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHanlder(w)
			return
		}

		var loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
