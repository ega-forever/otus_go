package internal

import (
	"bytes"
	"encoding/json"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type logRequest struct {
	uri     string
	payload interface{}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logRequest{uri: r.RequestURI, payload: ""}

		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			b, _ := ioutil.ReadAll(r.Body)
			rdr1 := ioutil.NopCloser(bytes.NewBuffer(b))
			rdr2 := ioutil.NopCloser(bytes.NewBuffer(b))
			_ = json.NewDecoder(rdr1).Decode(&log.payload)
			r.Body = rdr2
		}

		logger.Info(log)
		next.ServeHTTP(w, r)
	})
}
