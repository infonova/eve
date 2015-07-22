package events

type Log struct {
	Event
	//app srv specific
	Content            string `json:"content"`
	Uuid               string `json:"content"`
	BusinessProcessKey string `json:"businessprocesskey"`
	Mandator           string `json:"mandator"`
	Class              string `json:"class"`
	Thread             string `json:"thread"`
	Module             string `json:"module"`
	LogLevel           string `json:"loglevel"`
	//httpd specific
	Auth        string `json:"auth"`
	Bytes       int64  `json:"bytes"`
	ClientIp    string `json:"clientip"`
	HttpVersion string `json:"httpversion"`
	Ident       string `json:"ident"`
	Request     string `json:"request"`
	Response    string `json:"response"`
	Verb        string `json:"verb"`
	Origin      string `json:"origin"`
}
