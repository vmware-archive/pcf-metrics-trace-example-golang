package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/tedsuo/rata"
)

type traceHandler struct{}

func NewTraceHandler() traceHandler {
	return traceHandler{}
}

func (t traceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	delayParam := rata.Param(r, "delay")
	delay, err := strconv.ParseFloat(delayParam, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
