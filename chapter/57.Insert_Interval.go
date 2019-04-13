package chapter

func insert(intervals []Interval, newInterval Interval) []Interval {
	starts := func(p1, p2 *Interval) bool {
		return p1.Start < p2.Start
	}
	By(starts).Sort(intervals)
	res := make([]Interval, 0, len(intervals))
	i := 0
	for i < len(intervals) && intervals[i].End < newInterval.Start {
		res = append(res, intervals[i])
		i++
	}
	for i < len(intervals) && intervals[i].Start <= newInterval.End {
		start, end := intervals[i].Start, intervals[i].End
		if start > newInterval.Start {
			start = newInterval.Start
		}
		if end < newInterval.End {
			end = newInterval.End
		}
		newInterval = Interval{
			Start: start,
			End:   end,
		}
		i++
	}
	res = append(res, newInterval)
	// rest
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}
	return res
}
