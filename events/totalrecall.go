package events

type Totalrecall struct {
	ProjectId   string `json:"projectid" valid:"Required"`
	TargetId    string `json:"targetid" valid:"Required"`
	Application string `json:"application,omitempty"`
	Config      string `json:"config,omitempty"`
}
