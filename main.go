package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	ip                 = flag.String("ip", "0.0.0.0", "ip address (default 0.0.0.0)")
	port               = flag.String("port", "8080", "port (default 8080)")
	kafkaConnect       = flag.String("kafka_connect", "localhost:9092", "kafka broker list (default localhost:9092)")
	asyncEventProducer sarama.AsyncProducer
)

func main() {
	flag.Parse()

	router := NewRouter()

	brokerList := strings.Split(*kafkaConnect, ",")

	log.Println("Initializing kafka producer with connection string " + *kafkaConnect)
	asyncEventProducer = NewAsyncEventProducer(brokerList)

	log.Fatal(http.ListenAndServe(*ip+":"+*port, router))
}
