package graph

//单词接龙
//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
//
//每次转换只能改变一个字母。
//转换过程中的中间单词必须是字典中的单词。
//说明:
//
//如果不存在这样的转换序列，返回 0。
//所有单词具有相同的长度。
//所有单词只由小写字母组成。
//字典中不存在重复的单词。
//你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
//示例 1:
//
//输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     返回它的长度 5。
//示例 2:
//
//输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: 0
//
//解释: endWord "cog" 不在字典中，所以无法进行转换。

func ladderLength(beginWord string, endWord string, wordList []string) int {

	if len(beginWord) != len(endWord) {
		return 0
	}

	wordSet := make(map[string]struct{}, 0)
	for _, s := range wordList {
		wordSet[s] = struct{}{}
	}

	if _, ok := wordSet[endWord]; !ok {
		return 0
	}
	delete(wordSet, beginWord)
	delete(wordSet, endWord)
	findSet, resultSet := map[string]struct{}{beginWord: {}}, map[string]struct{}{endWord: {}}
	level := 1
	for len(findSet) > 0 && len(resultSet) > 0 {
		if len(findSet) > len(resultSet) {
			findSet, resultSet = resultSet, findSet
		}
		tmp := make(map[string]struct{})
		for f, _ := range findSet {
			fb := []byte(f)
			for i := range fb {
				for j := 'a'; j <= 'z'; j++ {
					fb[i] = byte(j)
					str := string(fb)
					if _, ok := resultSet[str]; ok {
						return level + 1
					}
					if _, ok := wordSet[str]; ok {
						delete(wordSet, str)
						tmp[str] = struct{}{}
					}
				}
				fb[i] = f[i]
			}
		}
		level++
		findSet = tmp
	}

	return 0
}
