package notes_test

import (
	"testing"

	"github.com/whynotavailable/musak/notes"
)

func TestMajor(t *testing.T) {
	ints := notes.Maj(notes.C3)
	if ints[1] != notes.E3 {
		t.Fatalf("Third Wrong, got %d not %d", ints[1], notes.E3)
	}

	if ints[2] != notes.G3 {
		t.Fatalf("Fifth Wrong, got %d not %d", ints[2], notes.G3)
	}
}

func TestMinor(t *testing.T) {
	ints := notes.Min(notes.C3)
	if ints[1] != notes.Ds3 {
		t.Fatalf("Third Wrong, got %d not %d", ints[1], notes.Ds3)
	}

	if ints[2] != notes.G3 {
		t.Fatalf("Fifth Wrong, got %d not %d", ints[2], notes.G3)
	}
}

func TestArp(t *testing.T) {
	n := []uint8{
		notes.C3, notes.E3, notes.G3,
	}

	n1goal := []uint8{
		notes.C3, notes.E3, notes.G3, notes.E3,
	}

	n2goal := []uint8{
		notes.C3, notes.E3, notes.G3, notes.C4, notes.E4, notes.G4, notes.E4, notes.C4, notes.G3, notes.E3,
	}

	n1 := notes.Arp(n, "updown", 1)
	n2 := notes.Arp(n, "updown", 2)

	if !NotesComp(n1goal, n1) {
		t.Fatal("n1 doesn't match")
	}

	if !NotesComp(n2goal, n2) {
		t.Fatal("n2 doesn't match")
	}
}

func TestExpand(t *testing.T) {
	n := []uint8{
		1, 2,
	}

	e1 := notes.Expand(n, 1)
	e2 := notes.Expand(n, 2)
	e3 := notes.Expand(n, 3)

	NCompF(t, e1, []uint8{
		1, 2,
	})

	NCompF(t, e2, []uint8{
		1, 2, 1, 2,
	})

	NCompF(t, e3, []uint8{
		1, 2, 1, 2, 1, 2,
	})
}

// TODO: Move this to a general location later
func NotesComp(arr1 []uint8, arr2 []uint8) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func NCompF(t *testing.T, arr1 []uint8, arr2 []uint8) {
	if !NotesComp(arr1, arr2) {
		t.Fatalf("%#v != %#v", arr1, arr2)
	}
}
