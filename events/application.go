package events

type Application struct {
	Content            string `json:"content,omitempty"`
	Uuid               string `json:"uuid,omitempty"`
	BusinessProcessKey string `json:"businessprocesskey,omitempty"`
	Mandator           string `json:"mandator,omitempty"`
	Class              string `json:"class,omitempty"`
	Thread             string `json:"thread,omitempty"`
	Module             string `json:"module,omitempty"`
	LogLevel           string `json:"loglevel,omitempty"`
}
