package m

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{
		//	name: "case1",
		//	args: args{
		//		s: "aa",
		//		p: "a",
		//	},
		//	want: false,
		//},
		//{
		//	name: "case2",
		//	args: args{
		//		s: "aa",
		//		p: "*",
		//	},
		//	want: true,
		//},
		//{
		//	name: "case3",
		//	args: args{
		//		s: "cb",
		//		p: "?b",
		//	},
		//	want: true,
		//},
		//{
		//	name: "case4",
		//	args: args{
		//		s: "cb",
		//		p: "?a",
		//	},
		//	want: false,
		//},
		//{
		//	name: "case5",
		//	args: args{
		//		s: "adceb",
		//		p: "*a*b",
		//	},
		//	want: true,
		//},
		//{
		//	name: "case6",
		//	args: args{
		//		s: "acdcb",
		//		p: "a*c?b",
		//	},
		//	want: false,
		//},
		{
			name: "case7",
			args: args{
				s: "abefcdgiescdfimde",
				p: "ab*cd?i*de",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
