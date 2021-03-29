package boolean

// Checks if antecedent ⇔ consequent (a ⇔ b).
//
//   // Eq(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Eq(true, true)   == true
//   Eq(false, false) == true
//   Eq(true, false)  == false
//   Eq(false, true)  == false
func Eq(a bool, b bool) bool {
  return a == b
}
