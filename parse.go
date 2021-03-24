package boolean

import "regexp"

// PARSE

// Converts string to boolean.
//
//   // Parse(s)
//   // s: a string
//
// Example:
//   Parse("1")        == true
//   Parse("truthy")   == true
//   Parse("not off")  == true
//   Parse("not true") == false
//   Parse("inactive") == false
//   Parse("disabled") == false
func Parse(s string) bool {
	var rfalse = regexp.MustCompile(`(negati|never|refus|wrong|fal|off)|\b(f|n|0)\b`)
	var rnegate = regexp.MustCompile(`\b(nay|nah|no|dis|un|in)`)
	var f = rfalse.MatchString(s)
	var n = len(rnegate.FindAllStringIndex(s, -1))%2 == 1
	return f == n
}
