package kata

var romans map[rune]int = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func Decode(roman string) int {
	var sum int
	for i, letter := range roman {
		if sum == 0 || (romans[roman[i-1]] >= romans[letter]) {
			sum += romans[letter]
		} else if i-1 != 0 {
			sum += (romans[letter] - romans[roman[i-1]]) - romans[roman[i-1]]
		} else {
			sum = romans[letter] - sum
		}
	}
	return sum
}
