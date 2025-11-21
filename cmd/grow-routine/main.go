package main

import (
	"context"
	"fmt"
	"time"

	// Подставь свои пути импорта!
	"grow-routine/internal/domain/sensor"
	"grow-routine/internal/infra/workers"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dataChan := make(chan *sensor.SensorPacket, 10)

	worker := workers.NewMockSensorConsumer(
		100*time.Millisecond,
		500*time.Millisecond,
	)

	fmt.Println("🚀 Starting sensor...")
	go worker.RunSensor(ctx, "sensor-1", dataChan)

	fmt.Println("👂 Waiting for data...")

	for {
		select {
		case packet := <-dataChan:
			fmt.Printf("Received: [%s] %.2f (Type: %s)\n",
				packet.SensorID, packet.Value, packet.Type)

		case <-ctx.Done():
			fmt.Println("🛑 Time is up! Shutting down.")
			return
		}
	}
}
