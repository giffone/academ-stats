package sorting

import "cmp"

func MergeSortInt[T cmp.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSortInt(arr[:mid])
	right := MergeSortInt(arr[mid:])

	return merge(left, right)
}

func merge[T cmp.Ordered](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}