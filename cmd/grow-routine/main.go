package main

import (
	"context"
	"fmt"
	"grow-routine/internal/app/hub"
	sensor_service "grow-routine/internal/app/sensor"
	"grow-routine/internal/domain/sensor"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dataChan := make(chan *sensor.SensorPacket, 10)

	worker := sensor_service.NewMockSensorConsumer(
		100*time.Millisecond,
		500*time.Millisecond,
	)
	hub := hub.NewHub(dataChan, 10)

	fmt.Println("🚀 Starting sensor...")
	go worker.RunSensor(ctx, "sensor-1", dataChan)
	fmt.Println("👂 Waiting for data...")
	go hub.Run(ctx)
	for {

	}
}
