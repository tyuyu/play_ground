package dynamic

import "testing"

func Test_longestIncreasingPath(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				matrix: [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}},
			},
			want: 4,
		}, {
			name: "cas21",
			args: args{
				matrix: [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
