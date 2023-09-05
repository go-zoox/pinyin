package pinyin

// Abbr returns the abbreviation of the hans.
func Abbr(hans string) string {
	return FirstChars(hans)
}
