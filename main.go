package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Shopify/sarama"
)

const (
	// Default title
	DefaultTitle = "EvE - Event Entrance"

	// Default hostname for http server
	DefaultHostname = "localhost"

	// Default bind address for http server
	DefaultBindAddress = ":8080"

	// Default kafka connect
	DefaultKafkaConnect = "localhost:9092"
)

var (
	configFile         = flag.String("config-file", "etc/eve.toml", "path to config file")
	asyncEventProducer sarama.AsyncProducer
)

type tomlConfig struct {
	Title string
	Http  httpInfo
	Kafka kafkaInfo
}

type httpInfo struct {
	Hostname    string
	BindAddress string `toml:"bind_address"`
}

type kafkaInfo struct {
	Connect string
}

func main() {
	flag.Parse()

	var config tomlConfig

	if _, err := toml.DecodeFile(*configFile, &config); err != nil {
		log.Println("No eve.toml configuration file found, taking defaults.")
		config.Title = DefaultTitle
		config.Http.Hostname = DefaultHostname
		config.Http.BindAddress = DefaultBindAddress
		config.Kafka.Connect = DefaultKafkaConnect
	}

	log.Println("Starting " + config.Title)

	router := NewRouter()

	brokerList := strings.Split(config.Kafka.Connect, ",")

	log.Println("Initializing kafka producer connecting to " + config.Kafka.Connect)
	asyncEventProducer = NewAsyncEventProducer(brokerList)

	log.Println("Starting http server " +
		config.Http.Hostname +
		config.Http.BindAddress)
	log.Fatal(http.ListenAndServe(config.Http.Hostname+config.Http.BindAddress, router))
}
