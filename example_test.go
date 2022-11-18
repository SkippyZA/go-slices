package slices_test

import (
	"fmt"
	"strconv"

	"github.com/skippyza/go-slices"
)

func ExampleAll() {
	isEven := func(i int) bool { return i%2 == 0 }

	result := slices.All([]int{2, 4, 6, 8, 10}, isEven)
	fmt.Println(result)
	result = slices.All([]int{1, 2, 4, 6, 8, 9, 10}, isEven)
	fmt.Println(result)

	// Output:
	// true
	// false
}

func ExampleAny() {
	isEven := func(i int) bool { return i%2 == 0 }

	result := slices.Any([]int{1, 2, 3, 4, 5}, isEven)
	fmt.Println(result)
	result = slices.Any([]int{1, 3, 5}, isEven)
	fmt.Println(result)

	// Output:
	// true
	// false
}

func ExampleEach() {
	printLine := func(s string) { fmt.Println(s) }

	slices.Each([]string{"test 1", "test 2", "test 3"}, printLine)

	// Output:
	// test 1
	// test 2
	// test 3
}

func ExampleFilter() {
	isEven := func(i int) bool { return i%2 == 0 }
	result := slices.Filter([]int{1, 2, 3, 4, 5}, isEven)
	fmt.Println(result)

	// Output:
	// [2 4]
}

func ExampleFind() {
	isEven := func(i int) bool { return i%2 == 0 }

	result, err := slices.Find([]int{1, 3, 4, 6, 7}, isEven)
	fmt.Println(result, err)
	result, err = slices.Find([]int{1, 3, 5, 7}, isEven)
	fmt.Println(result, err)

	// Output:
	// 4 <nil>
	// 0 item not found
}

func ExampleMap() {
	withPrefix := func(i int) string { return "pre-" + strconv.Itoa(i) }

	result := slices.Map([]int{1, 2, 3}, withPrefix)
	fmt.Println(result)

	// Output:
	// [pre-1 pre-2 pre-3]
}

func ExampleReduce() {
	sum := func(last, cur int) int { return last + cur }

	result := slices.Reduce([]int{2, 4, 6, 8}, sum, 0)
	fmt.Println(result)

	// Output:
	// 20
}
