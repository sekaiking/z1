package piscine

func Rot14(s string) string {

	runeSlice := []rune(s)

	// A-Z = 90-65
	// a-z = 97-122

	for i, c := range runeSlice {
		if c >= 65 && c <= 90 {
			runeSlice[i] = c + 14
			if (c + 14) > 90 {
				runeSlice[i] -= 26
			}
		} else if c >= 97 && c <= 122 {
			runeSlice[i] = c + 14
			if (c + 14) > 122 {
				runeSlice[i] -= 26
			}
		}
	}

	return string(runeSlice)
}
