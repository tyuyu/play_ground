package str

import "strings"

//给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
//
//你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
//
//要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
//
//文本的最后一行应为左对齐，且单词之间不插入额外的空格。
//
//说明:
//
//单词是指由非空格字符组成的字符序列。
//每个单词的长度大于 0，小于等于 maxWidth。
//输入单词数组 words 至少包含一个单词。
//示例:
//
//输入:
//words = ["This", "is", "an", "example", "of", "text", "justification."]
//maxWidth = 16
//输出:
//[
//   "This    is    an",
//   "example  of text",
//   "justification.  "
//]
//示例 2:
//
//输入:
//words = ["What","must","be","acknowledgment","shall","be"]
//maxWidth = 16
//输出:
//[
//  "What   must   be",
//  "acknowledgment  ",
//  "shall be        "
//]
//解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
//     因为最后一行应为左对齐，而不是左右两端对齐。
//     第二行同样为左对齐，这是因为这行只包含一个单词。
//示例 3:
//
//输入:
//words = ["Science","is","what","we","understand","well","enough","to","explain",
//         "to","a","computer.","Art","is","everything","else","we","do"]
//maxWidth = 20
//输出:
//[
//  "Science  is  what we",
//  "understand      well",
//  "enough to explain to",
//  "a  computer.  Art is",
//  "everything  else  we",
//  "do                  "
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/text-justification
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	tmp := make([]string, 0)
	for i := 0; i < len(words); i++ {
		old := strings.Join(tmp, " ")
		if len(old)+len(words[i]) < maxWidth {
			tmp = append(tmp, words[i])
			if i != len(words)-1 {
				continue
			}
		}
		wc := 0
		for _, w := range tmp {
			wc += len(w)
		}
		spaces := maxWidth - wc
		if len(tmp) == 1 {
			ww := tmp[0]
			for spaces > 0 {
				ww = ww + " "
				spaces--
			}
			res = append(res, ww)
		} else {
			each := spaces / (len(tmp) - 1)
			last := spaces % each
			ww := ""
			for x, w := range tmp {
				ww = ww + w
				if x == len(tmp)-1 {
					continue
				}
				for j := 0; j < each; j++ {
					ww = ww + " "
				}
				if last > 0 {
					ww = ww + " "
					last--
				}
			}
			res = append(res, ww)
		}

		tmp = []string{words[i]}
	}
	last := strings.Join(tmp, " ")
	sp := maxWidth - len(last)
	for sp > 0 {
		last += " "
		sp--
	}
	res = append(res, last)
	return res
}
