package selfnumber

func Sum(from int, to int) (sum int) {
	return NewSequence(from, to).FilterSelfNumber(from, to).Sum()
}
