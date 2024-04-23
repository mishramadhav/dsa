package bloomfilter_test

import (
	"testing"

	bloomfilter "github.com/mishramadhav/dsa/pkg/bloom_filter"
	"github.com/stretchr/testify/assert"
)

func TestBloomFilter(t *testing.T) {
	t.Parallel()

	t.Run("Add and Contains", func(t *testing.T) {
		t.Parallel()

		bf := bloomfilter.NewBloomFilter(100, 3)
		bf.Add([]byte("hello"))
		bf.Add([]byte("world"))

		assert.True(t, bf.Contains([]byte("hello")), "expected 'hello' to be in the Bloom filter")
		assert.True(t, bf.Contains([]byte("world")), "expected 'world' to be in the Bloom filter")
		assert.False(t, bf.Contains([]byte("foo")), "expected 'foo' to not be in the Bloom filter")
	})

	t.Run("NewBloomFilterWithAllowedFalsePositiveRate", func(t *testing.T) {
		t.Parallel()

		bf := bloomfilter.NewBloomFilterWithAllowedFalsePositiveRate(100, 0.01)
		bf.Add([]byte("hello"))
		bf.Add([]byte("world"))

		assert.True(t, bf.Contains([]byte("hello")), "expected 'hello' to be in the Bloom filter")
		assert.True(t, bf.Contains([]byte("world")), "expected 'world' to be in the Bloom filter")
		assert.False(t, bf.Contains([]byte("foo")), "expected 'foo' to not be in the Bloom filter")
	})

	t.Run("Reset", func(t *testing.T) {
		t.Parallel()

		bf := bloomfilter.NewBloomFilter(100, 3)
		bf.Add([]byte("hello"))
		bf.Add([]byte("world"))

		bf.Reset()

		assert.False(t, bf.Contains([]byte("hello")), "expected 'hello' to not be in the Bloom filter after reset")
		assert.False(t, bf.Contains([]byte("world")), "expected 'world' to not be in the Bloom filter after reset")
	})

	t.Run("GetOptimalHasherCount", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			n        int
			m        int
			expected int
		}{
			{100, 1000, 7},
			{1000, 10000, 7},
			{10000, 100000, 7},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, bloomfilter.GetOptimalHasherCountWithExpectedItemsAndOptimalArraySize(test.n, test.m), "unexpected number of hash functions")
		}
	})

	t.Run("GetOptimalArraySize", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			n        int
			p        float64
			expected int
		}{
			{100, 0.01, 958},
			{1000, 0.01, 9585},
			{10000, 0.01, 95850},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, bloomfilter.GetOptimalArraySizeWithExpectedItemsAndFalsePositiveProbability(test.n, test.p), "unexpected size of the bit array")
		}
	})
}
