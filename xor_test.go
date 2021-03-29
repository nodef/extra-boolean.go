package boolean

import "fmt"

func ExampleXor() {
  fmt.Println(Xor(true, false))              // true
  fmt.Println(Xor(true, true))               // false
  fmt.Println(Xor4(true, true, true, false)) // true
  fmt.Println(Xor4(true, true, true, true))  // false
  // Output:
  // true
  // false
  // true
  // false
}
