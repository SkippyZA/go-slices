package slices

import "errors"

var (
	ErrNotFound = errors.New("item not found")
)

type PredicateFn[T any] func(T) bool
type MapFn[IN any, OUT any] func(IN) OUT
type AccumulatorFn[IN any, OUT any] func(OUT, IN) OUT

func All[S ~[]IN, IN any](arr S, fn PredicateFn[IN]) bool {
	for _, el := range arr {
		if !fn(el) {
			return false
		}
	}

	return true
}

func Any[S ~[]IN, IN any](arr S, fn PredicateFn[IN]) bool {
	for _, el := range arr {
		if fn(el) {
			return true
		}
	}

	return false
}

func Each[S ~[]IN, IN any](arr S, fn func(IN)) {
	for _, el := range arr {
		fn(el)
	}
}

func Filter[S ~[]IN, IN any](arr S, fn PredicateFn[IN]) S {
	result := make(S, 0, len(arr))
	for _, el := range arr {
		if fn(el) {
			result = append(result, el)
		}
	}

	return result
}

func Find[S ~[]IN, IN any](arr S, fn PredicateFn[IN]) (IN, error) {
	for _, el := range arr {
		if fn(el) {
			return el, nil
		}
	}

	var tmp IN
	return tmp, ErrNotFound
}

func Map[S ~[]IN, IN any, OUT any](arr S, fn MapFn[IN, OUT]) []OUT {
	result := make([]OUT, 0, len(arr))
	for _, el := range arr {
		result = append(result, fn(el))
	}

	return result
}

func Reduce[S ~[]IN, IN any, OUT any](arr S, fn AccumulatorFn[IN, OUT], initial OUT) OUT {
	for _, el := range arr {
		initial = fn(initial, el)
	}

	return initial
}
