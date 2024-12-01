package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

// pastikan kafka topic sudah dibuat
var kafkaTopic = "orders-events"

func main() {
	// ----------------------------------------------------------------------------------------
	// initialization kafka producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	defer producer.Close()

	// -----------------------------------------------------------------------------------------
	// initialize route and server
	app := http.NewServeMux()
	app.HandleFunc("/orders", getOrder(producer))
	if err := http.ListenAndServe("localhost:8080", app); err != nil {
		panic(errors.Wrap(err, "main.http.ListenAndServe"))
	}
}

func getOrder(producer *kafka.Producer) http.HandlerFunc {
	type Product struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Price int64  `json:"price"`
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		dummyProduct := Product{
			Id:    "P001",
			Name:  "Laptop acer swift Go",
			Price: 8_900_000,
		}

		// serialisasi object ke json
		bytes, err := json.Marshal(dummyProduct)
		if err != nil {
			panic(errors.Wrap(err, "getOrder.json.Marshal"))
		}

		// publish/write message to kafka topic
		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &kafkaTopic,
				Partition: kafka.PartitionAny,
			},
			Value: bytes,
		}
		if err := producer.Produce(message, nil); err != nil {
			panic(errors.Wrap(err, "getOrder.producer.Produce"))
		}
		producer.Flush(1000 * 5)

		writer.Write(bytes)
	}
}
