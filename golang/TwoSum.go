package kata

import "fmt"

// TwoSum has
func TwoSum(numbers []int, target int) [2]int {
	var h = make(map[int]int)
	for i, v := range numbers {
		need := target - v
		if _, ok := h[v]; ok {
			fmt.Println(h[v], i)
			return [2]int{h[v], i}
		}
		h[need] = i
	}
	return [2]int{0, 0}
}
