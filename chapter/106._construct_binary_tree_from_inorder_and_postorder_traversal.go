package chapter

//中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]
//返回如下的二叉树
//   3
//   / \
//  9  20
//    /  \
//   15   7

// 包内已经定义了这个函数了，所以加了个2的后缀，粘到leetcode后需要去掉2
func buildTree_2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{inorder[0], nil, nil}
	}
	index := 0
	for index = range inorder {
		if inorder[index] == postorder[len(postorder)-1] {
			break
		}
	}
	return &TreeNode{inorder[index], buildTree(inorder[:index], postorder[:index]),
	buildTree(inorder[index+1:], postorder[index:len(postorder)-1])}
}
