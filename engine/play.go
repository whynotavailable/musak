package engine

import (
	"bytes"
	"cmp"
	"fmt"
	"github.com/whynotavailable/musak/models"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
	"slices"
)

func Compile(midiFile models.MidiRequest) []byte {
	var (
		bf bytes.Buffer
		tr smf.Track
	)

	// first track must have tempo and meter informations
	tr.Add(0, smf.MetaMeter(midiFile.MeterNum, midiFile.MeterDenum))
	tr.Add(0, smf.MetaTempo(midiFile.Tempo))

	slices.SortFunc(Events, func(a, b EventItem) int {
		return cmp.Compare(a.Pos, b.Pos)
	})

	var cursor uint64

	for _, event := range Events {
		pos := event.Pos

		var delta uint32 = uint32(pos - cursor)
		fmt.Println(pos)
		cursor = pos

		switch e := event.Event.(type) {
		case NoteOnEvent:
			fmt.Println("note on", delta, e.Note)
			tr.Add(delta, midi.NoteOn(e.Channel, e.Note, e.Velocity))
		case NoteOffEvent:
			fmt.Println("note off", delta, e.Note)
			tr.Add(delta, midi.NoteOff(e.Channel, e.Note))
		}
	}

	tr.Close(0)

	// create the SMF and add the tracks
	s := smf.New()
	s.TimeFormat = midiFile.TickRate
	s.Add(tr)
	s.WriteTo(&bf)
	return bf.Bytes()

}

// Used to build eventlist
type SongEvent interface{}

// When you add a note it will be turned into two events
type NoteEvent struct {
	Channel  uint8
	Pos      uint64
	Length   uint64
	Note     uint8
	Velocity uint8
}

type NoteOnEvent struct {
	Channel  uint8
	Pos      uint64
	Note     uint8
	Velocity uint8
}

type NoteOffEvent struct {
	Channel uint8
	Pos     uint64
	Note    uint8
}

type EventItem struct {
	Pos   uint64
	Event SongEvent
}

var Events []EventItem

func addEvent(pos uint64, event SongEvent) {
	fmt.Println("Added event", pos, event)
	Events = append(Events, EventItem{
		Pos:   pos,
		Event: event,
	})
}

// Clamp sets a max value to 127, useful for random values
func Clamp(num uint8) uint8 {
	if num > 127 {
		return 127
	} else {
		return num
	}
}

func AddNote(event NoteEvent) {
	event.Velocity = Clamp(event.Velocity)

	addEvent(event.Pos, NoteOnEvent{
		Channel:  event.Channel,
		Pos:      event.Pos,
		Note:     event.Note,
		Velocity: event.Velocity,
	})

	addEvent(event.Pos+event.Length, NoteOffEvent{
		Channel: event.Channel,
		Pos:     event.Pos + event.Length,
		Note:    event.Note,
	})
}

func Sequence(channel uint8, cursor uint64, notes []uint8, length uint64) uint64 {
	for _, note := range notes {
		AddNote(NoteEvent{
			Channel:  channel,
			Pos:      cursor,
			Length:   length,
			Note:     note,
			Velocity: 64,
		})

		cursor += length
	}

	return cursor
}
