package utils

type SlicesOfSlices [][]string

func (s SlicesOfSlices) Len() int { return len(s) }

func (s SlicesOfSlices) Less(i, j int) bool {
	if len(s[i]) < len(s[j]) {
		return true
	}
	if len(s[i]) > len(s[j]) {
		return false
	}
	// the length is equal, sort according to item(from first to last)
	for k := range s[i] {
		if s[i][k] < s[j][k] {
			return true
		}
		if s[i][k] > s[j][k] {
			return false
		}
	}

	return true
}

func (s SlicesOfSlices) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
