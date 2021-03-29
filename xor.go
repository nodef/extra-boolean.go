package boolean

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor(a bool, b bool) bool {
  return Xor2(a, b)
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor0() bool {
  return false
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor1(a bool) bool {
  return a
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor2(a bool, b bool) bool {
  return a != b
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor3(a bool, b bool, c bool) bool {
  return Xor2(Xor2(a, b), Xor1(c))
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor4(a bool, b bool, c bool, d bool) bool {
  return Xor2(Xor2(a, b), Xor2(c, d))
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor5(a bool, b bool, c bool, d bool, e bool) bool {
  return Xor2(Xor4(a, b, c, d), Xor1(e))
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
  return Xor2(Xor4(a, b, c, d), Xor2(e, f))
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
  return Xor2(Xor4(a, b, c, d), Xor3(e, f, g))
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
  return Xor2(Xor4(a, b, c, d), Xor4(e, f, g, h))
}
