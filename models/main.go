package models

type MidiRequest struct {
	Tempo      float64
	TickRate   uint16
	MeterNum   uint8
	MeterDenum uint8
}

// Sensible defaults
func NewMidiRequest() MidiRequest {
	return MidiRequest{
		Tempo:      120,
		TickRate:   96,
		MeterNum:   4,
		MeterDenum: 4,
	}
}
