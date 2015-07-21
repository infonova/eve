package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const maxLength int64 = 1024 * 512

func LogsIndex(w http.ResponseWriter, r *http.Request) {
	var content = &Log{}
	err := json.NewDecoder(io.LimitReader(r.Body, maxLength)).Decode(&content)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, content.Message+"\n", content.ProjectId+"\n", content.TargetId+"\n", content.Timestamp)
}

func MetricsIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Metrics endpoint")
}

func TracesIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Traces endpoint")
}
