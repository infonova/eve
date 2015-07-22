package events

type Metric struct {
	Event
	Input    string `json:"input"`
	Comp     string `json:"comp"`
	Module   string `json:"module"`
	Instance string `json:"instance"`
	Metric   string `json:"metric"`
	Service  string `json:"service"`
}
