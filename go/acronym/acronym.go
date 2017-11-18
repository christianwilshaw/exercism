package acronym

import "strings"
//Takes a string and returns an acronym.
func Abbreviate(s string) string {

	stringTofmt := strings.Split(s, "")
	var output = ""
	var length = len(stringTofmt)
	var previousChar = ""
	for i := 0; i < length; i++ {

		if i == 0 || previousChar == " " || previousChar == "-" {
			output += stringTofmt[i]
		}
		previousChar = stringTofmt[i]

	}
	return strings.ToUpper(output)
}
