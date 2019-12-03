package number

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 0,
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 5,
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 20,
			},
			want: false,
		},
		{
			name: "case4",
			args: args{
				matrix: [][]int{
					{-1, 3},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
				},
				target: 19,
			},
			want: true,
		},
		{
			name: "case6",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7, 9},
					{2, 4, 6, 8, 10},
					{11, 13, 15, 17, 19},
					{12, 14, 16, 18, 20},
					{21, 22, 23, 24, 25},
				},
				target: 11,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
