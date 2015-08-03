package events

type Collectd struct {
	Measurement string `json:"measurement" valid:"Required"`
	Name        string `json:"mname,omitempty"`
	Type        string `json:"mtype,omitempty"`
	Instance    string `json:"instance,omitempty"`
	Unit        string `json:"unit,omitempty"`
	Request     string `json:"request,omitempty"`
}
