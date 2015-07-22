package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	ip            = flag.String("ip", "0.0.0.0", "ip address (default 0.0.0.0")
	port          = flag.String("port", "8080", "port (default 8080)")
	kafka_connect = flag.String("kafka_connect", "localhost:9092", "kafka broker list (default localhost:9092")
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(*ip+":"+*port, router))
}
