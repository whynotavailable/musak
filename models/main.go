package models

import "gitlab.com/gomidi/midi/v2/smf"

type MidiRequest struct {
	Tempo      float64
	TickRate   smf.MetricTicks
	MeterNum   uint8
	MeterDenum uint8
}

// Sensible defaults
func NewMidiRequest() MidiRequest {
	return MidiRequest{
		Tempo:      120,
		TickRate:   smf.MetricTicks(96),
		MeterNum:   4,
		MeterDenum: 4,
	}
}
