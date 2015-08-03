package events

import (
	"encoding/json"
)

type Log struct {
	Totalrecall
	Logstash
	Httpd
	Application

	encoded []byte
	err     error
}

func (logEvent *Log) ensureEncoded() {
	if logEvent.encoded == nil && logEvent.err == nil {
		logEvent.encoded, logEvent.err = json.Marshal(logEvent)
	}
}

func (logEvent *Log) Length() int {
	logEvent.ensureEncoded()
	return len(logEvent.encoded)
}

func (logEvent *Log) Encode() ([]byte, error) {
	logEvent.ensureEncoded()
	return logEvent.encoded, logEvent.err
}

func (logEvent *Log) IsValid() error {
	ev := &EventValidator{}
	ev.validateEvent(&logEvent.Totalrecall)
	ev.validateEvent(&logEvent.Logstash)
	if ev.err != nil {
		return ev.err
	}
	return nil
}
