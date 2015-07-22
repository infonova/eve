package events

import "time"

type Endpoint struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Service string `json:"service"`
}

type Trace struct {
	Event
	TraceId  int64  `json:"traceId"`
	SpanId   int64  `json:"spanId"`
	Name     string `json:"name"`
	ParentId int64  `json:"parentId"`
	Tla      []struct {
		Timestamp time.Time `json:"timestamp"`
		Value     string    `json:"value"`
		Endpoint
	} `json:"tla"`
	Kva []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Endpoint
	} `json:"kva"`
}
