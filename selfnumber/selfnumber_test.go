package selfnumber

import "fmt"

func ExampleNewSequence() {
	fmt.Println(NewSequence(1, 5))

	// Output:
	// [1 2 3 4]
}

func ExampleDigits() {
	fmt.Println(digits(123))

	// Output:
	// [3 2 1]
}

func ExampleSumOfElements() {
	list := []int{1, 2, 3}
	fmt.Println(sumOfElements(list))

	// Output:
	// 6
}

func ExampleGenerate() {
	fmt.Println(generate(91))

	// Output:
	// 101
}

func ExampleSum() {
	sum := Sum(1, 32)

	ref := 1 + 3 + 5 + 7 + 9 + 20 + 31

	fmt.Println(sum == ref)

	// Output:
	// true
}

func ExampleSumTo5000() {
	sum := Sum(1, 5000)

	fmt.Println(sum)

	// Output:
	// 1227365
}

func ExampleSumTo5000000() {
	sum := Sum(1, 5000000)

	fmt.Println(sum)

	// Output:
	// 1222233933479
}
