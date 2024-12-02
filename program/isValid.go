package tracker
import "regexp"
// checks if the user input is valid text
func isValid(input string) bool {
	// Define a regular expression that matches a string that contains only letters
	re := regexp.MustCompile(`^[\pL0-9' ]*[\pL0-9][\pL0-9' ]*$`)
	// return the string that matches the regexp
	return re.MatchString(input)
}
