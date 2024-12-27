package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serialize(root *TreeNode) string {
	var result []string
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			result = append(result, "#")
			return
		}
		result = append(result, strconv.Itoa(node.Val))
		helper(node.Left)
		helper(node.Right)
	}
	helper(root)
	return strings.Join(result, ",")
}

func deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	var index int
	var helper func() *TreeNode
	helper = func() *TreeNode {
		if index >= len(values) {
			return nil
		}
		if values[index] == "#" {
			index++
			return nil
		}

		val, _ := strconv.Atoi(values[index])
		index++

		node := &TreeNode{Val: val}
		node.Left = helper()
		node.Right = helper()
		return node
	}

	return helper()
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
	}

	serialized := serialize(root)
	fmt.Println("Serialized tree:", serialized)

	deserialized := deserialize(serialized)
	fmt.Println("Deserialized tree (root value):", deserialized.Val)
}
