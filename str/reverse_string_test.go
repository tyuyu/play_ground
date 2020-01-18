package str

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "111",
			args: args{
				s: []byte{'h', 'e', 'l', 'l', 'o'},
			},
		},
		{
			name: "111",
			args: args{
				s: []byte{'h', 'e', 'l', 'l', 'o', 'x'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
		})
	}
}
