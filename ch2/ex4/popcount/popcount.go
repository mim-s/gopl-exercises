package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	n := 0
	for range 64 {
		n += int(x & 1)
		x >>= 1
	}
	return n
}

// Early exit version:
// func PopCount(x uint64) int {
// 	var pC int
// 	for ; x != 0; x >>= 1 {
// 		pC += int(x & 1)
// 	}
// 	return pC
// }
