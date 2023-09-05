package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		digits []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{digits: []int{1, 8, 3, 2}},
			want: 6,
		},
		{
			name: "test2",
			args: args{[]int{2, 3, 3, 2}},
			want: 3,
		}, {
			name: "test3",
			args: args{[]int{6, 2, 4, 7}},
			want: 0,
		}, {
			name: "test4",
			args: args{[]int{9, 9, 9, 9, 9}},
			want: 0,
		}, {
			name: "test5",
			args: args{[]int{-9, -1, -2, 0, 10}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.digits); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
