package boolean

import "fmt"

func ExampleAnd() {
	fmt.Println(And(true, true))               // true
	fmt.Println(And(true, false))              // false
	fmt.Println(And4(true, true, true, true))  // true
	fmt.Println(And4(true, false, true, true)) // false
	// Output:
	// true
	// false
	// true
	// false
}
