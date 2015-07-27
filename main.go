package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	ip                 = flag.String("ip", "0.0.0.0", "ip address")
	port               = flag.String("port", "8080", "port")
	kafkaConnect       = flag.String("kafka_connect", "localhost:9092", "kafka broker list")
	asyncEventProducer sarama.AsyncProducer
)

func main() {
	flag.Parse()

	router := NewRouter()

	brokerList := strings.Split(*kafkaConnect, ",")

	log.Println("Initializing kafka producer connecting to " + *kafkaConnect)
	asyncEventProducer = NewAsyncEventProducer(brokerList)

	log.Println("Starting http server listening on port " + *port)
	log.Fatal(http.ListenAndServe(*ip+":"+*port, router))
}
