package popcount

// pc[i] is the population count of i.
var pc [256]int

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + i&1
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var numPC int
	for i := range 8 {
		numPC += pc[byte(x>>(i*8))]
	}
	return numPC
}
