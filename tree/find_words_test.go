package tree

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				board: [][]byte{
					{'o', 'a', 'a', 'n'},
					{'e', 't', 'a', 'e'},
					{'i', 'h', 'k', 'r'},
					{'i', 'f', 'l', 'v'},
				},
				words: []string{"oath", "pea", "eat", "rain"},
			},
			want: []string{"eat", "oath"},
		}, {
			name: "case1",
			args: args{
				board: [][]byte{
					{'b', 'a', 'a', 'b', 'a', 'b'},
					{'a', 'b', 'a', 'a', 'a', 'a'},
					{'a', 'b', 'a', 'a', 'a', 'b'},
					{'a', 'b', 'a', 'b', 'b', 'a'},
					{'a', 'a', 'b', 'b', 'a', 'b'},
					{'a', 'a', 'b', 'b', 'b', 'a'},
					{'a', 'a', 'b', 'a', 'a', 'b'}},
				words: []string{"bbaabaabaaaaabaababaaaaababb",
					"aabbaaabaaabaabaaaaaabbaaaba",
					"babaababbbbbbbaabaababaabaaa",
					"bbbaaabaabbaaababababbbbbaaa",
					"babbabbbbaabbabaaaaaabbbaaab",
					"bbbababbbbbbbababbabbbbbabaa",
					"babababbababaabbbbabbbbabbba",
					"abbbbbbaabaaabaaababaabbabba",
					"aabaabababbbbbbababbbababbaa",
					"aabbbbabbaababaaaabababbaaba",
					"ababaababaaabbabbaabbaabbaba",
					"abaabbbaaaaababbbaaaaabbbaab",
					"aabbabaabaabbabababaaabbbaab",
					"baaabaaaabbabaaabaabababaaaa",
					"aaabbabaaaababbabbaabbaabbaa",
					"aaabaaaaabaabbabaabbbbaabaaa",
					"abbaabbaaaabbaababababbaabbb",
					"baabaababbbbaaaabaaabbababbb",
					"aabaababbaababbaaabaabababab",
					"abbaaabbaabaabaabbbbaabbbbbb",
					"aaababaabbaaabbbaaabbabbabab",
					"bbababbbabbbbabbbbabbbbbabaa",
					"abbbaabbbaaababbbababbababba",
					"bbbbbbbabbbababbabaabababaab",
					"aaaababaabbbbabaaaaabaaaaabb",
					"bbaaabbbbabbaaabbaabbabbaaba",
					"aabaabbbbaabaabbabaabababaaa", "abbababbbaababaabbababababbb", "aabbbabbaaaababbbbabbababbbb", "babbbaabababbbbbbbbbaabbabaa"},
			},
			want: []string{"aabbbbabbaababaaaabababbaaba", "abaabbbaaaaababbbaaaaabbbaab", "ababaababaaabbabbaabbaabbaba"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSSS(t *testing.T) {

	checkInclusion("ab", "eidboaoo")
}

func checkInclusion(s1 string, s2 string) bool {

	return false
}
