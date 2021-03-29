package boolean

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor(a bool, b bool) bool {
  return Nor2(a, b)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor0() bool {
  return !Or0()
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor1(a bool) bool {
  return !Or1(a)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor2(a bool, b bool) bool {
  return !Or2(a, b)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor3(a bool, b bool, c bool) bool {
  return !Or3(a, b, c)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor4(a bool, b bool, c bool, d bool) bool {
  return !Or4(a, b, c, d)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor5(a bool, b bool, c bool, d bool, e bool) bool {
  return !Or5(a, b, c, d, e)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
  return !Or6(a, b, c, d, e, f)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
  return !Or7(a, b, c, d, e, f, g)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
  return !Or8(a, b, c, d, e, f, g, h)
}
