# Go Kafka sample app


## Start Kafka

* `cd kafka`
* `docker-compose up -d`


## Run Consumer
* `cd consumer`
* `SERVICE_NAME=goKafkaConsumer INSECURE_MODE=true OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317 KAFKA_ADDRESS=localhost:9092 go run consumer.go`

## Run Producer
* `cd producer`
* `SERVICE_NAME=goKafkaProducer INSECURE_MODE=true OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317 KAFKA_ADDRESS=localhost:9092 go run producer.go`