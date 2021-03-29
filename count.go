package boolean

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count(a bool, b bool) int {
  return Count2(a, b)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count0() int {
  return 0
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count1(a bool) int {
  if a {
    return 1
  }
  return 0
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count2(a bool, b bool) int {
  return Count1(a) + Count1(b)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count3(a bool, b bool, c bool) int {
  return Count2(a, b) + Count1(c)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count4(a bool, b bool, c bool, d bool) int {
  return Count2(a, b) + Count2(c, d)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count5(a bool, b bool, c bool, d bool, e bool) int {
  return Count4(a, b, c, d) + Count1(e)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count6(a bool, b bool, c bool, d bool, e bool, f bool) int {
  return Count4(a, b, c, d) + Count2(e, f)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) int {
  return Count4(a, b, c, d) + Count3(e, f, g)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) int {
  return Count4(a, b, c, d) + Count4(e, f, g, h)
}
