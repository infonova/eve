package events

import (
	"errors"
	"log"

	"github.com/astaxie/beego/validation"
)

type EventInterface interface {
	IsValid() error
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
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}

		ev.err = errors.New("Invalid event")
	}
}
