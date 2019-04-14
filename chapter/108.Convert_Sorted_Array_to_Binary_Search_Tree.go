package chapter

func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree108(nums)
}

func buildTree108(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := (len(nums) - 1) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  buildTree108(nums[:mid]),
		Right: buildTree108(nums[mid+1:]),
	}
}
