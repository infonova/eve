package events

type Httpd struct {
	Auth        string `json:"auth,omitempty"`
	Bytes       int64  `json:"bytes,omitempty"`
	ClientIp    string `json:"clientip,omitempty"`
	HttpVersion string `json:"httpversion,omitempty"`
	Ident       string `json:"ident,omitempty"`
	Request     string `json:"request,omitempty"`
	Response    string `json:"response,omitempty"`
	Verb        string `json:"verb,omitempty"`
	Origin      string `json:"origin,omitempty"`
}
