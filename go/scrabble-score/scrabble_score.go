package scrabble

import (
	"strings"
)

type scrabbleLetterScores struct {
	letter string
	score  int
}

var scrabbleScores = []scrabbleLetterScores{
	{"A", 1}, {"E", 1}, {"I", 1}, {"O", 1}, {"U", 1}, {"L", 1}, {"N", 1}, {"R", 1}, {"S", 1}, {"T", 1},
	{"D", 2}, {"G", 2},
	{"B", 3}, {"C", 3}, {"M", 3}, {"P", 3},
	{"F", 4}, {"H", 4}, {"V", 4}, {"W", 4}, {"Y", 4}, {"K", 5}, {"J", 8}, {"X", 8}, {"Q", 10}, {"Z", 10},
}

//Score returns the scrabble score
func Score(word string) int {
	var outScore = 0
	if word == "" {
		return outScore
	}
	for i := 0; i < len(word); i++ {
		c := string(word[i])
		for _, scores := range scrabbleScores {
			if strings.ToLower(c) == strings.ToLower(scores.letter) {
				outScore += scores.score
			}
		}
	}

	return outScore

}
