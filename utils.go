package framed

import (
	"reflect"
	"slices"
)

// SliceMap converts slices of T to slices of K
func SliceMap[T any, K any](src []T, cb func(int, T) K) []K {
	result := make([]K, len(src))

	for idx, item := range src {
		result[idx] = cb(idx, item)
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

// SliceUnique picks only unique values
func SliceUnique[T comparable](src []T) []T {
	values := make([]T, 0)
	found := make(map[T]bool)

	for _, item := range src {
		if !found[item] {
			found[item] = true
			values = append(values, item)
		}
	}

	return values
}

// MapReduce reduces map to a final value
func MapReduce[T comparable, K any, L any](src map[T]K, initial L, cb func(L, T, K) L) L {
	value := initial

	for key, item := range src {
		value = cb(value, key, item)
	}

	return value
}

// IsEmpty checks if the value is empty/nil
func IsEmpty(def *Definition, value any) bool {
	if def.EmptyChecker != nil {
		return def.EmptyChecker(value)
	}

	chk, ok := value.(IsEmptyChecker)
	if ok {
		return chk.IsEmptyCheck()
	}

	if def.Kind() == reflect.Ptr {
		return value == nil
	}

	if def.Kind() == reflect.Struct {
		return false
	}

	if def.Kind() == reflect.String {
		if value == nil {
			return true
		}

		return len(value.(string)) == 0
	}

	if slices.Contains(StringTranslatableKinds, def.Kind()) {
		return false
	}

	return true
}

// MapTable loops through every row for data generation
func MapTable[T any](t *Table, cb func(*Row) T) []T {
	result := make([]T, 0)

	for _, row := range t.Rows {
		result = append(result, cb(row))
	}

	return result
}

// ReduceTable loops through every row for aggregated data generation
func ReduceTable[T any](t *Table, initial T, cb func(T, *Row) T) T {
	value := initial

	for _, row := range t.Rows {
		value = cb(value, row)
	}

	return value
}
