package bloomfilter

import (
	"hash"
	"hash/fnv"
)

// BloomFilter is a probabilistic data structure that is used to test whether an element is a member of a set.
type BloomFilter struct {
	bits    []bool
	hashers []hash.Hash64
}

// NewBloomFilter creates a new BloomFilter with the given size and number of hash functions.
func NewBloomFilter(size int, numHashers int) *BloomFilter {
	bits := make([]bool, size)
	hashers := make([]hash.Hash64, numHashers)
	for i := 0; i < numHashers; i++ {
		hashers[i] = fnv.New64a()
	}

	return &BloomFilter{bits, hashers}
}

// NewBloomFilterWithAllowedFalsePositiveRate creates a new BloomFilter with the
// given number of expected items and allowed false positive rate.
//
// The optimal size of the bit array and the optimal number of hash functions are calculated
// based on the given number of expected items and allowed false positive rate.
func NewBloomFilterWithAllowedFalsePositiveRate(n int, p float64) *BloomFilter {
	m := GetOptimalArraySizeWithExpectedItemsAndFalsePositiveProbability(n, p)
	k := GetOptimalHasherCountWithExpectedItemsAndOptimalArraySize(n, m)
	return NewBloomFilter(m, k)
}

// Reset resets the BloomFilter.
func (bf *BloomFilter) Reset() {
	for i := range bf.bits {
		bf.bits[i] = false
	}
}

// Add adds an element to the BloomFilter.
func (bf *BloomFilter) Add(data []byte) {
	for _, hasher := range bf.hashers {
		hasher.Write(data)
		hash := hasher.Sum64()
		hasher.Reset()
		bf.bits[hash%uint64(len(bf.bits))] = true
	}
}

// Contains checks if an element is in the BloomFilter.
func (bf *BloomFilter) Contains(data []byte) bool {
	for _, hasher := range bf.hashers {
		hasher.Write(data)
		hash := hasher.Sum64()
		hasher.Reset()
		if !bf.bits[hash%uint64(len(bf.bits))] {
			return false
		}
	}
	return true
}
