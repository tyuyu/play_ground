package tree

//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
//示例:
//
//Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");
//trie.search("app");     // 返回 true
//说明:
//
//你可以假设所有的输入都是由小写字母 a-z 构成的。
//保证所有输入均为非空字符串。

type Trie struct {
	isEnd bool
	next  map[rune]*Trie
}

///** Initialize your data structure here. */
//func Constructor() Trie {
//	return Trie{
//		isEnd: false,
//		next:  make(map[rune]*Trie),
//	}
//}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	cur := t
	for _, r := range word {
		n, ok := cur.next[r]
		if ok {
			cur = n
		} else {
			n = &Trie{
				isEnd: false,
				next:  make(map[rune]*Trie),
			}
			cur.next[r] = n
		}
		cur = n
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
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
func (t *Trie) StartsWith(prefix string) bool {
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

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
