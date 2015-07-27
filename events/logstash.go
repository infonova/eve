package events

import (
	"time"
)

type Logstash struct {
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	Host      string    `json:"host"`
	Path      string    `json:"path,omitempty"`
	Tags      string    `json:"tags,omitempty"`
	Timestamp time.Time `json:"@timestamp"`
}
