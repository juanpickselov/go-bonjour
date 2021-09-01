package morestrings

func ReverseRunes(s string) string {
	runeList := []rune(s)
	for i, j := 0, len(runeList)-1; i < len(runeList)/2; i, j = i+1, j-1 {
		runeList[i], runeList[j] = runeList[j], runeList[i]
	}
	return string(runeList)
}
