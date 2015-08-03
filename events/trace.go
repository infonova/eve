package events

import (
	"encoding/json"
	"time"
)

type EndpointDefinition struct {
	Host    string `json:"host" valid:"Required"`
	Port    int    `json:"port" valid:"Required"`
	Service string `json:"service" valid:"Required"`
}

type Trace struct {
	Totalrecall
	Traceid   int64     `json:"traceid" valid:"Required"`
	Spanid    int64     `json:"spanid" valid:"Required"`
	Name      string    `json:"name" valid:"Required"`
	Parentid  int64     `json:"parentid" valid:"Required"`
	Timestamp time.Time `json:"@timestamp" valid:"Required"`
	Tla       []struct {
		Timestamp int64  `json:"timestamp" valid:"Required"`
		Value     string `json:"value" valid:"Required"`
		Endpoint  struct {
			EndpointDefinition
		} `json:"endpoint" valid:"Required"`
	} `json:"tla" valid:"Required"`
	Kva []struct {
		Key      string `json:"key" valid:"Required"`
		Value    string `json:"value" valid:"Required"`
		Endpoint struct {
			EndpointDefinition
		} `json:"endpoint,omitempty"`
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
	ev.validateEvent(&traceEvent.Totalrecall)
	ev.validateEvent(traceEvent)
	for _, item := range traceEvent.Tla {
		ev.validateEvent(item)
		ev.validateEvent(&item.Endpoint.EndpointDefinition)
	}
	for _, item := range traceEvent.Kva {
		ev.validateEvent(item)
		ev.validateEvent(&item.Endpoint.EndpointDefinition)
	}
	if ev.err != nil {
		return ev.err
	}
	return nil
}
