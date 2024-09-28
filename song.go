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

	var cursor uint64
	cursor = mkArp(1, cursor, notes.Min(notes.A1))
	cursor = mkArp(1, cursor, notes.Min(notes.A1))
	cursor = mkArp(1, cursor, notes.Maj(notes.G1))
	cursor = mkArp(1, cursor, notes.Maj(notes.F1))

	midiFile := models.NewMidiRequest()
	midiFile.Tempo = 120

	midiData := engine.Compile(midiFile)

	loopPlay(midiData)
}

func mkArp(channel uint8, cursor uint64, chord []uint8) uint64 {
	arp := notes.Arp(chord, "updown", 1)
	expanded := notes.Expand(arp, 2)

	return engine.Sequence(channel, cursor, expanded, 48)
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
