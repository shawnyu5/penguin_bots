package logger

import (
	"log"
	"net/http"
	"time"
)

// a logger middleware
type LoggerMidware struct {
	log  log.Logger
	next http.Handler
}

func (mw LoggerMidware) logger() {
	defer func(begin time.Time) {

	}(time.Now())
}
