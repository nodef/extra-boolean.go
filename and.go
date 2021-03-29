package boolean

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And(a bool, b bool) bool {
  return And2(a, b)
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And0() bool {
  return true
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And1(a bool) bool {
  return a
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And2(a bool, b bool) bool {
  return a && b
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And3(a bool, b bool, c bool) bool {
  return a && b && c
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And4(a bool, b bool, c bool, d bool) bool {
  return a && b && c && d
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And5(a bool, b bool, c bool, d bool, e bool) bool {
  return a && b && c && d && e
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
  return a && b && c && d && e && f
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
  return a && b && c && d && e && f && g
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
  return a && b && c && d && e && f && g && h
}
