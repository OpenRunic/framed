package framed

import (
	"slices"
)

// SliceMap converts slices of T to slices of K
func SliceMap[T any, K any](src []T, cb func(T) K) []K {
	result := make([]K, len(src))

	for idx, item := range src {
		result[idx] = cb(item)
	}

	return result
}

// SliceKeyMap convert slices of T to map of L[M]
func SliceKeyMap[T any, L comparable, M any](src []T, cb func(T, int) (L, M)) map[L]M {
	result := make(map[L]M)

	for idx, item := range src {
		key, value := cb(item, idx)
		result[key] = value
	}

	return result
}

// SliceFilter filters slice values based on filter func
func SliceFilter[T any](src []T, cb func(T) bool) []T {
	result := make([]T, 0)

	for _, item := range src {
		if cb(item) {
			result = append(result, item)
		}
	}

	return result
}

// SlicePick picks slice values
func SlicePick[T comparable](src []T, keys []T) []T {
	return SliceFilter(src, func(t T) bool {
		return slices.Contains(keys, t)
	})
}

// SlicePick omits slice values
func SliceOmit[T comparable](src []T, keys []T) []T {
	return SliceFilter(src, func(t T) bool {
		return !slices.Contains(keys, t)
	})
}
