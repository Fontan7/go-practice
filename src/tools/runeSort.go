package tools

type SortR []rune

func (s SortR) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortR) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortR) Len() int {
	return len(s)
}