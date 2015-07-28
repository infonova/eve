package events

import (
	"encoding/json"
	"time"
)

type Endpoint struct {
	Host    string `json:"host" valid:"Required"`
	Port    int    `json:"port" valid:"Required"`
	Service string `json:"service" valid:"Required"`
}

type Trace struct {
	Event
	TraceId  int64  `json:"traceId" valid:"Required"`
	SpanId   int64  `json:"spanId" valid:"Required"`
	Name     string `json:"name" valid:"Required"`
	ParentId int64  `json:"parentId" valid:"Required"`
	Tla      []struct {
		Timestamp time.Time `json:"timestamp" valid:"Required"`
		Value     string    `json:"value" valid:"Required"`
		Endpoint
	} `json:"tla" valid:"Required"`
	Kva []struct {
		Key   string `json:"key" valid:"Required"`
		Value string `json:"value" valid:"Required"`
		Endpoint
	} `json:"kva,omitempty"`

	encoded []byte
	err     error
}

func (traceEvent *Trace) ensureEncoded() {
	if traceEvent.encoded == nil && traceEvent.err == nil {
		traceEvent.encoded, traceEvent.err = json.Marshal(traceEvent)
	}
}

func (traceEvent *Trace) Length() int {
	traceEvent.ensureEncoded()
	return len(traceEvent.encoded)
}

func (traceEvent *Trace) Encode() ([]byte, error) {
	traceEvent.ensureEncoded()
	return traceEvent.encoded, traceEvent.err
}

func (traceEvent *Trace) IsValid() error {
	ev := &EventValidator{}
	ev.validateEvent(&traceEvent.Event)
	ev.validateEvent(&traceEvent)
	for _, item := range traceEvent.Tla {
		ev.validateEvent(&item)
	}
	for _, item := range traceEvent.Kva {
		ev.validateEvent(&item)
	}
	if ev.err != nil {
		return ev.err
	}
	return nil
}
