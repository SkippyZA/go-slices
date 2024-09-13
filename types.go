package goslices

type PredicateFn[T any] func(T) bool
type MapFn[T any, U any] func(T) U
type AccumulatorFn[T any, U any] func(U, T) U
