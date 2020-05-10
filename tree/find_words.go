package tree

import "sort"

//给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
//
//示例:
//
//输入:
//words = ["oath","pea","eat","rain"] and board =
//[
//  ['o','a','a','n'],
//  ['e','t','a','e'],
//  ['i','h','k','r'],
//  ['i','f','l','v']
//]
//
//输出: ["eat","oath"]
//说明:
//你可以假设所有输入都由小写字母 a-z 组成。
//
//提示:
//
//你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
//如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？ 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。

func findWords(board [][]byte, words []string) []string {

	res := make([]string, 0)

	t := trie{
		isEnd: false,
		next:  make(map[rune]*trie, 0),
	}

	for _, w := range words {
		t.insert(w)
	}

	yl := len(board)
	if yl == 0 {
		return res
	}
	xl := len(board[0])
	if xl == 0 {
		return res
	}
	for y := 0; y < yl; y++ {
		for x := 0; x < xl; x++ {
			bg := &node{
				word:    string(board[y][x]),
				p:       pos{x, y},
				parents: make(map[pos]*node, 0),
				board:   &board,
				res:     &res,
				t:       &t,
			}
			bg.check()
		}
	}

	tmp := make([]string, 0)
	ccc := make(map[string]struct{}, 0)
	for _, r := range res {
		_, ok := ccc[r]
		if !ok {
			ccc[r] = struct{}{}
			tmp = append(tmp, r)
		}
	}
	sort.Strings(tmp)
	return tmp
}

type pos struct {
	x, y int
}

type node struct {
	word    string
	p       pos
	parents map[pos]*node
	board   *[][]byte
	res     *[]string
	t       *trie
}

func (n *node) neighbor(xl, yl int) []pos {
	neighbors := make([]pos, 0)
	now := n.p
	if now.x-1 >= 0 {
		p := pos{now.x - 1, now.y}
		_, ok := n.parents[p]
		if !ok {
			neighbors = append(neighbors, p)
		}
	}
	if now.y-1 >= 0 {
		p := pos{now.x, now.y - 1}
		_, ok := n.parents[p]
		if !ok {
			neighbors = append(neighbors, p)
		}
	}
	if now.x+1 < xl {
		p := pos{now.x + 1, now.y}
		_, ok := n.parents[p]
		if !ok {
			neighbors = append(neighbors, p)
		}
	}
	if now.y+1 < yl {
		p := pos{now.x, now.y + 1}
		_, ok := n.parents[p]
		if !ok {
			neighbors = append(neighbors, p)
		}
	}
	return neighbors
}

func (n *node) check() {
	if !n.t.startsWith(n.word) {
		return
	}
	if n.t.search(n.word) {
		*n.res = append(*n.res, n.word)
	}
	yl := len(*n.board)
	xl := len((*n.board)[0])
	nexts := n.neighbor(xl, yl)
	p := n.parents
	p[n.p] = n
	for _, next := range nexts {
		node := node{
			word:    n.word + string((*n.board)[next.y][next.x]),
			p:       next,
			parents: p,
			board:   n.board,
			res:     n.res,
			t:       n.t,
		}
		node.check()
	}
}

type trie struct {
	isEnd bool
	next  map[rune]*trie
}

/** Inserts a word into the trie. */
func (t *trie) insert(word string) {
	cur := t
	for _, r := range word {
		n, ok := cur.next[r]
		if ok {
			cur = n
		} else {
			n = &trie{
				isEnd: false,
				next:  make(map[rune]*trie),
			}
			cur.next[r] = n
		}
		cur = n
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (t *trie) search(word string) bool {
	cur := t
	for _, r := range word {
		n, ok := cur.next[r]
		if !ok {
			return false
		}
		cur = n
	}
	return cur.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *trie) startsWith(prefix string) bool {
	cur := t
	for _, r := range prefix {
		n, ok := cur.next[r]
		if !ok {
			return false
		}
		cur = n
	}
	return true
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) / 2
	root := &TreeNode{
		Val: nums[m],
	}
	left := nums[:m]
	right := nums[m+1:]
	it := root
	for i := len(left) - 1; i >= 0; i-- {
		it.Left = &TreeNode{
			Val: left[i],
		}
		it = it.Left
	}
	it = root
	for _, n := range right {
		it.Right = &TreeNode{
			Val: n,
		}
		it.Right = it
	}
	return root
}
