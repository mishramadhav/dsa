package binarysearch

import "cmp"

// Find returns the index of the target value in the sorted slice.
// If the target value is not found, it returns the index where the target value should be inserted
// and the second return value is false.
// The slice must be sorted in non-decreasing order.
func Find[T cmp.Ordered](arr []T, target T) (int, bool) {
	idx := LowerBound(arr, target)
	if idx < len(arr) && arr[idx] == target {
		return idx, true
	}

	return idx, false
}

// LowerBound returns the first index of the target value in the sorted slice.
// If the target value is not found, it returns the index where the target value should be inserted.
// The slice must be sorted in non-decreasing order.
func LowerBound[T cmp.Ordered](arr []T, target T) int {
	l, r := 0, len(arr)
	for l < r {
		m := l + (r-l)/2
		if cmp.Less(arr[m], target) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

// UpperBound returns the last index of the target value in the sorted slice.
// If the target value is not found, it returns the index where the target value should be inserted.
// The slice must be sorted in non-decreasing order.
func UpperBound[T cmp.Ordered](arr []T, target T) int {
	l, r := 0, len(arr)
	for l < r {
		m := l + (r-l)/2
		if !cmp.Less(target, arr[m]) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

// EqualRange returns the range of indices where the target value is located in the sorted slice.
// If the target value is not found, it returns the range where the target value should be inserted.
// The slice must be sorted in non-decreasing order.
func EqualRange[T cmp.Ordered](arr []T, target T) (int, int) {
	return LowerBound(arr, target), UpperBound(arr, target)
}
