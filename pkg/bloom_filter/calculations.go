package bloomfilter

import "math"

// GetOptimalHasherCountWithExpectedItemsAndOptimalArraySize returns the optimal number of hash functions
// to use for a Bloom filter with n elements and m bits.
//
// The optimal number of hash functions is given by the formula m/n * ln(2).
//
// The returned number of hash functions would be an integer.
func GetOptimalHasherCountWithExpectedItemsAndOptimalArraySize(n int, m int) int {
	return int(math.Ceil(float64(m) / float64(n) * math.Ln2))
}

// GetOptimalArraySizeWithExpectedItemsAndFalsePositiveProbability returns the optimal size of the bit array
// to use for a Bloom filter with n elements and a false positive probability p.
//
// The optimal size of the bit array is given by the formula -n * (ln(p) / ln(2)^2).
//
// The returned size would be an integer.
func GetOptimalArraySizeWithExpectedItemsAndFalsePositiveProbability(n int, p float64) int {
	return -1 * int(math.Ceil(float64(n)*math.Log(p)/math.Pow(math.Ln2, 2)))
}
