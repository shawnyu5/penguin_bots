package middleware

import (
	"log"
	"net/http"
	"time"
)

type LoggerInter interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// logger middleware
type Logger struct {
	Next func(w http.ResponseWriter, r *http.Request)
}

// ServeHTTP wrapper around a http handler with logging
func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Next(w, r)
	log.Printf("method=%s path=%s took=%s", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger returns a new logger middleware
func NewLogger(handlerToWrap func(w http.ResponseWriter, r *http.Request)) Logger {
	return Logger{Next: handlerToWrap}

}
