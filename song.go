package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/whynotavailable/musak/engine"
	"github.com/whynotavailable/musak/models"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
	"gitlab.com/gomidi/midi/v2/smf"
	"os"

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

	loopPlay(midiData)
}

func loopPlay(file []byte) {
	sig := make(chan int, 1)
	finish := make(chan int, 1)

	go func() {
		for {
			select {
			case <-sig:
				finish <- 1
				return
			default:
			}

			playSong(file)
		}
	}()
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
	sig <- 1
	<-finish
	reset()
}

func reset() {
	out, _ := midi.FindOutPort("IAC Driver Bus 1")
	send, _ := midi.SendTo(out)
	_ = send(midi.Reset())
}

func playSong(file []byte) {
	out, err := midi.FindOutPort("IAC Driver Bus 1")
	if err != nil {
		fmt.Printf("can't find qsynth")
		return
	}

	rd := bytes.NewReader(file)

	// read and play it
	fullTrack := smf.ReadTracksFrom(rd)
	fmt.Println(fullTrack.SMF().Format())
	fullTrack.Play(out)

	fmt.Println(notes.A4)

}
