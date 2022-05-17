module github.com/sermojohn/otelkafka/test

go 1.17

require (
	github.com/confluentinc/confluent-kafka-go v1.8.2
	github.com/ory/dockertest/v3 v3.8.1
	github.com/sermojohn/otelkafka v0.1.0
	github.com/stretchr/testify v1.7.1
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
	go.uber.org/goleak v1.1.12
)

replace github.com/sermojohn/otelkafka => ../
