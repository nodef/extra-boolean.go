package boolean

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor(a bool, b bool) bool {
  return Xnor2(a, b)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor0() bool {
  return !Xor0()
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor1(a bool) bool {
  return !Xor1(a)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor2(a bool, b bool) bool {
  return !Xor2(a, b)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor3(a bool, b bool, c bool) bool {
  return !Xor3(a, b, c)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor4(a bool, b bool, c bool, d bool) bool {
  return !Xor4(a, b, c, d)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor5(a bool, b bool, c bool, d bool, e bool) bool {
  return !Xor5(a, b, c, d, e)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
  return !Xor6(a, b, c, d, e, f)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
  return !Xor7(a, b, c, d, e, f, g)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
  return !Xor8(a, b, c, d, e, f, g, h)
}
