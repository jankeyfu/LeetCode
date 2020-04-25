package competition

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		n        int
		relation [][]int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:        5,
				relation: [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}},
				k:        3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n, tt.args.relation, tt.args.k); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
