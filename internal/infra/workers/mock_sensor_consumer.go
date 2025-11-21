package workers

import (
	"context"
	"grow-routine/internal/domain/sensor"
	"math/rand"
	"time"
)

// MockSensorConsumer acts as an adapter for external data sources.
// In production, this would consume messages from Kafka/MQTT.
// Here it generates mock data to simulate external telemetry load.
type MockSensorConsumer struct {
	minDelay time.Duration
	maxDelay time.Duration
}

func NewMockSensorConsumer(minDelay, maxDelay time.Duration) *MockSensorConsumer {
	if minDelay <= 0 {
		minDelay = 100 * time.Millisecond
	}
	if maxDelay <= minDelay {
		maxDelay = minDelay + 100*time.Millisecond
	}

	return &MockSensorConsumer{
		minDelay: minDelay,
		maxDelay: maxDelay,
	}
}

func (c *MockSensorConsumer) RunSensor(ctx context.Context, id string, out chan<- *sensor.SensorPacket) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		select {
		case <-ctx.Done():
			return
		default:
			sleepTime := c.minDelay + time.Duration(rng.Intn(int(c.maxDelay-c.minDelay)))
			time.Sleep(sleepTime)

			packet := &sensor.SensorPacket{
				SensorID:  id,
				Type:      sensor.TypeTemperature,
				Value:     20.0 + rng.Float64()*15.0,
				Timestamp: time.Now().Unix(),
			}

			select {
			case out <- packet:
			case <-ctx.Done():
				return
			}
		}
	}
}
