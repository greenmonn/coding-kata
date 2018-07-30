package roman_numerals

import "fmt"

func ExampleRomanNumerals() {
	fmt.Println(Roman(1))
	fmt.Println(Roman(2))
	fmt.Println(Roman(3))
	fmt.Println(Roman(4))
	fmt.Println(Roman(5))
	fmt.Println(Roman(7))

	// Output:
	// I
	// II
	// III
	// IV
	// V
	// VII
}

func ExampleRomanNumeralsTwoDigits() {
	fmt.Println(Roman(10))
	fmt.Println(Roman(39))
	fmt.Println(Roman(67))

	// Output:
	// X
	// XXXIX
	// LXVII
}

func ExampleRomanNumeralsThreeDigits() {
	fmt.Println(Roman(246))
	fmt.Println(Roman(207))

	// Output:
	// CCXLVI
	// CCVII
}

func ExampleRomanNumeralsFourDigits() {
	fmt.Println(Roman(1066))
	fmt.Println(Roman(1776))
	fmt.Println(Roman(1954))

	// Output:
	// MLXVI
	// MDCCLXXVI
	// MCMLIV
}

func ExampleUpperBound() {
	fmt.Println(boundary(1))
	fmt.Println(boundary(30))
	fmt.Println(boundary(103))
	fmt.Println(boundary(899))

	// Output:
	// 10
	// 100
	// 1000
	// 1000
}