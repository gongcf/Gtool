package str

import "unicode"

// IsNumOrLetter checks the specified rune is number or letter.
func (gs *Gstr) IsNumOrLetter(r rune) bool {
	return ('0' <= r && '9' >= r) || gs.IsLetter(r)
}

// IsLetter checks the specified rune is letter.
func (*Gstr) IsLetter(r rune) bool {
	return 'a' <= r && 'z' >= r || 'A' <= r && 'Z' >= r
}

// ContainChinese checks the specified string whether contains chinese.
func (*Gstr) ContainChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}
