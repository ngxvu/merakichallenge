package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_1",
			args: args{
				S: "nice",
				T: "nicer",
			},
			want: "ADD r",
		},
		{
			name: "test_2",
			args: args{
				S: "test",
				T: "tent",
			},
			want: "CHANGE s n",
		},
		{
			name: "test_3",
			args: args{
				S: "beans",
				T: "banes",
			},
			want: "MOVE e",
		},
		{
			name: "test_4",
			args: args{
				S: "0",
				T: "odd",
			},
			want: "IMPOSSIBLE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
