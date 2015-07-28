package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	var config tomlConfig

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := toml.DecodeFile(dir+"/../etc/eve.toml", &config); err != nil {
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
