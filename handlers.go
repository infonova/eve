package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"git.infonova.at/totalrecall/eve/events"
	"github.com/Shopify/sarama"
)

const maxLength int64 = 1024 * 512

func LogsIndex(w http.ResponseWriter, r *http.Request) {
	var content = &events.Log{}
	err := json.NewDecoder(io.LimitReader(r.Body, maxLength)).Decode(&content)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	asyncEventProducer.Input() <- &sarama.ProducerMessage{
		Topic: "logs",
		Value: content,
	}

	w.WriteHeader(http.StatusOK)
}

func MetricsIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Metrics endpoint")
}

func TracesIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Traces endpoint")
}
