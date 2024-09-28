package notes

// TODO: May want to change this to not auto include the root later
func Intervals(root uint8, intervals ...uint8) []uint8 {
	ret := make([]uint8, len(intervals)+1)

	ret[0] = root

	for i, val := range intervals {
		ret[i+1] = root + val
	}

	return ret
}

func Maj(root uint8) []uint8 {
	return Intervals(root, 4, 7)
}

func Min(root uint8) []uint8 {
	return Intervals(root, 3, 7)
}

func Arp(notes []uint8, dir string, octaves uint8) []uint8 {
	if octaves == 0 {
		octaves = 1
	}

	tempNotes := make([]uint8, len(notes)*int(octaves))

	for i := 0; i < int(octaves); i++ {
		for j := 0; j < len(notes); j++ {
			tempNotes[(i*len(notes))+j] = notes[j] + uint8(i*12)
		}
	}

	var out []uint8

	if dir == "updown" {
		l := (len(tempNotes) * 2) - 2
		out = make([]uint8, l)

		i := 0

		for i = 0; i < len(tempNotes); i++ {
			out[i] = tempNotes[i]
		}

		for j := len(tempNotes) - 2; j > 0; j-- {
			out[i] = tempNotes[j]
			i++
		}
	}

	return out
}

func Expand(arr []uint8, times int) []uint8 {
	if times < 2 {
		return arr
	}

	out := make([]uint8, len(arr)*times)
	l := len(arr)

	for i := 0; i < l*times; i++ {
		out[i] = arr[i%l]
	}

	return out
}
