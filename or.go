package boolean

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or(a bool, b bool) bool {
  return Or2(a, b)
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or0() bool {
  return false
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or1(a bool) bool {
  return a
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or2(a bool, b bool) bool {
  return a || b
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or3(a bool, b bool, c bool) bool {
  return a || b || c
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or4(a bool, b bool, c bool, d bool) bool {
  return a || b || c || d
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or5(a bool, b bool, c bool, d bool, e bool) bool {
  return a || b || c || d || e
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
  return a || b || c || d || e || f
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
  return a || b || c || d || e || f || g
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
  return a || b || c || d || e || f || g || h
}
