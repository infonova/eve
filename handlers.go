package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"git.infonova.at/totalrecall/eve/events"
	"github.com/Shopify/sarama"
)

const maxLength int64 = 1024 * 512

type loggedWriter struct {
	w http.ResponseWriter
	r *http.Request
	t time.Time
}

func (w *loggedWriter) WriteHeader(status int) {
	log.Printf(
		"%s\t%s\t%s\t%s\t%d\t%s",
		w.r.RemoteAddr,
		w.r.Method,
		w.r.RequestURI,
		w.r.Proto,
		status,
		time.Since(w.t),
	)
}

func (w *loggedWriter) Header() http.Header { return w.w.Header() }

func LogsIndex(w http.ResponseWriter, r *http.Request) {
	var writer = &loggedWriter{w, r, time.Now()}
	var logEvent = &events.Log{}

	err := json.NewDecoder(io.LimitReader(r.Body, maxLength)).Decode(&logEvent)
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		log.Println("Error during decoding: " + err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := logEvent.IsValid(); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		jsonEvent, _ := json.Marshal(logEvent)
		log.Println(err)
		fmt.Println(string(jsonEvent))
		return
	}

	asyncEventProducer.Input() <- &sarama.ProducerMessage{
		Topic: "logs",
		Value: logEvent,
	}

	writer.WriteHeader(http.StatusOK)
}

func MetricsIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Metrics endpoint")
}

func TracesIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Traces endpoint")
}
