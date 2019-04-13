package chapter

import (
	"fmt"
	"testing"
)

func Test_insert(t *testing.T) {
	intervals := []Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 6, End: 9},
		// Interval{Start: 8, End: 10},
		// Interval{Start: 15, End: 18},
	}
	res := insert(intervals, Interval{Start: 2, End: 5})
	fmt.Println(res)
}
