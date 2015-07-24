package events

import "encoding/json"

type Log struct {
	Event
	//app srv specific
	Content            string `json:"content"`
	Uuid               string `json:"uuid"`
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

func (log *Log) ensureEncoded() {
	if log.encoded == nil && log.err == nil {
		log.encoded, log.err = json.Marshal(log)
	}
}

func (log *Log) Length() int {
	log.ensureEncoded()
	return len(log.encoded)
}

func (log *Log) Encode() ([]byte, error) {
	log.ensureEncoded()
	return log.encoded, log.err
}
