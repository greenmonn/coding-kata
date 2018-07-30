package roman_numerals

// list := [I, V, X, L, C, D, M]

var (
	m = map[int]string{
		0: "",
		1: "I",
		5: "V",
		10: "X",
		50: "L",
		100: "C",
		500: "D",
		1000: "M",
	}
)

func boundary(number int) int {
	bounds := []int{10, 100, 1000}

	for _, bound := range bounds {
		if number < bound {
			return bound
		}
	}

	return 100000
}

func Roman(number int) (romanNumeral string) {
	bound := boundary(number)
	half := bound/2
	base := bound/10

	if bound > 1000 {
		return m[1000] + Roman(number-1000)
	}

	if number < half - base {
		for i := 0; i < number/base; i++ {
			romanNumeral += m[base]
		}

		remains := number % base

		if remains == 0 {
			return romanNumeral
		} else {
			return romanNumeral + Roman(number % base)
		}

	} else if number < bound/2 {

		return m[base] + m[half] + Roman(number - half + base)

	} else if number < bound - bound/10 {

		return m[half] + Roman(number - half)

	} else if number < bound {

		return m[base] + m[bound] + Roman(number - bound + base)

	}

	return
}
