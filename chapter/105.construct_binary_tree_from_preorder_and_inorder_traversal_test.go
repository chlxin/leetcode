package chapter

import (
	"testing"
	"fmt"
)

//
//func Test_buildTree(t *testing.T) {
//	type args struct {
//		preorder []int
//		inorder  []int
//	}
//	tests := []struct {
//		name string
//		args args
//		want *TreeNode
//	}{
//
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("buildTree() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}


func TestBuildTree(t *testing.T) {
	tree := buildTree([]int{1, 2, 3}, []int{3, 2, 1})
	fmt.Println(tree)
}