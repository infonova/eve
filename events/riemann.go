package events

type Riemann struct {
	Service     string  `json:"service" valid:"Required"`
	State       string  `json:"state,omitempty"`
	Time        uint64  `json:"time,omitempty"`
	Description string  `json:"description,omitempty"`
	Metric      float64 `json:"metric" valid:"Required"`
	Ttl         uint64  `json:"ttl,omitempty"`
}
