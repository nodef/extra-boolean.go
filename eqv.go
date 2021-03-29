package boolean

// Checks if antecedent ⇔ consequent (a ⇔ b).
//
//   // Eqv(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Eqv(true, true)   == true
//   Eqv(false, false) == true
//   Eqv(true, false)  == false
//   Eqv(false, true)  == false
func Eqv(a bool, b bool) bool {
  return Eq(a, b)
}
