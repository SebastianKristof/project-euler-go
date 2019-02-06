package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfMultiplesOf(t *testing.T) {
	tests := []struct {
		limit       int
		multiplesOf []int
		expected	int
	}{
		{
			limit:       10,
			multiplesOf: []int{3, 5},
			expected:	23,
		},
		{
			limit:       5,
			multiplesOf: []int{3, 5},
			expected: 3,
		},
		{
			limit: 6,
			multiplesOf: []int{2,3,5},
			expected: 14,
		},
		{
			limit: 6,
			multiplesOf: []int{5},
			expected: 5,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, sumOfMultiplesOf(test.limit, test.multiplesOf))
	}
}

func TestIsMultiple(t *testing.T) {
	tests := []struct {
		number int
		numbers []int
		expected	bool
	}{
		{
			number:       10,
			numbers: []int{3, 5},
			expected:	true,
		},
		{
			number:       7,
			numbers: []int{3, 5},
			expected: false,
		},
		// {
		// 	limit: 6,
		// 	multiplesOf: []int{2,3,5},
		// 	expected: 14,
		// },
		// {
		// 	limit: 6,
		// 	multiplesOf: []int{5},
		// 	expected: 5,
		// },
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isMultiple(test.number, test.numbers))
	}
}
