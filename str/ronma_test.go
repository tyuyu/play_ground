package str

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "case2",
			args: args{
				num: 4,
			},
			want: "IV",
		}, {
			name: "case9",
			args: args{
				num: 9,
			},
			want: "IX",
		}, {
			name: "case4",
			args: args{
				num: 58,
			},
			want: "LVIII",
		}, {
			name: "case5",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "case6",
			args: args{
				num: 10,
			},
			want: "X",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
