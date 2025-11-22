package hub

import (
	"context"
	"grow-routine/internal/domain/sensor"
)

type Hub struct {
	batchSize int
	in        <-chan *sensor.SensorPacket // Только для чтения!
}

func NewHub(in <-chan *sensor.SensorPacket, batchSize int) *Hub {
	return &Hub{in: in, batchSize: batchSize}
}

func (h *Hub) Send(pool []*sensor.SensorPacket) error {
	return nil
}

func (h *Hub) Run(ctx context.Context) {
	batch := make([]*sensor.SensorPacket, 0, h.batchSize)

	for {
		select {
		case packet := <-h.in:
			batch = append(batch, packet)

			if len(batch) >= h.batchSize {
				h.processBatch(batch)
				batch = batch[:0]
			}

		case <-ctx.Done():
			if len(batch) > 0 {
				h.processBatch(batch)
			}
			return
		}
	}
}

func (h *Hub) processBatch(batch []*sensor.SensorPacket) {
	type aggData struct {
		sum   float64
		count int
		base  *sensor.SensorPacket
	}

	sums := make(map[string]*aggData)

	for _, p := range batch {
		key := p.SensorID + string(p.Type)
		if _, ok := sums[key]; !ok {
			sums[key] = &aggData{base: p}
		}
		sums[key].sum += p.Value
		sums[key].count++
	}

	result := make([]*sensor.SensorPacket, 0, len(sums))
	for _, data := range sums {
		outPacket := *data.base
		outPacket.Value = data.sum / float64(data.count)
		result = append(result, &outPacket)
	}

	h.Send(result)
}
