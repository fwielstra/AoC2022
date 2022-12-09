package day6

type empty = struct{}

var flag = empty{}

func FindUniqueChunk(input string, size int) int {
	for i := 0; i < len(input)-size-1; i++ {
		chunk := input[i : i+size]
		// A map of empty structs is the most memory-efficient set in go
		m := map[rune]empty{}
		for _, r := range chunk {
			m[r] = flag
		}

		// chunk is not unique
		if len(m) != size {
			continue
		}
		return i + size
	}
	return -1
}

func FindStartMarker(input string) int {
	return FindUniqueChunk(input, 4)
}

func FindStartMessage(input string) int {
	return FindUniqueChunk(input, 14)
}
