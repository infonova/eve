package main

import "time"

type Log struct {
	Message   string    `json:"message"`
	ProjectId string    `json:"projectid"`
	TargetId  string    `json:"targetid"`
	Timestamp time.Time `json:"@timestamp"`
}
