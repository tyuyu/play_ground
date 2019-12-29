package str

import (
	"reflect"
	"testing"
)

func Test_wordBreaks(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"case1",
			args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			[]string{"cats and dog", "cat sand dog"},
		},
		{
			"case2",
			args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			[]string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			"case3",
			args{
				s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreaks(tt.args.s, tt.args.wordDict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreaks() = %v, want %v", got, tt.want)
			} else {
				t.Log("got is ", got)
			}

		})
	}
}
