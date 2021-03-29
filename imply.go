package boolean

// Checks if antecedent ⇒ consequent (a ⇒ b).
//
//   // Imply(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Imply(true, true)   == true
//   Imply(false, true)  == true
//   Imply(false, false) == true
//   Imply(true, false)  == false
func Imply(a bool, b bool) bool {
  return !a || b
}
