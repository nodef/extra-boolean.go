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
//  And(true, true)              == true
//  And(true, false)             == false
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
//  And(true, true)              == true
//  And(true, false)             == false
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
//  And(true, true)              == true
//  And(true, false)             == false
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
//  And(true, true)              == true
//  And(true, false)             == false
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
//  And(true, true)              == true
//  And(true, false)             == false
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
//  And(true, true)              == true
//  And(true, false)             == false
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
//  And(true, true)              == true
//  And(true, false)             == false
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
//  And(true, true)              == true
//  And(true, false)             == false
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
//  And(true, true)              == true
//  And(true, false)             == false
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
//  And(true, true)              == true
//  And(true, false)             == false
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
//  Or(true, false)                == true
//  Or(false, false)               == false
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
//  Or(true, false)                == true
//  Or(false, false)               == false
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
//  Or(true, false)                == true
//  Or(false, false)               == false
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
//  Or(true, false)                == true
//  Or(false, false)               == false
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
//  Or(true, false)                == true
//  Or(false, false)               == false
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
//  Or(true, false)                == true
//  Or(false, false)               == false
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
//  Or(true, false)                == true
//  Or(false, false)               == false
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
//  Or(true, false)                == true
//  Or(false, false)               == false
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
//  Or(true, false)                == true
//  Or(false, false)               == false
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
//  Or(true, false)                == true
//  Or(false, false)               == false
//  Or4(false, true, false, true)   == true
//  Or4(false, false, false, false) == false
func Or8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return a || b || c || d || e || f || g || h
}

/*

-- XOR (VARIABLE)
{-|
Checks if odd no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xor)
	-- xor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xor True False            == True
	xor True True             == False
	xor4 True True True False == True
	xor4 True True True True  == False
-}
xor : Bool -> Bool -> Bool
xor = xor2


{-|
Checks if odd no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xor)
	-- xor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xor True False            == True
	xor True True             == False
	xor4 True True True False == True
	xor4 True True True True  == False
-}
xor0 : Bool
xor0 =
False


{-|
Checks if odd no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xor)
	-- xor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xor True False            == True
	xor True True             == False
	xor4 True True True False == True
	xor4 True True True True  == False
-}
xor1 : Bool -> Bool
xor1 a =
a


{-|
Checks if odd no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xor)
	-- xor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xor True False            == True
	xor True True             == False
	xor4 True True True False == True
	xor4 True True True True  == False
-}
xor2 : Bool -> Bool -> Bool
xor2 = Basics.xor


{-|
Checks if odd no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xor)
	-- xor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xor True False            == True
	xor True True             == False
	xor4 True True True False == True
	xor4 True True True True  == False
-}
xor3 : Bool -> Bool -> Bool -> Bool
xor3 a b c =
xor2 (xor2 a b) (xor1 c)


{-|
Checks if odd no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xor)
	-- xor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xor True False            == True
	xor True True             == False
	xor4 True True True False == True
	xor4 True True True True  == False
-}
xor4 : Bool -> Bool -> Bool -> Bool -> Bool
xor4 a b c d =
xor2 (xor2 a b) (xor2 c d)


{-|
Checks if odd no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xor)
	-- xor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xor True False            == True
	xor True True             == False
	xor4 True True True False == True
	xor4 True True True True  == False
-}
xor5 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool
xor5 a b c d e =
xor2 (xor4 a b c d) (xor1 e)


{-|
Checks if odd no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xor)
	-- xor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xor True False            == True
	xor True True             == False
	xor4 True True True False == True
	xor4 True True True True  == False
-}
xor6 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
xor6 a b c d e f =
xor2 (xor4 a b c d) (xor2 e f)


{-|
Checks if odd no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xor)
	-- xor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xor True False            == True
	xor True True             == False
	xor4 True True True False == True
	xor4 True True True True  == False
-}
xor7 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
xor7 a b c d e f g =
xor2 (xor4 a b c d) (xor3 e f g)


{-|
Checks if odd no. of values are true.
[📘](https://github.com/elmw/extra-boolean/wiki/xor)
	-- xor[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	xor True False            == True
	xor True True             == False
	xor4 True True True False == True
	xor4 True True True True  == False
-}
xor8 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
xor8 a b c d e f g h =
xor2 (xor4 a b c d) (xor4 e f g h)




-- COUNT (VARIABLE)
{-|
Count no. of true values.
[📘](https://github.com/elmw/extra-boolean/wiki/count)
	-- count[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	count True True               == 2
	count True False              == 1
	count4 True True True False   == 3
	count4 False True False False == 1
-}
count : Bool -> Bool -> Int
count = count2


{-|
Count no. of true values.
[📘](https://github.com/elmw/extra-boolean/wiki/count)
	-- count[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	count True True               == 2
	count True False              == 1
	count4 True True True False   == 3
	count4 False True False False == 1
-}
count0 : Int
count0 =
0


{-|
Count no. of true values.
[📘](https://github.com/elmw/extra-boolean/wiki/count)
	-- count[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	count True True               == 2
	count True False              == 1
	count4 True True True False   == 3
	count4 False True False False == 1
-}
count1 : Bool -> Int
count1 a =
if a then 1 else 0


{-|
Count no. of true values.
[📘](https://github.com/elmw/extra-boolean/wiki/count)
	-- count[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	count True True               == 2
	count True False              == 1
	count4 True True True False   == 3
	count4 False True False False == 1
-}
count2 : Bool -> Bool -> Int
count2 a b =
count1 a + count1 b


{-|
Count no. of true values.
[📘](https://github.com/elmw/extra-boolean/wiki/count)
	-- count[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	count True True               == 2
	count True False              == 1
	count4 True True True False   == 3
	count4 False True False False == 1
-}
count3 : Bool -> Bool -> Bool -> Int
count3 a b c =
count2 a b + count1 c


{-|
Count no. of true values.
[📘](https://github.com/elmw/extra-boolean/wiki/count)
	-- count[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	count True True               == 2
	count True False              == 1
	count4 True True True False   == 3
	count4 False True False False == 1
-}
count4 : Bool -> Bool -> Bool -> Bool -> Int
count4 a b c d =
count2 a b + count2 c d


{-|
Count no. of true values.
[📘](https://github.com/elmw/extra-boolean/wiki/count)
	-- count[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	count True True               == 2
	count True False              == 1
	count4 True True True False   == 3
	count4 False True False False == 1
-}
count5 : Bool -> Bool -> Bool -> Bool -> Bool -> Int
count5 a b c d e =
count4 a b c d + count1 e


{-|
Count no. of true values.
[📘](https://github.com/elmw/extra-boolean/wiki/count)
	-- count[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	count True True               == 2
	count True False              == 1
	count4 True True True False   == 3
	count4 False True False False == 1
-}
count6 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Int
count6 a b c d e f =
count4 a b c d + count2 e f


{-|
Count no. of true values.
[📘](https://github.com/elmw/extra-boolean/wiki/count)
	-- count[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	count True True               == 2
	count True False              == 1
	count4 True True True False   == 3
	count4 False True False False == 1
-}
count7 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Int
count7 a b c d e f g =
count4 a b c d + count3 e f g


{-|
Count no. of true values.
[📘](https://github.com/elmw/extra-boolean/wiki/count)
	-- count[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	count True True               == 2
	count True False              == 1
	count4 True True True False   == 3
	count4 False True False False == 1
-}
count8 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Int
count8 a b c d e f g h =
count4 a b c d + count4 e f g h




-- NAND (VARIABLE)
{-|
Checks if any value is false.
[📘](https://github.com/elmw/extra-boolean/wiki/nand)
	-- nand[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nand True False            == True
	nand True True             == False
	nand4 True True False True == True
	nand4 True True True True  == False
-}
nand : Bool -> Bool -> Bool
nand = nand2


{-|
Checks if any value is false.
[📘](https://github.com/elmw/extra-boolean/wiki/nand)
	-- nand[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nand True False            == True
	nand True True             == False
	nand4 True True False True == True
	nand4 True True True True  == False
-}
nand0 : Bool
nand0 =
not <| and0


{-|
Checks if any value is false.
[📘](https://github.com/elmw/extra-boolean/wiki/nand)
	-- nand[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nand True False            == True
	nand True True             == False
	nand4 True True False True == True
	nand4 True True True True  == False
-}
nand1 : Bool -> Bool
nand1 a =
not <| and1 a


{-|
Checks if any value is false.
[📘](https://github.com/elmw/extra-boolean/wiki/nand)
	-- nand[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nand True False            == True
	nand True True             == False
	nand4 True True False True == True
	nand4 True True True True  == False
-}
nand2 : Bool -> Bool -> Bool
nand2 a b =
not <| and2 a b


{-|
Checks if any value is false.
[📘](https://github.com/elmw/extra-boolean/wiki/nand)
	-- nand[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nand True False            == True
	nand True True             == False
	nand4 True True False True == True
	nand4 True True True True  == False
-}
nand3 : Bool -> Bool -> Bool -> Bool
nand3 a b c =
not <| and3 a b c


{-|
Checks if any value is false.
[📘](https://github.com/elmw/extra-boolean/wiki/nand)
	-- nand[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nand True False            == True
	nand True True             == False
	nand4 True True False True == True
	nand4 True True True True  == False
-}
nand4 : Bool -> Bool -> Bool -> Bool -> Bool
nand4 a b c d =
not <| and4 a b c d


{-|
Checks if any value is false.
[📘](https://github.com/elmw/extra-boolean/wiki/nand)
	-- nand[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nand True False            == True
	nand True True             == False
	nand4 True True False True == True
	nand4 True True True True  == False
-}
nand5 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool
nand5 a b c d e =
not <| and5 a b c d e


{-|
Checks if any value is false.
[📘](https://github.com/elmw/extra-boolean/wiki/nand)
	-- nand[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nand True False            == True
	nand True True             == False
	nand4 True True False True == True
	nand4 True True True True  == False
-}
nand6 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
nand6 a b c d e f =
not <| and6 a b c d e f


{-|
Checks if any value is false.
[📘](https://github.com/elmw/extra-boolean/wiki/nand)
	-- nand[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nand True False            == True
	nand True True             == False
	nand4 True True False True == True
	nand4 True True True True  == False
-}
nand7 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
nand7 a b c d e f g =
not <| and7 a b c d e f g


{-|
Checks if any value is false.
[📘](https://github.com/elmw/extra-boolean/wiki/nand)
	-- nand[n] a b ...
	-- a: 1st boolean
	-- b: 2nd boolean
	nand True False            == True
	nand True True             == False
	nand4 True True False True == True
	nand4 True True True True  == False
-}
nand8 : Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool -> Bool
nand8 a b c d e f g h =
not <| and8 a b c d e f g h




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
