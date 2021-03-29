package boolean

// Checks if value is false.
//
//   // Not(a, b)
//   // a: a boolean
//
// Example:
//   Not(false) == true
//   Not(true)  == false
func Not(a bool) bool {
  return !a
}
