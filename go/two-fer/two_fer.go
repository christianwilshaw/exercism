//Packagetwofer provides methods for sharing.
package twofer

// ShareWith examines the input stringnd decides which output string to return.
func ShareWith(s string) string {
	output := ""

	if s == "" {
		output = "One for you, one for me."
	} else {
		output = "One for " + s + ", " + "one for me."
	}
	return output
}
