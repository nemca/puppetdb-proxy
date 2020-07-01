package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func (s *server) logHTTP(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := statusWriter{ResponseWriter: w}
		handler.ServeHTTP(&sw, r)
		duration := time.Now().Sub(start)
		// addr := strings.Split(r.RemoteAddr, ":")[0] // TODO: add parsing IPv6 address and port
		s.Log.Infof("%s \"%s %s %s\" %d %f", r.RemoteAddr, r.Method, r.RequestURI, r.Proto, sw.status, duration.Seconds())
	})
}

func (s *server) initLogger() {
	s.Log = log.New()
	s.Log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile(opts.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		s.Log.Fatal(err)
	}
	s.Log.Out = file
	s.Log.SetLevel(log.Level(opts.LogLevel))
}
