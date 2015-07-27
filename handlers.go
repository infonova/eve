package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"git.infonova.at/totalrecall/eve/events"
	"github.com/Shopify/sarama"
)

const maxLength int64 = 1024 * 512

func LogsIndex(w http.ResponseWriter, r *http.Request) {
	var logEvent = &events.Log{}

	err := json.NewDecoder(io.LimitReader(r.Body, maxLength)).Decode(&logEvent)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		log.Println("Error during decoding: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := logEvent.IsValid(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonEvent, _ := json.Marshal(logEvent)
		log.Println(err)
		fmt.Println(string(jsonEvent))
		return
	}

	asyncEventProducer.Input() <- &sarama.ProducerMessage{
		Topic: "logs",
		Value: logEvent,
	}

	w.WriteHeader(http.StatusOK)
}

func MetricsIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Metrics endpoint")
}

func TracesIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Traces endpoint")
}
