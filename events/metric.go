package events

type Metric struct {
	Event
	Input    string `json:"input,omitempty"`
	Comp     string `json:"comp,omitempty"`
	Module   string `json:"module,omitempty"`
	Instance string `json:"instance,omitempty"`
	Metric   string `json:"metric" valid:"Required"`
	Service  string `json:"service" valid:"Required"`
}
