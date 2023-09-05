package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test_1",
			args: args{A: [][]int{
				{4, 3, 4, 5, 3},
				{2, 7, 3, 8, 4},
				{1, 7, 6, 5, 2},
				{8, 4, 9, 5, 5},
			}},
			want: 3,
		},
		{name: "test_2",
			args: args{A: [][]int{
				{2, 2, 1, 1},
				{2, 2, 2, 2},
				{1, 2, 2, 2},
			}},
			want: 2,
		},
		{name: "test_3",
			args: args{A: [][]int{
				{7, 2, 4},
				{2, 7, 6},
				{9, 5, 1},
				{4, 3, 8},
				{3, 5, 4},
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.A); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
