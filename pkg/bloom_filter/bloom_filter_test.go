package bloomfilter_test

import (
	"testing"

	bloomfilter "github.com/mishramadhav/dsa/pkg/bloom_filter"
)

func TestBloomFilter(t *testing.T) {
	t.Parallel()

	t.Run("Add and Contains", func(t *testing.T) {
		t.Parallel()

		bf := bloomfilter.NewBloomFilter(100, 3)
		bf.Add([]byte("hello"))
		bf.Add([]byte("world"))

		if !bf.Contains([]byte("hello")) {
			t.Error("expected 'hello' to be in the Bloom filter")
		}
		if !bf.Contains([]byte("world")) {
			t.Error("expected 'world' to be in the Bloom filter")
		}
		if bf.Contains([]byte("foo")) {
			t.Error("expected 'foo' to not be in the Bloom filter")
		}
	})

	t.Run("NewBloomFilterWithAllowedFalsePositiveRate", func(t *testing.T) {
		t.Parallel()

		bf := bloomfilter.NewBloomFilterWithAllowedFalsePositiveRate(100, 0.01)
		bf.Add([]byte("hello"))
		bf.Add([]byte("world"))

		if !bf.Contains([]byte("hello")) {
			t.Error("expected 'hello' to be in the Bloom filter")
		}
		if !bf.Contains([]byte("world")) {
			t.Error("expected 'world' to be in the Bloom filter")
		}
		if bf.Contains([]byte("foo")) {
			t.Error("expected 'foo' to not be in the Bloom filter")
		}
	})

	t.Run("Reset", func(t *testing.T) {
		t.Parallel()

		bf := bloomfilter.NewBloomFilter(100, 3)
		bf.Add([]byte("hello"))
		bf.Add([]byte("world"))

		bf.Reset()

		if bf.Contains([]byte("hello")) {
			t.Error("expected 'hello' to not be in the Bloom filter after reset")
		}
		if bf.Contains([]byte("world")) {
			t.Error("expected 'world' to not be in the Bloom filter after reset")
		}
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
			got := bloomfilter.GetOptimalHasherCountWithExpectedItemsAndOptimalArraySize(test.n, test.m)
			if got != test.expected {
				t.Errorf("unexpected number of hash functions: got=%d, expected=%d", got, test.expected)
			}
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
			got := bloomfilter.GetOptimalArraySizeWithExpectedItemsAndFalsePositiveProbability(test.n, test.p)
			if got != test.expected {
				t.Errorf("unexpected size of the bit array: got=%d, expected=%d", got, test.expected)
			}
		}
	})
}
