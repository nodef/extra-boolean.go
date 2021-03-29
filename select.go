package boolean

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select(i int, a bool, b bool) bool {
  return Select2(i, a, b)
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select0(i int) bool {
  return false
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select1(i int, a bool) bool {
  switch i {
  case 0:
    return a
  default:
    return false
  }
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select2(i int, a bool, b bool) bool {
  switch i {
  case 0:
    return a
  case 1:
    return b
  default:
    return false
  }
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select3(i int, a bool, b bool, c bool) bool {
  switch i {
  case 0:
    return a
  case 1:
    return b
  case 2:
    return c
  default:
    return false
  }
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select4(i int, a bool, b bool, c bool, d bool) bool {
  switch i {
  case 0:
    return a
  case 1:
    return b
  case 2:
    return c
  case 3:
    return d
  default:
    return false
  }
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select5(i int, a bool, b bool, c bool, d bool, e bool) bool {
  switch i {
  case 0:
    return a
  case 1:
    return b
  case 2:
    return c
  case 3:
    return d
  case 4:
    return e
  default:
    return false
  }
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select6(i int, a bool, b bool, c bool, d bool, e bool, f bool) bool {
  switch i {
  case 0:
    return a
  case 1:
    return b
  case 2:
    return c
  case 3:
    return d
  case 4:
    return e
  case 5:
    return f
  default:
    return false
  }
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select7(i int, a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
  switch i {
  case 0:
    return a
  case 1:
    return b
  case 2:
    return c
  case 3:
    return d
  case 4:
    return e
  case 5:
    return f
  case 6:
    return g
  default:
    return false
  }
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select8(i int, a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
  switch i {
  case 0:
    return a
  case 1:
    return b
  case 2:
    return c
  case 3:
    return d
  case 4:
    return e
  case 5:
    return f
  case 6:
    return g
  case 7:
    return h
  default:
    return false
  }
}
