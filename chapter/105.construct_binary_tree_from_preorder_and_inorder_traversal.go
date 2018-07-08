package chapter

// 取先序的第一个数作为跟，然后继续定位左子树的区域
// 方法是把preorder后面的元素和inorder开头的元素进行一一比对，索引一起往后+直到发现不相等，确定这个不相等元素和preorder的0号元素相等，得到两个索引值，作为分化，出三个partition
// 剩下就是递归构建一颗树的过程了
// 题目假设了所有元素都不会相等，但是实际上假如元素重复的话，无法确认根，会使问题变得复杂,假如重复怎么办？可能需要增加个方法，来判断是否能构成一课树，遍历的同时对左边和右边同时校验，可能出现多解

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}
	index := 0
	for index = range inorder {
		//找根元素
		if inorder[index] == preorder[0] {
			break
		}
	}

	return &TreeNode{preorder[0], buildTree(preorder[1:index+1], inorder[:index]),
	buildTree(preorder[index+1:], inorder[index+1:])}
}