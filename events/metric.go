package events

import (
	"encoding/json"
)

type Metric struct {
	Totalrecall
	Logstash
	Riemann
	Collectd

	encoded []byte
	err     error
}

func (metricEvent *Metric) ensureEncoded() {
	if metricEvent.encoded == nil && metricEvent.err == nil {
		metricEvent.encoded, metricEvent.err = json.Marshal(metricEvent)
	}
}

func (metricEvent *Metric) Length() int {
	metricEvent.ensureEncoded()
	return len(metricEvent.encoded)
}

func (metricEvent *Metric) Encode() ([]byte, error) {
	metricEvent.ensureEncoded()
	return metricEvent.encoded, metricEvent.err
}

func (metricEvent *Metric) IsValid() error {
	ev := &EventValidator{}
	ev.validateEvent(&metricEvent.Totalrecall)
	ev.validateEvent(&metricEvent.Logstash)
	ev.validateEvent(&metricEvent.Riemann)
	ev.validateEvent(&metricEvent.Collectd)

	if ev.err != nil {
		return ev.err
	}
	return nil
}
