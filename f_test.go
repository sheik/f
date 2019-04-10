package main

import (
	"testing"
)

func TestExpandRange(t *testing.T) {
	SuccessTable := []struct {
		input  []string
		output []int
	}{
		{[]string{"1", "2", "3"}, []int{1, 2, 3}},
		{[]string{"4", "5"}, []int{4, 5}},
		{[]string{"4-6"}, []int{4, 5, 6}},
		{[]string{"4-6"}, []int{4, 5, 6}},
		{[]string{"1-10"}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]string{"1", "7-10"}, []int{1, 7, 8, 9, 10}},
		{[]string{"1-3"}, []int{1, 2, 3}},
	}

	for _, entry := range SuccessTable {
		result, err := ExpandRange(entry.input)
		if err != nil {
			t.Error(err)
		}
		if len(result) != len(entry.output) {
			t.Errorf("Output list lengths do not match: %v / %v", result, entry.output)
		}
		for i, n := range entry.output {
			if result[i] != n {
				t.Errorf("Expected: %v / Got: %v", entry.output, result)
			}
		}
	}

	FailTable := []struct {
		input []string
	}{
		{[]string{"a"}},
		{[]string{"b"}},
		{[]string{"10-7"}},
	}

	for i, entry := range FailTable {
		_, err := ExpandRange(entry.input)
		if err == nil {
			t.Errorf("Expected a non-nil error value for FailTable entry %d", i+1)
		}
	}
}
