package boolean

// Checks if antecedent ⇒ consequent (a ⇒ b).
//
//   // Imp(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Imp(true, true)   == true
//   Imp(false, true)  == true
//   Imp(false, false) == true
//   Imp(true, false)  == false
func Imp(a bool, b bool) bool {
  return Imply(a, b)
}
