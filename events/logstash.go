package events

import (
	"time"
)

type Logstash struct {
	Host      string    `json:"host" valid:"Required"`
	Timestamp time.Time `json:"@timestamp" valid:"Required"`
	Message   string    `json:"message,omitempty"`
	Type      string    `json:"type,omitempty"`
	Path      string    `json:"path,omitempty"`
	Tags      string    `json:"tags,omitempty"`
}
