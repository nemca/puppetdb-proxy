// server.go
package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type server struct {
	Router *mux.Router
	Log    *log.Logger
}

func newServer() *server {
	s := new(server)

	s.Router = mux.NewRouter()
	s.Router.Use(s.logHTTP)
	s.Router.Use(s.metricsMiddleware)
	s.initRoutes()

	s.initLogger()

	return s
}

func (s *server) run(addr string) {
	s.Log.Infof("Run server on a %s", addr)
	s.Log.Fatal(http.ListenAndServe(addr, s.Router))
}
