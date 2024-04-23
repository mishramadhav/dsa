package binarysearch

import (
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "value exists",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 3,
			},
			want:  2,
			want1: true,
		},
		{
			name: "value does not exist",
			args: args{
				arr:    []int{1, 2, 4, 5},
				target: 3,
			},
			want:  2,
			want1: false,
		},
		{
			name: "value is smaller than all",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "value is larger than all",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 6,
			},
			want:  5,
			want1: false,
		},
		{
			name: "empty slice",
			args: args{
				arr:    []int{},
				target: 3,
			},
			want:  0,
			want1: false,
		},
		{
			name: "value exists, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 3,
			},
			want:  2,
			want1: true,
		},
		{
			name: "value does not exist, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 6,
			},
			want:  7,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Find(tt.args.arr, tt.args.target)
			if got != tt.want {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLowerBound(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "value exists",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 3,
			},
			want: 2,
		},
		{
			name: "value does not exist",
			args: args{
				arr:    []int{1, 2, 4, 5},
				target: 3,
			},
			want: 2,
		},
		{
			name: "value is smaller than all",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 0,
			},
			want: 0,
		},
		{
			name: "value is larger than all",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 6,
			},
			want: 5,
		},
		{
			name: "empty slice",
			args: args{
				arr:    []int{},
				target: 3,
			},
			want: 0,
		},
		{
			name: "value exists, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 3,
			},
			want: 2,
		},
		{
			name: "value does not exist, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 6,
			},
			want: 7,
		},
		{
			name: "value is smaller than all, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 0,
			},
			want: 0,
		},
		{
			name: "value is larger than all, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 6,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerBound(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("LowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "value exists",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 3,
			},
			want: 3,
		},
		{
			name: "value does not exist",
			args: args{
				arr:    []int{1, 2, 4, 5},
				target: 3,
			},
			want: 2,
		},
		{
			name: "value is smaller than all",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 0,
			},
			want: 0,
		},
		{
			name: "value is larger than all",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 6,
			},
			want: 5,
		},
		{
			name: "empty slice",
			args: args{
				arr:    []int{},
				target: 3,
			},
			want: 0,
		},
		{
			name: "value exists, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 3,
			},
			want: 5,
		},
		{
			name: "value does not exist, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 6,
			},
			want: 7,
		},
		{
			name: "value is smaller than all, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 0,
			},
			want: 0,
		},
		{
			name: "value is larger than all, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 6,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperBound(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("UpperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualRange(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "value exists",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 3,
			},
			want:  2,
			want1: 3,
		},
		{
			name: "value does not exist",
			args: args{
				arr:    []int{1, 2, 4, 5},
				target: 3,
			},
			want:  2,
			want1: 2,
		},
		{
			name: "value is smaller than all",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 0,
			},
			want:  0,
			want1: 0,
		},
		{
			name: "value is larger than all",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 6,
			},
			want:  5,
			want1: 5,
		},
		{
			name: "empty slice",
			args: args{
				arr:    []int{},
				target: 3,
			},
			want:  0,
			want1: 0,
		},
		{
			name: "value exists, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 3,
			},
			want:  2,
			want1: 5,
		},
		{
			name: "value does not exist, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 6,
			},
			want:  7,
			want1: 7,
		},
		{
			name: "value is smaller than all, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 0,
			},
			want:  0,
			want1: 0,
		},
		{
			name: "value is larger than all, multiple same values",
			args: args{
				arr:    []int{1, 2, 3, 3, 3, 4, 5},
				target: 6,
			},
			want:  7,
			want1: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := EqualRange(tt.args.arr, tt.args.target)
			if got != tt.want {
				t.Errorf("EqualRange() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("EqualRange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
