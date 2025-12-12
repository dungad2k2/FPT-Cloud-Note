package main

import "testing"

func TestAdd(t *testing.T){
	tests := [] struct {
		name string
		inputX int
		inputY int
		expected int
	}{
		{"Positive numbers", 2, 3, 5},
		{"Negative numbers", -1, -2, -3},
		{"Mixed numbers", -5, 5, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.inputX, tc.inputY)
			if result != tc.expected{
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}