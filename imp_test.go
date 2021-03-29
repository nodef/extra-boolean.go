package boolean

import "fmt"

func ExampleImp() {
	fmt.Println(Imp(true, true))   // true
	fmt.Println(Imp(false, true))  // true
	fmt.Println(Imp(false, false)) // true
	fmt.Println(Imp(true, false))  // false
	// Output:
	// true
	// true
	// true
	// false
}
