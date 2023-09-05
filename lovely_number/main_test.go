package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_1",
			args: args{
				A: 0,
				B: 0,
			},
			want: 1,
		},
		{
			name: "test_2",
			args: args{
				A: 1,
				B: 111,
			},
			want: 110,
		},
		{
			name: "test_3",
			args: args{
				A: 100000,
				B: 100000,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
