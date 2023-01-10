package ccypher

func ShiftRune(r rune, shift int) rune {
	if r >= 'a' && r < 'z' {
		return r + rune(shift)
	}
	if r == 'z' {
		return 'a'
	}
	return r
}
