package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type okResponse struct {
	Ok int `json:"ok"`
}

type RestService struct {
	srv *http.Server
}

func NewRestService(port int) *RestService {
	r := mux.NewRouter()
	r.HandleFunc("/", handleRootEndpoint)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + strconv.Itoa(port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &RestService{srv: srv}
}

func handleRootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ok := okResponse{Ok: 1}
	_ = json.NewEncoder(w).Encode(&ok)

}

func (rs *RestService) Start() error {
	return rs.srv.ListenAndServe()
}

func (rs *RestService) Disconnect() error {
	return rs.srv.Close()
}
