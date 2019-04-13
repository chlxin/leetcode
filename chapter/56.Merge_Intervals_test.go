package chapter

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	intervals := []Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 2, End: 6},
		Interval{Start: 8, End: 10},
		Interval{Start: 15, End: 18},
	}
	res := merge(intervals)
	fmt.Println(res)
}
