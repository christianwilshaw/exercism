package isogram

import (
	"strings"
)

//IsIsogram returns true if the function is an isogram
func IsIsogram(in string) bool {
	outputCount := 0
	for i := 0; i < len(in); i++ {
		c := strings.ToLower(string(in[i]))
		outputCount = 0
		for j := 0; j < len(in); j++ {
			if c == strings.ToLower(string(in[j])) {
				if c == "-" || c == " " {
					break
				}
				outputCount++
				if outputCount > 1 {
					return false
				}
			}
		}
	}
	return true

}
