package main

import "fmt"

type TreeNode struct {
	lchild *TreeNode
	rchild *TreeNode
	num    int
}

func BuildTree(a []int, n *int) *TreeNode {
	if *n > len(a) {
		return nil
	}
	if a[*n] != -1 {
		Root := &TreeNode{
			num: a[*n],
		}
		*n++
		Root.lchild = BuildTree(a, n)
		Root.rchild = BuildTree(a, n)
		return Root
	} else {
		*n++
		return nil
	}
}
func printTree(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.num)
		printTree(root.lchild)
		printTree(root.rchild)
	} else {
		return
	}
}
func main() {
	var x int = 0
	var a = []int{1, 2, -1, -1, 3, -1, -1}
	t := BuildTree(a, &x)
	printTree(t)
}
