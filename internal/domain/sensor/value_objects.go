package sensor

type MeasurementType string

const (
	TypeTemperature MeasurementType = "TEMP"
	TypeHumidity    MeasurementType = "HUM"
)

type SensorPacket struct {
	SensorID  string          `json:"sensor_id"`
	Type      MeasurementType `json:"type"`
	Value     float64         `json:"value"`
	Timestamp int64           `json:"timestamp"`
}
