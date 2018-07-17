package selfnumber

type Sequence []int

func NewSequence(from int, to int) (sequence Sequence) {
	sequence = make([]int, to-from)

	for number := from; number < to; number++ {
		index := number - from
		sequence[index] = number
	}

	return
}

func (sequence Sequence) FilterSelfNumber(from int, to int) Sequence {
	for i := 1; i < to; i++ {
		number := generate(i)
		if  from <= number && number < to {
			index := number - from
			sequence[index] = 0
		}
	}

	return sequence
}

func (sequence Sequence) Sum() (sum int) {
	for _, value := range sequence {
		sum += value
	}
	return
}

func generate(number int) int {
	return number + sumOfElements(digits(number))
}

func sumOfElements(list []int) (sum int) {
	for _, value := range list {
		sum += value
	}
	return
}

func digits(number int) (result []int) {
	result = make([]int, 0)

	for number >= 10 {
		result = append(result, number%10)
		number /= 10
	}

	result = append(result, number)

	return
}

