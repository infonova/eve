package events

type Riemann struct {
	Metric      float64 `json:"metric" valid:"Required"`
	Service     string  `json:"service,omitempty"`
	State       string  `json:"state,omitempty"`
	Time        uint64  `json:"time,omitempty"`
	Description string  `json:"description,omitempty"`
	Ttl         uint64  `json:"ttl,omitempty"`
}
