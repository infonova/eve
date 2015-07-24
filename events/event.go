package events

import (
	"time"
)

type Event struct {
	//totalrecall specific fields
	ProjectId   string `json:"projectid"`
	TargetId    string `json:"targetid"`
	Application string `json:"application"`
	Config      string `json:"config"`
	//eventstash specific fields
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	Host      string    `json:"host"`
	Path      string    `json:"path"`
	Tags      string    `json:"tags"`
	Timestamp time.Time `json:"@timestamp"`

	encoded []byte
	err     error
}
