package graph

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{0, 1}},
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
		{
			name: "case4",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{1, 0}, {1, 2}, {0, 1}},
			},
			want: false,
		},
		{
			name: "case5",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
