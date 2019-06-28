package kata

import "math/bits"

// CountBits has
func CountBits(n uint) int {
	return bits.OnesCount(n)
}
