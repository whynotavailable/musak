package main

import (
	"bytes"
	"fmt"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
	"gitlab.com/gomidi/midi/v2/smf"
)

func main() {
	defer midi.CloseDriver()

	out, err := midi.FindOutPort("IAC Driver Bus 1")
	if err != nil {
		fmt.Printf("can't find qsynth")
		return
	}
	fmt.Printf("outports: %+v\n", midi.GetOutPorts())

	clock := smf.MetricTicks(96)
	fmt.Printf("%d %d\n", clock.Ticks4th(), clock.Ticks32th())

	rd := bytes.NewReader(mkSMF())

	// read and play it
	fullTrack := smf.ReadTracksFrom(rd)
	fmt.Println(fullTrack.SMF().Format())
	fullTrack.Play(out)
}

func mkSMF() []byte {
	var (
		bf    bytes.Buffer
		clock = smf.MetricTicks(96) // resolution: 96 ticks per quarternote 960 is also common
		tr    smf.Track
	)

	// first track must have tempo and meter informations
	tr.Add(0, smf.MetaMeter(4, 4))
	tr.Add(0, smf.MetaTempo(100))
	tr.Add(0, midi.NoteOn(0, midi.Ab(3), 64))
	tr.Add(clock.Ticks8th(), midi.NoteOn(0, midi.C(4), 64))
	// duration: a quarter note (96 ticks in our case)
	tr.Add(clock.Ticks4th()*2, midi.NoteOff(0, midi.Ab(3)))
	tr.Add(0, midi.NoteOff(0, midi.C(4)))
	tr.Close(0)

	// create the SMF and add the tracks
	s := smf.New()
	s.TimeFormat = clock
	s.Add(tr)
	s.WriteTo(&bf)
	return bf.Bytes()
}
