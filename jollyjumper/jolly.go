package jollyjumper

type IntSequence []int

func IsJolly(sequence IntSequence) bool {
	length := len(sequence)-1

	diffs := sequence.diffs()

	sort(diffs)

	return diffs.isContinuousSequence(1, length)
}

func (sequence IntSequence) diffs() (diffs IntSequence){
	length := len(sequence) - 1
	diffs = make([]int, length)

	for i := 0; i<length; i++ {
		diffs[i] = abs(sequence[i] - sequence[i+1])
	}

	return
}

func (sequence IntSequence) isContinuousSequence(from int, to int) bool {
	for i := from; i <= to; i++ {
		index := i-from
		if sequence[index] != i {
			return false
		}
	}
	return true
}

func abs(number int) int {
	if number < 0 {
		return -1 * number
	}

	return number
}

func sort(list []int) {
	for i := 0; i<len(list); i++ {
		for j := i+1; j<len(list); j++ {
			if list[i] > list[j] {
				swap(&list[i], &list[j])
			}
		}
	}
}

func swap(left *int, right *int) {
	tmp := *left
	*left = *right
	*right = tmp
}
