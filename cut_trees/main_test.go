package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{A: []int{3, 4, 5, 3, 7}},
			want: 3,
		}, {
			name: "case2",
			args: args{A: []int{1, 3, 1, 2}},
			want: 0,
		}, {
			name: "case3",
			args: args{[]int{1, 2, 3, 4}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
