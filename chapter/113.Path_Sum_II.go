package chapter

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	backtrack113(&res, make([]int, 0), root, sum)
	return res
}

func backtrack113(res *[][]int, nums []int, node *TreeNode, target int) {
	if node == nil {
		return
	}
	// leaf node
	if node.Left == nil && node.Right == nil {
		if node.Val == target {
			nums = append(nums, node.Val)
			tmpRes := make([]int, len(nums))
			copy(tmpRes, nums)
			*res = append(*res, tmpRes)
		}
		return
	}
	if node.Left != nil {
		nums = append(nums, node.Val)
		backtrack113(res, nums, node.Left, target-node.Val)
		nums = nums[:len(nums)-1]
	}
	if node.Right != nil {
		nums = append(nums, node.Val)
		backtrack113(res, nums, node.Right, target-node.Val)
		nums = nums[:len(nums)-1]
	}
}
