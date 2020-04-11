package tree

import (
	"reflect"
	"testing"
)

func TestArrayToTree(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case1",
			args: args{
				a: "3,5,1,6,2,0,8,null,null,7,4",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToTree(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayToTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
