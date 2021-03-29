package boolean

// Checks if antecedent ⇏ consequent (a ⇏ b).
//
//   // Nimply(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Nimply(true, false)  == true
//   Nimply(true, true)   == false
//   Nimply(false, true)  == false
//   Nimply(false, false) == false
func Nimply(a bool, b bool) bool {
  return !Imply(a, b)
}
