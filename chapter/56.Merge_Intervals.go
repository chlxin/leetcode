package chapter

import "sort"

// Interval 时间间隔
type Interval struct {
	Start int
	End   int
}

// IntervalSorter 排序
type IntervalSorter struct {
	intervals []Interval
	by        func(p1, p2 *Interval) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *IntervalSorter) Len() int {
	return len(s.intervals)
}

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *Interval) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(intervals []Interval) {
	ps := &IntervalSorter{
		intervals: intervals,
		by:        by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// Swap is part of sort.Interface.
func (s *IntervalSorter) Swap(i, j int) {
	s.intervals[i], s.intervals[j] = s.intervals[j], s.intervals[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *IntervalSorter) Less(i, j int) bool {
	return s.by(&s.intervals[i], &s.intervals[j])
}

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

func merge(intervals []Interval) []Interval {
	starts := func(p1, p2 *Interval) bool {
		return p1.Start < p2.Start
	}
	By(starts).Sort(intervals)
	res := make([]Interval, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		for i+1 < len(intervals) && interval.End >= intervals[i+1].Start {
			end := interval.End
			if end < intervals[i+1].End {
				end = intervals[i+1].End
			}
			interval = Interval{
				Start: interval.Start,
				End:   end,
			}
			i++
		}
		res = append(res, interval)
	}
	return res
}
