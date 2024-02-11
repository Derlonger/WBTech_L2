package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		words    []string
		expected map[string][]string
	}{
		{
			name:  "Example",
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
		},
		{
			name:     "No Anagrams",
			words:    []string{"apple", "banana", "cherry", "orange"},
			expected: map[string][]string{},
		},
		{
			name:  "Mixed Case",
			words: []string{"Пятак", "Пятка", "Тяпка"},
			expected: map[string][]string{
				"акптя": {"пятак", "пятка", "тяпка"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findAnagrams(&tc.words)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
