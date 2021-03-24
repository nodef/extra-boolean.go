// Boolean data type has two possible truth values to represent logic.
//
// 📦 Package: https://package.elm-lang.org/packages/elmw/extra-boolean/latest/
// 📘 Wiki: https://github.com/elmw/extra-boolean/wiki.
package Boolean

import "regexp"

//  PARSE

// Converts string to boolean.
//
// Parameters:
//  // s: a string
//
// Example:
//  Parse("1")        == true
//  Parse("truthy")   == true
//  Parse("not off")  == true
//  Parse("not true") == false
//  Parse("inactive") == false
//  Parse("disabled") == false
func Parse(s string) bool {
	var rfalse = regexp.MustCompile(`(negati|never|refus|wrong|fal|off)|\b(f|n|0)\b`)
	var rnegate = regexp.MustCompile(`\b(nay|nah|no|dis|un|in)`)
	var f = rfalse.MatchString(s)
	var n = len(rnegate.FindAllStringIndex(s, -1))%2 == 1
	return f == n
}

// NOT, EQ, NEQ, IMPLY, NIMPLY (FIXED)

// Checks if value is false.
//
// Parameters:
//  // a: a boolean
//
// Example:
//  Not(false) == true
//  Not(true)  == false
func Not(a bool) bool {
	return !a
}

// Checks if antecedent ⇔ consequent (a ⇔ b).
//
// Parameters:
//  // a: antecedent
//  // b: consequent
//
// Example:
//  Eq(true, true)   == true
//  Eq(false, false) == true
//  Eq(true, false)  == false
//  Eq(false, true)  == false
func Eq(a bool, b bool) bool {
	return a == b
}

// Checks if antecedent ⇎ consequent (a ⇎ b).
//
// Parameters:
//  // a: antecedent
//  // b: consequent
//
// Example:
//  Neq(true, false)  == true
//  Neq(false, true)  == true
//  Neq(true, true)   == false
//  Neq(false, false) == false
func Neq(a bool, b bool) bool {
	return a != b
}

// Checks if antecedent ⇒ consequent (a ⇒ b).
//
// Parameters:
//  // a: antecedent
//  // b: consequent
//
// Example:
//  Imply(true, true)   == true
//  Imply(false, true)  == true
//  Imply(false, false) == true
//  Imply(true, false)  == false
func Imply(a bool, b bool) bool {
	return !a || b
}

// Checks if antecedent ⇏ consequent (a ⇏ b).
//
// Parameters:
//  // a: antecedent
//  // b: consequent
//
// Example:
//  Nimply(true, false)  == true
//  Nimply(true, true)   == false
//  Nimply(false, true)  == false
//  Nimply(false, false) == false
func Nimply(a bool, b bool) bool {
	return !Imply(a, b)
}

// AND (VARIABLE)

// Checks if all values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  And(true, true)               == true
//  And(true, false)              == false
//  And4(true, true, true, true)  == true
//  And4(true, false, true, true) == false
func And(a bool, b bool) bool {
	return And2(a, b)
}

// Checks if all values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  And(true, true)               == true
//  And(true, false)              == false
//  And4(true, true, true, true)  == true
//  And4(true, false, true, true) == false
func And0() bool {
	return true
}

// Checks if all values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  And(true, true)               == true
//  And(true, false)              == false
//  And4(true, true, true, true)  == true
//  And4(true, false, true, true) == false
func And1(a bool) bool {
	return a
}

// Checks if all values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  And(true, true)               == true
//  And(true, false)              == false
//  And4(true, true, true, true)  == true
//  And4(true, false, true, true) == false
func And2(a bool, b bool) bool {
	return a && b
}

// Checks if all values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  And(true, true)               == true
//  And(true, false)              == false
//  And4(true, true, true, true)  == true
//  And4(true, false, true, true) == false
func And3(a bool, b bool, c bool) bool {
	return a && b && c
}

// Checks if all values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  And(true, true)               == true
//  And(true, false)              == false
//  And4(true, true, true, true)  == true
//  And4(true, false, true, true) == false
func And4(a bool, b bool, c bool, d bool) bool {
	return a && b && c && d
}

// Checks if all values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  And(true, true)               == true
//  And(true, false)              == false
//  And4(true, true, true, true)  == true
//  And4(true, false, true, true) == false
func And5(a bool, b bool, c bool, d bool, e bool) bool {
	return a && b && c && d && e
}

// Checks if all values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  And(true, true)               == true
//  And(true, false)              == false
//  And4(true, true, true, true)  == true
//  And4(true, false, true, true) == false
func And6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
	return a && b && c && d && e && f
}

// Checks if all values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  And(true, true)               == true
//  And(true, false)              == false
//  And4(true, true, true, true)  == true
//  And4(true, false, true, true) == false
func And7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	return a && b && c && d && e && f && g
}

// Checks if all values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  And(true, true)               == true
//  And(true, false)              == false
//  And4(true, true, true, true)  == true
//  And4(true, false, true, true) == false
func And8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return a && b && c && d && e && f && g && h
}

// OR (VARIABLE)

// Checks if any value is true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Or(true, false)                 == true
//  Or(false, false)                == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or(a bool, b bool) bool {
	return Or2(a, b)
}

// Checks if any value is true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Or(true, false)                 == true
//  Or(false, false)                == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or0() bool {
	return false
}

// Checks if any value is true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Or(true, false)                 == true
//  Or(false, false)                == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or1(a bool) bool {
	return a
}

// Checks if any value is true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Or(true, false)                 == true
//  Or(false, false)                == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or2(a bool, b bool) bool {
	return a || b
}

// Checks if any value is true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Or(true, false)                 == true
//  Or(false, false)                == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or3(a bool, b bool, c bool) bool {
	return a || b || c
}

// Checks if any value is true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Or(true, false)                 == true
//  Or(false, false)                == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or4(a bool, b bool, c bool, d bool) bool {
	return a || b || c || d
}

// Checks if any value is true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Or(true, false)                 == true
//  Or(false, false)                == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or5(a bool, b bool, c bool, d bool, e bool) bool {
	return a || b || c || d || e
}

// Checks if any value is true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Or(true, false)                 == true
//  Or(false, false)                == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
	return a || b || c || d || e || f
}

// Checks if any value is true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Or(true, false)                 == true
//  Or(false, false)                == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	return a || b || c || d || e || f || g
}

// Checks if any value is true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Or(true, false)                 == true
//  Or(false, false)                == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return a || b || c || d || e || f || g || h
}

// XOR (VARIABLE)

// Checks if odd no. of values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Xor(true, false)              == true
//  Xor(true, true)               == false
//  Xor4(true, true, true, false) == true
//  Xor4(true, true, true, true)  == false
func Xor(a bool, b bool) bool {
	return Xor2(a, b)
}

// Checks if odd no. of values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Xor(true, false)              == true
//  Xor(true, true)               == false
//  Xor4(true, true, true, false) == true
//  Xor4(true, true, true, true)  == false
func Xor0() bool {
	return false
}

// Checks if odd no. of values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Xor(true, false)              == true
//  Xor(true, true)               == false
//  Xor4(true, true, true, false) == true
//  Xor4(true, true, true, true)  == false
func Xor1(a bool) bool {
	return a
}

// Checks if odd no. of values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Xor(true, false)              == true
//  Xor(true, true)               == false
//  Xor4(true, true, true, false) == true
//  Xor4(true, true, true, true)  == false
func Xor2(a bool, b bool) bool {
	return a != b
}

// Checks if odd no. of values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Xor(true, false)              == true
//  Xor(true, true)               == false
//  Xor4(true, true, true, false) == true
//  Xor4(true, true, true, true)  == false
func Xor3(a bool, b bool, c bool) bool {
	return Xor2(Xor2(a, b), Xor1(c))
}

// Checks if odd no. of values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Xor(true, false)              == true
//  Xor(true, true)               == false
//  Xor4(true, true, true, false) == true
//  Xor4(true, true, true, true)  == false
func Xor4(a bool, b bool, c bool, d bool) bool {
	return Xor2(Xor2(a, b), Xor2(c, d))
}

// Checks if odd no. of values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Xor(true, false)              == true
//  Xor(true, true)               == false
//  Xor4(true, true, true, false) == true
//  Xor4(true, true, true, true)  == false
func Xor5(a bool, b bool, c bool, d bool, e bool) bool {
	return Xor2(Xor4(a, b, c, d), Xor1(e))
}

// Checks if odd no. of values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Xor(true, false)              == true
//  Xor(true, true)               == false
//  Xor4(true, true, true, false) == true
//  Xor4(true, true, true, true)  == false
func Xor6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
	return Xor2(Xor4(a, b, c, d), Xor2(e, f))
}

// Checks if odd no. of values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Xor(true, false)              == true
//  Xor(true, true)               == false
//  Xor4(true, true, true, false) == true
//  Xor4(true, true, true, true)  == false
func Xor7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	return Xor2(Xor4(a, b, c, d), Xor3(e, f, g))
}

// Checks if odd no. of values are true.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Xor(true, false)              == true
//  Xor(true, true)               == false
//  Xor4(true, true, true, false) == true
//  Xor4(true, true, true, true)  == false
func Xor8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return Xor2(Xor4(a, b, c, d), Xor4(e, f, g, h))
}

// COUNT (VARIABLE)

// Counts no. of true values.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Count(true, true)                 == 2
//  Count(true, false)                == 1
//  Count4(true, true, true, false)   == 3
//  Count4(false, true, false, false) == 1
func Count(a bool, b bool) int {
	return Count2(a, b)
}

// Counts no. of true values.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Count(true, true)                 == 2
//  Count(true, false)                == 1
//  Count4(true, true, true, false)   == 3
//  Count4(false, true, false, false) == 1
func Count0() int {
	return 0
}

// Counts no. of true values.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Count(true, true)                 == 2
//  Count(true, false)                == 1
//  Count4(true, true, true, false)   == 3
//  Count4(false, true, false, false) == 1
func Count1(a bool) int {
	if a {
		return 1
	}
	return 0
}

// Counts no. of true values.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Count(true, true)                 == 2
//  Count(true, false)                == 1
//  Count4(true, true, true, false)   == 3
//  Count4(false, true, false, false) == 1
func Count2(a bool, b bool) int {
	return Count1(a) + Count1(b)
}

// Counts no. of true values.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Count(true, true)                 == 2
//  Count(true, false)                == 1
//  Count4(true, true, true, false)   == 3
//  Count4(false, true, false, false) == 1
func Count3(a bool, b bool, c bool) int {
	return Count2(a, b) + Count1(c)
}

// Counts no. of true values.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Count(true, true)                 == 2
//  Count(true, false)                == 1
//  Count4(true, true, true, false)   == 3
//  Count4(false, true, false, false) == 1
func Count4(a bool, b bool, c bool, d bool) int {
	return Count2(a, b) + Count2(c, d)
}

// Counts no. of true values.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Count(true, true)                 == 2
//  Count(true, false)                == 1
//  Count4(true, true, true, false)   == 3
//  Count4(false, true, false, false) == 1
func Count5(a bool, b bool, c bool, d bool, e bool) int {
	return Count4(a, b, c, d) + Count1(e)
}

// Counts no. of true values.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Count(true, true)                 == 2
//  Count(true, false)                == 1
//  Count4(true, true, true, false)   == 3
//  Count4(false, true, false, false) == 1
func Count6(a bool, b bool, c bool, d bool, e bool, f bool) int {
	return Count4(a, b, c, d) + Count2(e, f)
}

// Counts no. of true values.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Count(true, true)                 == 2
//  Count(true, false)                == 1
//  Count4(true, true, true, false)   == 3
//  Count4(false, true, false, false) == 1
func Count7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) int {
	return Count4(a, b, c, d) + Count3(e, f, g)
}

// Counts no. of true values.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Count(true, true)                 == 2
//  Count(true, false)                == 1
//  Count4(true, true, true, false)   == 3
//  Count4(false, true, false, false) == 1
func Count8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) int {
	return Count4(a, b, c, d) + Count4(e, f, g, h)
}

// NAND (VARIABLE)

// Checks if any value is false.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Nand(true, false)              == true
//  Nand(true, true)               == false
//  Nand4(true, true, false, true) == true
//  Nand4(true, true, true, true)  == false
func Nand(a bool, b bool) bool {
	return Nand2(a, b)
}

// Checks if any value is false.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Nand(true, false)              == true
//  Nand(true, true)               == false
//  Nand4(true, true, false, true) == true
//  Nand4(true, true, true, true)  == false
func Nand0() bool {
	return !And0()
}

// Checks if any value is false.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Nand(true, false)              == true
//  Nand(true, true)               == false
//  Nand4(true, true, false, true) == true
//  Nand4(true, true, true, true)  == false
func Nand1(a bool) bool {
	return !And1(a)
}

// Checks if any value is false.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Nand(true, false)              == true
//  Nand(true, true)               == false
//  Nand4(true, true, false, true) == true
//  Nand4(true, true, true, true)  == false
func Nand2(a bool, b bool) bool {
	return !And2(a, b)
}

// Checks if any value is false.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Nand(true, false)              == true
//  Nand(true, true)               == false
//  Nand4(true, true, false, true) == true
//  Nand4(true, true, true, true)  == false
func Nand3(a bool, b bool, c bool) bool {
	return !And3(a, b, c)
}

// Checks if any value is false.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Nand(true, false)              == true
//  Nand(true, true)               == false
//  Nand4(true, true, false, true) == true
//  Nand4(true, true, true, true)  == false
func Nand4(a bool, b bool, c bool, d bool) bool {
	return !And4(a, b, c, d)
}

// Checks if any value is false.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Nand(true, false)              == true
//  Nand(true, true)               == false
//  Nand4(true, true, false, true) == true
//  Nand4(true, true, true, true)  == false
func Nand5(a bool, b bool, c bool, d bool, e bool) bool {
	return !And5(a, b, c, d, e)
}

// Checks if any value is false.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Nand(true, false)              == true
//  Nand(true, true)               == false
//  Nand4(true, true, false, true) == true
//  Nand4(true, true, true, true)  == false
func Nand6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
	return !And6(a, b, c, d, e, f)
}

// Checks if any value is false.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Nand(true, false)              == true
//  Nand(true, true)               == false
//  Nand4(true, true, false, true) == true
//  Nand4(true, true, true, true)  == false
func Nand7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	return !And7(a, b, c, d, e, f, g)
}

// Checks if any value is false.
//
// Parameters:
//  // a: 1st boolean
//  // b: 2nd boolean
//
// Example:
//  Nand(true, false)              == true
//  Nand(true, true)               == false
//  Nand4(true, true, false, true) == true
//  Nand4(true, true, true, true)  == false
func Nand8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return !And8(a, b, c, d, e, f, g, h)
}

/*
-- NOR (VARIABLE)
{-|
Checks if all values are false.
[📘](https://github.com/elmw/extra-boolean/wiki/nor)
	-- nor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nor False False              == True
	nor True False               == False
	nor4 False False False False == True
	nor4 False False True False  == False
-}
nor : Bool -> Bool -> Bool
nor = nor2


{-|
Checks if all values are false.
[📘](https://github.com/elmw/extra-boolean/wiki/nor)
	-- nor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nor False False              == True
	nor True False               == False
	nor4 False False False False == True
	nor4 False False True False  == False
-}
nor0 : Bool
nor0 =
not or0


{-|
Checks if all values are false.
[📘](https://github.com/elmw/extra-boolean/wiki/nor)
	-- nor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nor False False              == True
	nor True False               == False
	nor4 False False False False == True
	nor4 False False True False  == False
-}
nor1 : Bool -> Bool
nor1 a =
not <| or1 a


{-|
Checks if all values are false.
[📘](https://github.com/elmw/extra-boolean/wiki/nor)
	-- nor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nor False False              == True
	nor True False               == False
	nor4 False False False False == True
	nor4 False False True False  == False
-}
nor2 : Bool -> Bool -> Bool
nor2 a b =
not <| or2 a b


{-|
Checks if all values are false.
[📘](https://github.com/elmw/extra-boolean/wiki/nor)
	-- nor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nor False False              == True
	nor True False               == False
	nor4 False False False False == True
	nor4 False False True False  == False
-}
nor3 : Bool -> Bool -> Bool -> Bool
nor3 a b c =
not <| or3 a b c


{-|
Checks if all values are false.
[📘](https://github.com/elmw/extra-boolean/wiki/nor)
	-- nor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nor False False              == True
	nor True False               == False
	nor4 False False False False == True
	nor4 False False True False  == False
-}
nor4 : Bool -> Bool -> Bool -> Bool -> Bool
nor4 a b c d =
not <| or4 a b c d


{-|
Checks if all values are false.
[📘](https://github.com/elmw/extra-boolean/wiki/nor)
	-- nor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nor False False              == True
	nor True False               == False
	nor4 False False False False == True
	nor4 False False True False  == False
-}
nor5 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool
nor5 a b c d e =
not <| or5 a b c d e


{-|
Checks if all values are false.
[📘](https://github.com/elmw/extra-boolean/wiki/nor)
	-- nor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nor False False              == True
	nor True False               == False
	nor4 False False False False == True
	nor4 False False True False  == False
-}
nor6 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
nor6 a b c d e f =
not <| or6 a b c d e f


{-|
Checks if all values are false.
[📘](https://github.com/elmw/extra-boolean/wiki/nor)
	-- nor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nor False False              == True
	nor True False               == False
	nor4 False False False False == True
	nor4 False False True False  == False
-}
nor7 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
nor7 a b c d e f g =
not <| or7 a b c d e f g


{-|
Checks if all values are false.
[📘](https://github.com/elmw/extra-boolean/wiki/nor)
	-- nor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nor False False              == True
	nor True False               == False
	nor4 False False False False == True
	nor4 False False True False  == False
-}
nor8 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
nor8 a b c d e f g h =
not <| or8 a b c d e f g h




-- XNOR (VARIABLE)
{-|
Checks if even no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xnor)
	-- xnor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xnor True True              == True
	xnor False True             == False
	xnor4 True True False False == True
	xnor4 True True True False  == False
-}
xnor : Bool -> Bool -> Bool
xnor = xnor2


{-|
Checks if even no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xnor)
	-- xnor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xnor True True              == True
	xnor False True             == False
	xnor4 True True False False == True
	xnor4 True True True False  == False
-}
xnor0 : Bool
xnor0 =
not xor0


{-|
Checks if even no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xnor)
	-- xnor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xnor True True              == True
	xnor False True             == False
	xnor4 True True False False == True
	xnor4 True True True False  == False
-}
xnor1 : Bool -> Bool
xnor1 a =
not <| xor1 a


{-|
Checks if even no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xnor)
	-- xnor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xnor True True              == True
	xnor False True             == False
	xnor4 True True False False == True
	xnor4 True True True False  == False
-}
xnor2 : Bool -> Bool -> Bool
xnor2 a b =
not <| xor2 a b


{-|
Checks if even no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xnor)
	-- xnor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xnor True True              == True
	xnor False True             == False
	xnor4 True True False False == True
	xnor4 True True True False  == False
-}
xnor3 : Bool -> Bool -> Bool -> Bool
xnor3 a b c =
not <| xor3 a b c


{-|
Checks if even no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xnor)
	-- xnor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xnor True True              == True
	xnor False True             == False
	xnor4 True True False False == True
	xnor4 True True True False  == False
-}
xnor4 : Bool -> Bool -> Bool -> Bool -> Bool
xnor4 a b c d =
not <| xor4 a b c d


{-|
Checks if even no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xnor)
	-- xnor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xnor True True              == True
	xnor False True             == False
	xnor4 True True False False == True
	xnor4 True True True False  == False
-}
xnor5 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool
xnor5 a b c d e =
not <| xor5 a b c d e


{-|
Checks if even no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xnor)
	-- xnor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xnor True True              == True
	xnor False True             == False
	xnor4 True True False False == True
	xnor4 True True True False  == False
-}
xnor6 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
xnor6 a b c d e f =
not <| xor6 a b c d e f


{-|
Checks if even no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xnor)
	-- xnor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xnor True True              == True
	xnor False True             == False
	xnor4 True True False False == True
	xnor4 True True True False  == False
-}
xnor7 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
xnor7 a b c d e f g =
not <| xor7 a b c d e f g


{-|
Checks if even no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xnor)
	-- xnor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xnor True True              == True
	xnor False True             == False
	xnor4 True True False False == True
	xnor4 True True True False  == False
-}
xnor8 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
xnor8 a b c d e f g h =
not <| xor8 a b c d e f g h




-- SELECT (VARIABLE)
{-|
Checks if ith value is true.
[📘](https://github.com/elmw/extra-boolean/wiki/select)
	-- select[n] i a b ...
	-- i: index
	-- a: 1st boolean
	-- b: 2nd boolean
	select 0 True False             == True
	select 1 True False             == False
	select4 1 True True False False == True
	select4 2 True True False False == False
-}
select : Int -> Bool -> Bool -> Bool
select = select2


{-|
Checks if ith value is true.
[📘](https://github.com/elmw/extra-boolean/wiki/select)
	-- select[n] i a b ...
	-- i: index
	-- a: 1st boolean
	-- b: 2nd boolean
	select 0 True False             == True
	select 1 True False             == False
	select4 1 True True False False == True
	select4 2 True True False False == False
-}
select0 : Int -> Bool
select0 _ =
False


{-|
Checks if ith value is true.
[📘](https://github.com/elmw/extra-boolean/wiki/select)
	-- select[n] i a b ...
	-- i: index
	-- a: 1st boolean
	-- b: 2nd boolean
	select 0 True False             == True
	select 1 True False             == False
	select4 1 True True False False == True
	select4 2 True True False False == False
-}
select1 : Int -> Bool -> Bool
select1 i a =
case i of
	0 -> a
	_ -> False


{-|
Checks if ith value is true.
[📘](https://github.com/elmw/extra-boolean/wiki/select)
	-- select[n] i a b ...
	-- i: index
	-- a: 1st boolean
	-- b: 2nd boolean
	select 0 True False             == True
	select 1 True False             == False
	select4 1 True True False False == True
	select4 2 True True False False == False
-}
select2 : Int -> Bool -> Bool -> Bool
select2 i a b =
case i of
	0 -> a
	1 -> b
	_ -> False


{-|
Checks if ith value is true.
[📘](https://github.com/elmw/extra-boolean/wiki/select)
	-- select[n] i a b ...
	-- i: index
	-- a: 1st boolean
	-- b: 2nd boolean
	select 0 True False             == True
	select 1 True False             == False
	select4 1 True True False False == True
	select4 2 True True False False == False
-}
select3 : Int -> Bool -> Bool -> Bool -> Bool
select3 i a b c =
case i of
	0 -> a
	1 -> b
	2 -> c
	_ -> False


{-|
Checks if ith value is true.
[📘](https://github.com/elmw/extra-boolean/wiki/select)
	-- select[n] i a b ...
	-- i: index
	-- a: 1st boolean
	-- b: 2nd boolean
	select 0 True False             == True
	select 1 True False             == False
	select4 1 True True False False == True
	select4 2 True True False False == False
-}
select4 : Int -> Bool -> Bool -> Bool -> Bool -> Bool
select4 i a b c d =
case i of
	0 -> a
	1 -> b
	2 -> c
	3 -> d
	_ -> False


{-|
Checks if ith value is true.
[📘](https://github.com/elmw/extra-boolean/wiki/select)
	-- select[n] i a b ...
	-- i: index
	-- a: 1st boolean
	-- b: 2nd boolean
	select 0 True False             == True
	select 1 True False             == False
	select4 1 True True False False == True
	select4 2 True True False False == False
-}
select5 : Int -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
select5 i a b c d e =
case i of
	0 -> a
	1 -> b
	2 -> c
	3 -> d
	4 -> e
	_ -> False


{-|
Checks if ith value is true.
[📘](https://github.com/elmw/extra-boolean/wiki/select)
	-- select[n] i a b ...
	-- i: index
	-- a: 1st boolean
	-- b: 2nd boolean
	select 0 True False             == True
	select 1 True False             == False
	select4 1 True True False False == True
	select4 2 True True False False == False
-}
select6 : Int -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
select6 i a b c d e f =
case i of
	0 -> a
	1 -> b
	2 -> c
	3 -> d
	4 -> e
	5 -> f
	_ -> False


{-|
Checks if ith value is true.
[📘](https://github.com/elmw/extra-boolean/wiki/select)
	-- select[n] i a b ...
	-- i: index
	-- a: 1st boolean
	-- b: 2nd boolean
	select 0 True False             == True
	select 1 True False             == False
	select4 1 True True False False == True
	select4 2 True True False False == False
-}
select7 : Int -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
select7 i a b c d e f g =
case i of
	0 -> a
	1 -> b
	2 -> c
	3 -> d
	4 -> e
	5 -> f
	6 -> g
	_ -> False


{-|
Checks if ith value is true.
[📘](https://github.com/elmw/extra-boolean/wiki/select)
	-- select[n] i a b ...
	-- i: index
	-- a: 1st boolean
	-- b: 2nd boolean
	select 0 True False             == True
	select 1 True False             == False
	select4 1 True True False False == True
	select4 2 True True False False == False
-}
select8 : Int -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
select8 i a b c d e f g h =
case i of
	0 -> a
	1 -> b
	2 -> c
	3 -> d
	4 -> e
	5 -> f
	6 -> g
	7 -> h
	_ -> False



-- EQV, IMP (SHORTCUTS)
{-|
Checks if antecedent ⇔ consequent (a ⇔ b).
[📘](https://github.com/elmw/extra-boolean/wiki/eqv)
	-- eqv a b
	-- a: antecedent
	-- b: consequent
	eqv True True   == True
	eqv False False == True
	eqv True False  == False
	eqv False True  == False
-}
eqv : Bool -> Bool -> Bool
eqv = eq


{-|
Checks if antecedent ⇒ consequent (a ⇒ b).
[📘](https://github.com/elmw/extra-boolean/wiki/imp)
	-- imp a b
	-- a: antecedent
	-- b: consequent
	imp True True   == True
	imp False True  == True
	imp False False == True
	imp True False  == False
-}
imp : Bool -> Bool -> Bool
imp = imply

*/
