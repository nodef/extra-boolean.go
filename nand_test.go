package boolean

import "fmt"

func ExampleNand() {
  fmt.Println(Nand(true, false))              // true
  fmt.Println(Nand(true, true))               // false
  fmt.Println(Nand4(true, true, false, true)) // true
  fmt.Println(Nand4(true, true, true, true))  // false
  // Output:
  // true
  // false
  // true
  // false
}
