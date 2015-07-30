package events

import (
	"time"
)

type Logstash struct {
	Message   string    `json:"message" valid:"Required"`
	Type      string    `json:"type,omitempty"`
	Host      string    `json:"host" valid:"Required"`
	Path      string    `json:"path,omitempty"`
	Tags      string    `json:"tags,omitempty"`
	Timestamp time.Time `json:"@timestamp" valid:"Required"`
}
