package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{"All is true", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		}, false},
		{"All is false", []int{1, 1, 1}, func(i int) bool {
			if i == 1 {
				return true
			}
			return false
		}, true},
	}

	for _, tt := range tests {
		if got := All(tt.slice, tt.predicate); got != tt.want {
			t.Errorf("Name: %v All() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{"Any is true", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		}, true},
		{"Any is false", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i == 6 {
				return true
			}
			return false
		}, false},
	}

	for _, tt := range tests {
		if got := Any(tt.slice, tt.predicate); got != tt.want {
			t.Errorf("Name: %v Any() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNone(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{"None is false", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		}, false},
		{"None is true", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i == 6 {
				return true
			}
			return false
		}, true},
	}

	for _, tt := range tests {
		if got := None(tt.slice, tt.predicate); got != tt.want {
			t.Errorf("Name: %v None() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      int
	}{
		{"Index found", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		}, 2},
		{"Index not found", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i == 6 {
				return true
			}
			return false
		}, -1},
	}

	for _, tt := range tests {
		if got := Index(tt.slice, tt.predicate); got != tt.want {
			t.Errorf("Name: %v Index() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      int
	}{
		{"Find found result", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		}, 3},
	}

	for _, tt := range tests {
		if got := Find(tt.slice, tt.predicate); got != tt.want {
			t.Errorf("Name: %v Find() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// this example must panic!
func ExampleFind() {
	//sl := []int{1,2,3,4,5}
	//Find(sl, func(elem int) bool {
	//	if elem == 99 {
	//		return true
	//	}
	//	return false
	//})
}