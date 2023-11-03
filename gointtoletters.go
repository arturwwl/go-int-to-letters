package gointtoletters

// IntToLetters formats number (int) to letters
func IntToLetters(number int) (letters string) {
	return Int32ToLetters(int32(number))
}

// Int32ToLetters formats number (int32) to letters
func Int32ToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += Int32ToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}
