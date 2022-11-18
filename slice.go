package slices

import "errors"

var (
	ErrNotFound = errors.New("item not found")
)

type PredicateFn[T any] func(T) bool
type MapFn[T any, U any] func(T) U
type AccumulatorFn[T any, U any] func(U, T) U

func All[S ~[]T, T any](arr S, fn PredicateFn[T]) bool {
	for _, el := range arr {
		if !fn(el) {
			return false
		}
	}

	return true
}

func Any[S ~[]T, T any](arr S, fn PredicateFn[T]) bool {
	for _, el := range arr {
		if fn(el) {
			return true
		}
	}

	return false
}

func Each[S ~[]T, T any](arr S, fn func(T)) {
	for _, el := range arr {
		fn(el)
	}
}

func Filter[S ~[]T, T any](arr S, fn PredicateFn[T]) S {
	result := make(S, 0, len(arr))
	for _, el := range arr {
		if fn(el) {
			result = append(result, el)
		}
	}

	return result
}

func Find[S ~[]T, T any](arr S, fn PredicateFn[T]) (T, error) {
	for _, el := range arr {
		if fn(el) {
			return el, nil
		}
	}

	var tmp T
	return tmp, ErrNotFound
}

func Map[S ~[]T, T any, U any](arr S, fn MapFn[T, U]) []U {
	result := make([]U, 0, len(arr))
	for _, el := range arr {
		result = append(result, fn(el))
	}

	return result
}

func Reduce[S ~[]T, T any, U any](arr S, fn AccumulatorFn[T, U], initial U) U {
	for _, el := range arr {
		initial = fn(initial, el)
	}

	return initial
}
