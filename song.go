package main

import (
	"fmt"

	"github.com/whynotavailable/musak/engine"
	"github.com/whynotavailable/musak/models"
	"github.com/whynotavailable/musak/notes"
	"gitlab.com/gomidi/midi/v2"
)

func main() {
	defer midi.CloseDriver()

	fmt.Println("song")
	song := models.NewMidiRequest()

	engine.Play(song)

	fmt.Println(notes.A4)
}
