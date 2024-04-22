package slice_test

import (
	"newsapp/pkg/slice"
	"testing"
)

type TestCaseMapFromUintToUint64 struct {
	input    []uint
	expected []uint64
}

type TestCaseMapFromUint64ToUint struct {
	input    []uint64
	expected []uint
}

func TestMapFromUintToUint64(t *testing.T) {
	testCase := TestCaseMapFromUintToUint64{
		input:    []uint{1, 2, 3, 4, 5},
		expected: []uint64{1, 2, 3, 4, 5},
	}

	expected := slice.MapFromUintToUint64(testCase.input)

	for _, u := range testCase.expected {
		var exist bool
		for _, u2 := range expected {
			if u == u2 {
				exist = true
			}
		}

		if !exist {
			t.Errorf("expected %v, got %v", testCase.expected, expected)
		}
	}

}
func TestMapFromUint64ToUint(t *testing.T) {
	testCase := TestCaseMapFromUint64ToUint{
		input:    []uint64{1, 2, 3, 4, 5},
		expected: []uint{1, 2, 3, 4, 5},
	}

	expected := slice.MapFromUint64ToUint(testCase.input)

	for _, u := range testCase.expected {
		var exist bool
		for _, u2 := range expected {
			if u == u2 {
				exist = true
			}
		}

		if !exist {
			t.Errorf("expected %v, got %v", testCase.expected, expected)
		}
	}

}
