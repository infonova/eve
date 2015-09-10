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

	var status = handleEvent(logEvent, r)

	if status == http.StatusOK {
		logEvent.Totalrecall.Topic = "logs"
		asyncEventProducer.Input() <- &sarama.ProducerMessage{
			Topic: "logs",
			Value: logEvent,
		}
	}

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(status)
}

func MetricsIndex(w http.ResponseWriter, r *http.Request) {
	var writer = &loggedWriter{w, r, time.Now()}
	var metricEvent = &events.Metric{}

	var status = handleEvent(metricEvent, r)

	if status == http.StatusOK {
		metricEvent.Totalrecall.Topic = "metrics"
		asyncEventProducer.Input() <- &sarama.ProducerMessage{
			Topic: "metrics",
			Value: metricEvent,
		}
	}

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(status)
}

func TracesIndex(w http.ResponseWriter, r *http.Request) {
	var writer = &loggedWriter{w, r, time.Now()}
	var traceEvent = &events.Trace{}

	var status = handleEvent(traceEvent, r)
	if status == http.StatusOK {
		traceEvent.Totalrecall.Topic = "traces"
		asyncEventProducer.Input() <- &sarama.ProducerMessage{
			Topic: "traces",
			Value: traceEvent,
		}
	}

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(status)
}

func handleEvent(event interface{}, r *http.Request) int {
	var myEvent events.EventInterface
	switch eventType := event.(type) {
	case *events.Log:
		myEvent, _ = event.(*events.Log)
	case *events.Metric:
		myEvent, _ = event.(*events.Metric)
	case *events.Trace:
		myEvent, _ = event.(*events.Trace)
	default:
		log.Println("Unexpected type %T", eventType)
		return http.StatusBadRequest
	}

	err := json.NewDecoder(io.LimitReader(r.Body, maxLength)).Decode(&myEvent)
	if err != nil {
		log.Println("Error during decoding: " + err.Error())
		return http.StatusBadRequest
	}
	if err := myEvent.IsValid(); err != nil {
		jsonEvent, _ := json.Marshal(myEvent)
		log.Println(err)
		fmt.Println(string(jsonEvent))
		return http.StatusBadRequest
	}

	return http.StatusOK
}
