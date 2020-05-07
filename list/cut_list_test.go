package list

import (
	"reflect"
	"testing"
)

//给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
//
//你应当保留两个分区中每个节点的初始相对位置。
//
//示例:
//
//输入: head = 1->4->3->2->5->2, x = 3
//输出: 1->2->2->4->3->5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/partition-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val:  2,
										Next: nil,
									},
								},
							},
						},
					},
				},
				x: 3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
