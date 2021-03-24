package boolean

import (
	"fmt"
)

// PARSE

func ExampleParse() {
	fmt.Println(Parse("1"))        // true
	fmt.Println(Parse("truthy"))   // true
	fmt.Println(Parse("not off"))  // true
	fmt.Println(Parse("not true")) // false
	fmt.Println(Parse("inactive")) // false
	fmt.Println(Parse("disabled")) // false
	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
}

// NOT, EQ, NEQ, IMPLY, NIMPLY (FIXED)

func ExampleNot() {
	fmt.Println(Not(false)) // true
	fmt.Println(Not(true))  // false
	// Output:
	// true
	// false
}

func ExampleEq() {
	fmt.Println(Eq(true, true))   // true
	fmt.Println(Eq(false, false)) // true
	fmt.Println(Eq(true, false))  // false
	fmt.Println(Eq(false, true))  // false
	// Output:
	// true
	// true
	// false
	// false
}

func ExampleNeq() {
	fmt.Println(Neq(true, false))  // true
	fmt.Println(Neq(false, true))  // true
	fmt.Println(Neq(true, true))   // false
	fmt.Println(Neq(false, false)) // false
	// Output:
	// true
	// true
	// false
	// false
}

func ExampleImply() {
	fmt.Println(Imply(true, true))   // true
	fmt.Println(Imply(false, true))  // true
	fmt.Println(Imply(false, false)) // true
	fmt.Println(Imply(true, false))  // false
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleNimply() {
	fmt.Println(Nimply(true, false))  // true
	fmt.Println(Nimply(true, true))   // false
	fmt.Println(Nimply(false, true))  // false
	fmt.Println(Nimply(false, false)) // false
	// Output:
	// true
	// false
	// false
	// false
}

// AND (VARIABLE)

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

// OR (VARIABLE)

func ExampleOr() {
	fmt.Println(Or(true, false))                 // true
	fmt.Println(Or(false, false))                // false
	fmt.Println(Or4(false, true, false, true))   // true
	fmt.Println(Or4(false, false, false, false)) // false
	// Output:
	// true
	// false
	// true
	// false
}

// XOR (VARIABLE)

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

// COUNT (VARIABLE)

func ExampleCount() {
	fmt.Println(Count(true, true))                 // 2
	fmt.Println(Count(true, false))                // 1
	fmt.Println(Count4(true, true, true, false))   // 3
	fmt.Println(Count4(false, true, false, false)) // 1
	// Output:
	// 2
	// 1
	// 3
	// 1
}

// NAND (VARIABLE)

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

// NOR (VARIABLE)

func ExampleNor() {
	fmt.Println(Nor(false, false))                // true
	fmt.Println(Nor(true, false))                 // false
	fmt.Println(Nor4(false, false, false, false)) // true
	fmt.Println(Nor4(false, false, true, false))  // false
	// Output:
	// true
	// false
	// true
	// false
}

// XNOR (VARIABLE)

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
func ExampleXnor() {
	fmt.Println(Xnor(true, true))                // true
	fmt.Println(Xnor(false, true))               // false
	fmt.Println(Xnor4(true, true, false, false)) // true
	fmt.Println(Xnor4(true, true, true, false))  // false
	// Output:
	// true
	// false
	// true
	// false
}

// SELECT (VARIABLE)

func ExampleSelect() {
	fmt.Println(Select(0, true, false))               // true
	fmt.Println(Select(1, true, false))               // false
	fmt.Println(Select4(1, true, true, false, false)) // true
	fmt.Println(Select4(2, true, true, false, false)) // false
	// Output:
	// true
	// false
	// true
	// false
}

// EQV, IMP (SHORTCUTS)

func ExampleEqv() {
	fmt.Println(Eqv(true, true))   // true
	fmt.Println(Eqv(false, false)) // true
	fmt.Println(Eqv(true, false))  // false
	fmt.Println(Eqv(false, true))  // false
	// Output:
	// true
	// true
	// false
	// false
}

// Checks if antecedent ⇒ consequent (a ⇒ b).
//
//   // Imp(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
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

// func TestParse(t *testing.T) {
// 	var a, b bool
// 	a = Parse("1")
// 	b = true
// 	if a != b {
// 		t.Errorf(`Parse("1") = %t, expected %t`, a, b)
// 	}

// 	a = Parse("truthy")
// 	b = true
// 	if a != b {
// 		t.Errorf(`Parse("truthy") = %t, expected %t`, a, b)
// 	}

// 	a = Parse("not off")
// 	b = true
// 	if a != b {
// 		t.Errorf(`Parse("not off") = %t, expected %t`, a, b)
// 	}

// 	a = Parse("not true")
// 	b = false
// 	if a != b {
// 		t.Errorf(`Parse("not true") = %t, expected %t`, a, b)
// 	}

// 	a = Parse("inactive")
// 	b = false
// 	if a != b {
// 		t.Errorf(`Parse("inactive") = %t, expected %t`, a, b)
// 	}

// 	a = Parse("disabled")
// 	b = false
// 	if a != b {
// 		t.Errorf(`Parse("disabled") = %t, expected %t`, a, b)
// 	}
// }
