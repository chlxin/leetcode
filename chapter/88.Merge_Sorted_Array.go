package chapter

func merge88(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	index := m + n - 1
	for index >= 0 {
		if p1 >= 0 && p2 >= 0 {
			if nums1[p1] > nums2[p2] {
				nums1[index] = nums1[p1]
				p1--
			} else {
				nums1[index] = nums2[p2]
				p2--
			}
		} else if p1 >= 0 {
			nums1[index] = nums1[p1]
			p1--
		} else if p2 >= 0 {
			nums1[index] = nums2[p2]
			p2--
		}
		index--
	}
}
