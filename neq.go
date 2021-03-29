package boolean

// Checks if antecedent ⇎ consequent (a ⇎ b).
//
//   // Neq(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Neq(true, false)  == true
//   Neq(false, true)  == true
//   Neq(true, true)   == false
//   Neq(false, false) == false
func Neq(a bool, b bool) bool {
  return a != b
}
