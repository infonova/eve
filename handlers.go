package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"git.infonova.at/totalrecall/eve/events"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/prometheus/storage/remote"
)

const maxLength int64 = 1024 * 512

type loggedWriter struct {
	w http.ResponseWriter
	r *http.Request
	t time.Time
}

func (w *loggedWriter) WriteHeader(status int) {
	w.w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.w.WriteHeader(status)
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

func JsonLogsIndex(w http.ResponseWriter, r *http.Request) {
	var writer = &loggedWriter{w, r, time.Now()}
	var logEvent = &events.Log{}

	var status = handleJsonEvent(logEvent, r)

	if status == http.StatusOK {
		logEvent.Totalrecall.Topic = "logs"
		asyncEventProducer.Input() <- &sarama.ProducerMessage{
			Topic: "logs",
			Value: logEvent,
		}
	}

	writer.WriteHeader(status)
}

func JsonMetricsIndex(w http.ResponseWriter, r *http.Request) {
	var writer = &loggedWriter{w, r, time.Now()}
	var metricEvent = &events.Metric{}

	var status = handleJsonEvent(metricEvent, r)

	if status == http.StatusOK {
		metricEvent.Totalrecall.Topic = "metrics"
		asyncEventProducer.Input() <- &sarama.ProducerMessage{
			Topic: "metrics",
			Value: metricEvent,
		}
	}

	writer.WriteHeader(status)
}

func JsonTracesIndex(w http.ResponseWriter, r *http.Request) {
	var writer = &loggedWriter{w, r, time.Now()}
	var traceEvent = &events.Trace{}

	var status = handleJsonEvent(traceEvent, r)
	if status == http.StatusOK {
		traceEvent.Totalrecall.Topic = "traces"
		asyncEventProducer.Input() <- &sarama.ProducerMessage{
			Topic: "traces",
			Value: traceEvent,
		}
	}

	writer.WriteHeader(status)
}

func handleJsonEvent(event interface{}, r *http.Request) int {
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

func PrometheusIndex(w http.ResponseWriter, r *http.Request) {
	reqBuf, err := ioutil.ReadAll(snappy.NewReader(r.Body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req remote.WriteRequest
	if err := proto.Unmarshal(reqBuf, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, ts := range req.Timeseries {
		jsonmap := map[string]interface{}{}

		for _, l := range ts.Labels {
			jsonmap[l.Name] = l.Value
		}

		for _, s := range ts.Samples {
			jsonmap["value"] = s.Value
			jsonmap["timestamp"] = s.TimestampMs
		}

		b, err := json.Marshal(jsonmap)
		if err != nil {
			fmt.Println(jsonmap)
			fmt.Println(err)
			return
		}

		asyncEventProducer.Input() <- &sarama.ProducerMessage{
			Topic: "prometheus",
			Value: sarama.StringEncoder(b),
		}
	}
}
