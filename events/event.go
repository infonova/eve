package events

import (
	"errors"

	"github.com/astaxie/beego/validation"
)

type Event struct {
	ProjectId   string `json:"projectid" valid:"Required"`
	TargetId    string `json:"targetid" valid:"Required"`
	Application string `json:"application,omitempty"`
	Config      string `json:"config,omitempty"`
}

type EventValidator struct {
	err error
}

func (ev *EventValidator) validateEvent(event interface{}) {
	if ev.err != nil {
		return
	}

	valid := validation.Validation{}
	b, err := valid.Valid(event)
	if err != nil {
		ev.err = err
	}

	if !b {
		ev.err = errors.New("Invalid event")
	}
}
