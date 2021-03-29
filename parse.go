package boolean

import "strconv"

// Converts string to boolean.
//
//   // Parse(s)
//   // s: a string
//
// Example:
//   Parse("1")            == true
//   Parse("truthy")       == true
//   Parse("Not Off")      == true
//   Parse("Not Inactive") == true
//   Parse("cold")         == false
//   Parse("inactive")     == false
//   Parse("Negative Yes") == false
//   Parse("Negative Aye") == false
func Parse(s string) bool {
  if rNum.MatchString(s) {
    if n, err := strconv.Atoi(s); err == nil {
      return n > 0
    }
  }
  t := rTru.MatchString(s)
  f := rFal.MatchString(s)
  n := len(rNeg.FindAllStringIndex(s, -1)) % 2 == 1
  return Nimply(f, t) == n
}
