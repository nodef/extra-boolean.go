package boolean

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand(a bool, b bool) bool {
  return Nand2(a, b)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand0() bool {
  return !And0()
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand1(a bool) bool {
  return !And1(a)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand2(a bool, b bool) bool {
  return !And2(a, b)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand3(a bool, b bool, c bool) bool {
  return !And3(a, b, c)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand4(a bool, b bool, c bool, d bool) bool {
  return !And4(a, b, c, d)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand5(a bool, b bool, c bool, d bool, e bool) bool {
  return !And5(a, b, c, d, e)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
  return !And6(a, b, c, d, e, f)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
  return !And7(a, b, c, d, e, f, g)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
  return !And8(a, b, c, d, e, f, g, h)
}
