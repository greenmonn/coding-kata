package jollyjumper

import "fmt"

// 1 4 2 3 => Jolly
// 1 4 2 -1 6 => Not Jolly
// 11 7 4 2 1 6 => Jolly

func ExampleIsJolly() {
	sequence1 := []int{1, 4, 2, 3}
	sequence2 := []int{1, 4, 2, -1, 6}

	fmt.Println(IsJolly(sequence1))
	fmt.Println(IsJolly(sequence2))

	// Output:
	// true
	// false
}

func ExampleIsContinuousSequence() {
	sequence1 := IntSequence{1, 2, 3, 4}
	sequence2 := IntSequence{1, 4, 2, -1, 6}

	fmt.Println(sequence1.isContinuousSequence(1, 4))
	fmt.Println(sequence2.isContinuousSequence(1, 4))

	// Output:
	// true
	// false
}

func ExampleDiffs() {
	sequence := IntSequence{1, 4, 2, 3}

	fmt.Println(sequence.diffs())

	// Output:
	// [3 2 1]
}

func ExampleSort() {
	list := []int{2, 1, 3}

	sort(list)

	fmt.Println(list)

	// Output:
	// [1 2 3]
}
