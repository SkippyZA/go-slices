package slices_test

import (
	"strconv"
	"testing"

	"github.com/skippyza/go-slices"
	"github.com/stretchr/testify/assert"
)

var (
	isPositive     = func(el int) bool { return el > 0 }
	isEven         = func(i int) bool { return i%2 == 0 }
	doubleFn       = func(i int) int { return i + i }
	appendToString = func(total string, el int) string { return total + strconv.Itoa(el) }
)

func TestAll(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := slices.All(arr, isPositive)
	assert.Equal(t, true, result)

	arr = []int{1, 2, 3, -1}
	result = slices.All(arr, isPositive)
	assert.Equal(t, false, result)
}

func TestEach(t *testing.T) {
	count := 0
	counter := func(in int) { count++ }

	arr := []int{1, 2, 3, 4}
	slices.Each(arr, counter)
	assert.Equal(t, 4, count)
}

func TestFind(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result, err := slices.Find(arr, isEven)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)

	arr = []int{1, 3, 5, 7}
	result, err = slices.Find(arr, isEven)
	assert.ErrorIs(t, err, slices.ErrNotFound)
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := slices.Map(arr, doubleFn)
	assert.Equal(t, []int{2, 4, 6, 8}, result)
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := slices.Filter(arr, isEven)
	assert.Equal(t, []int{2, 4}, result)
}

func TestReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := slices.Reduce(arr, appendToString, "")
	assert.Equal(t, "1234", result)
}

func TestAny(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := slices.Any(arr, isPositive)
	assert.Equal(t, true, result)

	arr = []int{1, 2, 3, -1}
	result = slices.Any(arr, isPositive)
	assert.Equal(t, true, result)

	arr = []int{-1, -2, -3, -4}
	result = slices.Any(arr, isPositive)
	assert.Equal(t, false, result)
}
