package opgg

import (
	"strings"
)

func arrayElementIsPresent(ref []string, element string) bool {
	for _, i := range ref {
		if i == element {
			return true
		}
	}
	return false
}

func findAndDeleteCommonPattern(ref []string) []string {
	commonWords := make(map[string]uint)

	for _, l := range ref {
		words := strings.Split(l, " ")
		for _, w := range words {
			commonWords[w]++
		}
	}
	mostUsedWords := make([]string, 0)
	var mostUsedTime uint = 0
	for _, cw := range commonWords {
		if cw > mostUsedTime {
			mostUsedTime = cw
		}
	}
	if mostUsedTime <= 1 {
		return ref
	}
	for words, used := range commonWords {
		if used >= mostUsedTime {
			mostUsedWords = append(mostUsedWords, words)
		}
	}
	for i, l := range ref {
		words := strings.Split(l, " ")
		cleanedLine := make([]string, 0)
		for _, w := range words {
			if !arrayElementIsPresent(mostUsedWords, w) {
				cleanedLine = append(cleanedLine, w)
			}
		}
		ref[i] = strings.Join(cleanedLine, " ")
	}
	return ref
}
