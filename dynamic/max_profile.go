package dynamic

func maxProfit(prices []int) int {

	if len(prices) < 4 {
		return 0
	}

	max := 0
	for i := 2; i < len(prices)-1; i++ {
		l := maxp(prices[:i])
		if l > 0 {
			r := maxp(prices[i:])
			if r > 0 && l+r > max {
				max = l + r
			}
		}
	}
	return max
}

func maxp(prices []int) int {
	min := prices[0]
	max := 0
	for i, p := range prices {
		if i == 0 {
			min = p
			continue
		}
		if p < min {
			min = p
			continue
		}
		if p-min > max {
			max = p - min
		}
	}
	return max
}

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}

	search := make([][]bool, m)
	for i := range search {
		search[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if search[i][j] {
				continue
			}
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				if board[i][j] != 'O' {
					continue
				}
				mark(board, search, i, j, m, n)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func mark(board [][]byte, search [][]bool, i, j, m, n int) {

	if i < 0 || j < 0 || i >= m || n >= n {
		return
	}

	if search[i][j] {
		return
	}

	search[i][j] = true

	if board[i][j] == 'O' {
		board[i][j] = '#'
	}

	mark(board, search, i-1, j, m, n)
	mark(board, search, i+1, j, m, n)
	mark(board, search, i, j-1, m, n)
	mark(board, search, i, j+1, m, n)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	set := make(map[int]*Node, 0)
	check := make(map[int]bool, 0)
	copyNode(node, set, check)
	return set[1]
}

func copyNode(node *Node, set map[int]*Node, check map[int]bool) {

	if check[node.Val] {
		return
	}

	check[node.Val] = true

	cp := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	if n, ok := set[node.Val]; ok {
		cp = n
	}
	for _, nn := range node.Neighbors {
		cc := nn
		cpnn, ok := set[nn.Val]
		if !ok {
			cpnn = &Node{
				Val:       nn.Val,
				Neighbors: make([]*Node, 0),
			}
			set[nn.Val] = cpnn
		}
		cp.Neighbors = append(cp.Neighbors, cpnn)
		copyNode(cc, set, check)
	}
}
