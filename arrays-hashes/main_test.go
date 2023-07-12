package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// add test cases here, especially edge or corner cases
		{"empty slice", args{[]int{}}, false},
		{"one element", args{[]int{1}}, false},
		{"two elements", args{[]int{1, 2}}, false},
		{"two elements, duplicate", args{[]int{1, 1}}, true},
		{"three elements, duplicate", args{[]int{1, 2, 1}}, true},
		{"three elements, no duplicate", args{[]int{1, 2, 3}}, false},
		{"four elements, duplicate", args{[]int{1, 2, 3, 1}}, true},
		{"four elements, no duplicate", args{[]int{1, 2, 3, 4}}, false},
		{"five elements, duplicate", args{[]int{1, 2, 3, 4, 1}}, true},
		{"five elements, no duplicate", args{[]int{1, 2, 3, 4, 5}}, false},
		{"ten elements, duplicate", args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, true},
		{"four elements, no duplicate", args{[]int{1, 2, 3, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("ContainsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "gòlang",
				t: "langgò",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
