package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[null]"
	}
	res := []string{strconv.Itoa(root.Val)}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[0]
		if node != nil {
			nodes = append(nodes, node.Left, node.Right)
			if node.Left == nil {
				res = append(res, "null")
			} else {
				res = append(res, strconv.Itoa(node.Left.Val))
			}
			if node.Right == nil {
				res = append(res, "null")
			} else {
				res = append(res, strconv.Itoa(node.Right.Val))
			}
		}
		nodes = nodes[1:]
	}
	i := len(res) - 1
	for ; i > 0; i-- {
		if res[i] != "null" {
			break
		}
	}
	res = res[:i+1]
	return fmt.Sprintf("[%s]", strings.Join(res, ","))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = strings.Replace(data, "[", "", -1)
	data = strings.Replace(data, "]", "", -1)
	strs := strings.Split(data, ",")

	if len(strs) == 0 {
		return nil
	}
	root := build(strs[0])
	queue := []*TreeNode{root}
	strs = strs[1:]
	for len(queue) > 0 && len(strs) > 0 {
		if len(strs) >= 2 {
			l, r := strs[0], strs[1]
			ln, rn := build(l), build(r)
			if ln != nil {
				queue = append(queue, ln)
				queue[0].Left = ln
			}
			if rn != nil {
				queue = append(queue, rn)
				queue[0].Right = rn
			}
			queue = queue[1:]
			strs = strs[2:]
		} else {
			l := strs[0]
			ln := build(l)
			if ln != nil {
				queue = append(queue, ln)
				queue[0].Left = ln
			}
			queue = queue[1:]
			strs = strs[1:]
		}
	}

	return root
}

func build(s string) *TreeNode {
	if s == "null" {
		return nil
	}
	t := &TreeNode{}
	t.Val, _ = strconv.Atoi(s)
	return t
}
