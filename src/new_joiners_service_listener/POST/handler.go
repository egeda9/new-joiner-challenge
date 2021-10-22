package main

import (
	"context"
	"fmt"
	mapper "handler/subscriber/func"
	"log"
	"os"
	"time"

	servicebus "github.com/Azure/azure-service-bus-go"
	"github.com/joho/godotenv"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func getAppInsightsClient() appinsights.TelemetryClient {
	instrumentationKey := os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY")
	if instrumentationKey == "" {
		fmt.Println("FATAL: expected environment variable APPINSIGHTS_INSTRUMENTATIONKEY not set")
		return nil
	}

	telemetryConfig := appinsights.NewTelemetryConfiguration(instrumentationKey)

	// Configure how many items can be sent in one call to the data collector:
	telemetryConfig.MaxBatchSize = 8192

	// Configure the maximum delay before sending queued telemetry:
	telemetryConfig.MaxBatchInterval = 2 * time.Second

	client := appinsights.NewTelemetryClientFromConfig(telemetryConfig)

	return client
}

func processMessage(message []byte) {
	m := new(mapper.Mapper)
	stringBody := string(message)

	client := getAppInsightsClient()
	client.TrackTrace("SBQ Message: "+stringBody, appinsights.Information)

	m.Map(stringBody)
}

func main() {
	httpInvokerPort, exists := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_HTTPWORKER_PORT: " + httpInvokerPort)
	}

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file -tags debug")
	}

	client := getAppInsightsClient()

	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	connStr := os.Getenv("SERVICEBUS_CONNECTION_STRING")
	if connStr == "" {
		client.TrackException("FATAL: expected environment variable SERVICEBUS_CONNECTION_STRING not set")
		fmt.Println("FATAL: expected environment variable SERVICEBUS_CONNECTION_STRING not set")
		return
	}

	// Create a client to communicate with a Service Bus Namespace.
	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
	if err != nil {
		client.TrackException(err)
		fmt.Println(err)
		return
	}

	// Create a client to communicate with the queue. (The queue must have already been created, see `QueueManager`)
	q, err := ns.NewQueue("joinerqueue")
	if err != nil {
		client.TrackException(err)
		fmt.Println("FATAL: ", err)
		return
	}

	for {

		err = q.ReceiveOne(
			ctx,
			servicebus.HandlerFunc(func(ctx context.Context, message *servicebus.Message) error {
				fmt.Println("SBQ Message: " + string(message.Data))
				processMessage(message.Data)
				return message.Complete(ctx)
			}))

		if err != nil {
			client.TrackTrace("queue receiver returned an error -- context.DeadlineExceeded", appinsights.Error)
			client.TrackException("queue receiver returned an error -- context.DeadlineExceeded")

			if err == context.DeadlineExceeded {
				log.Println("Waiting for messages")
			} else {

				log.Println("queue receiver returned an error", err) // error is "context canceled"
			}
		}

		time.Sleep(10 * time.Second)
	}
}
