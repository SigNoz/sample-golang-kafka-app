package main

// SIGUSR1 toggle the pause/resume consumption
import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	config "github.com/SigNoz/sample-golang-kafka-app/config"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
)

var (
	kafkaAddress = os.Getenv("KAFKA_ADDRESS")
)

func main() {

	cleanup := config.InitTracer()
	defer cleanup(context.Background())

	log.Println("Starting a new Sarama producer")

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10 // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	var err error
	producer, err := sarama.NewSyncProducer(strings.Split(kafkaAddress, ","), config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	producer = otelsarama.WrapSyncProducer(config, producer)

	for i := 0; i < 20; i++ {
		message := fmt.Sprintf("Hello %d", i)
		partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{Topic: "quickstart", Key: sarama.StringEncoder("message"),
			Value: sarama.StringEncoder(message)})
		if err != nil {
			log.Panicf("Error from consumer: %v", err)
		} else {
			// The tuple (topic, partition, offset) can be used as a unique identifier
			// for a message in a Kafka cluster.
			log.Printf("Your data is stored with unique identifier quickstart/%d/%d\n", partition, offset)
		}
		time.Sleep(2 * time.Second)
	}

	defer producer.Close()

}
