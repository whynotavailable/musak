package main

import (
	"bytes"
	"fmt"
	"github.com/whynotavailable/musak/engine"
	"github.com/whynotavailable/musak/models"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
	"gitlab.com/gomidi/midi/v2/smf"

	"github.com/whynotavailable/musak/notes"
	"gitlab.com/gomidi/midi/v2"
)

func main() {
	defer midi.CloseDriver()
	clock := smf.MetricTicks(96) // resolution: 96 ticks per quarternote 960 is also common

	chord := notes.Min(notes.A1)
	arp := notes.Arp(chord, "updown", 1)
	expanded := notes.Expand(arp, 4)

	fmt.Println(clock.Ticks8th())
	engine.Sequence(0, 0, expanded, uint64(clock.Ticks8th()))

	midiFile := models.NewMidiRequest()
	midiFile.Tempo = 100

	midiData := engine.Compile(midiFile)

	out, err := midi.FindOutPort("IAC Driver Bus 1")
	if err != nil {
		fmt.Printf("can't find qsynth")
		return
	}

	rd := bytes.NewReader(midiData)

	// read and play it
	fullTrack := smf.ReadTracksFrom(rd)
	fmt.Println(fullTrack.SMF().Format())
	fullTrack.Play(out)

	fmt.Println(notes.A4)
}
