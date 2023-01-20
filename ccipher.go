package ccypher

// func ShiftRune(r rune, shift int) (rune, error) {

// 	shift %= 26

//		// lowercase
//		if shift < 0 || shift > 26 {
//			log.Fatal("The shift number has to be between 0 and 26")
//		}
//		if r >= 'a' && r <= 'z' {
//			r += rune(shift)
//		}
//		if r > 'z' {
//			r -= 26
//		}
//		// uppercase
//		if r >= 'A' && r <= 'Z' {
//			r += rune(shift)
//		}
//		if r > 'Z' {
//			r -= 26
//		}
//		return r, nil
//	}

func CipherText(s string) []string {
	var newS []string
	for _, r := range s {
		r = ShiftRune(r, 1)
		newS = append(newS, string(r))
	}
	return newS
}

func ShiftRune(r rune, shift int) rune {
	return r + rune(shift)
}

func DecipherText(s string) []string {
	var deciphered []string
	for _, r := range s {
		r = ShiftRune(r, -1)
		deciphered = append(deciphered, string(r))
	}
	return deciphered
}
