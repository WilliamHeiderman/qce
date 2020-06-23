package matrix

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSuccess(t *testing.T) {
	// Predefined
	var y int = 0
	var x int = 0
	var rows []int = nil
	expected := []string{
		fmt.Sprintf("- It returns an empty slice of slices."),
		fmt.Sprintf("- It should not return an error."),
		fmt.Sprintf("- Every column should have a capacity equal to the y arg: %v.", y),
		fmt.Sprintf("- Every column should have a length equal to the y arg: %v.", y),
		fmt.Sprintf("- Every row should have a capacity equal to the x arg: %v.", x),
		fmt.Sprintf("- Every row has a length equal to the x arg: %v.", x),
		"- Every row should be an empty slice.",
	}

	// Log
	t.Log(
		"\nPredefined: No rows are passed in as args.",
		fmt.Sprintf("\n- y arg is %v.", y),
		fmt.Sprintf("\n- x arg is %v.", x),
		fmt.Sprintf("\n- rows arg is %v.", rows),
		"\nExpected Result:",
		"\n"+strings.Join(expected, "\n"),
	)

	// Run Test
	m1, err := Create(x, y)

	// Assertions
	assert.IsType(t, [][]int{}, m1, expected[0])
	assert.Equal(t, [][]int{}, m1, expected[0])
	assert.NoError(t, err, expected[1])
	assert.Equal(t, 0, cap(m1), expected[2])
	assert.Equal(t, 0, len(m1), expected[3])
	for _, row := range m1 {
		assert.IsType(t, []int{}, row, "Every row should be a slice of integers.")
		assert.Equal(t, []int{}, row, "Every row should be an empty slice of integers.")
		assert.Equal(t, 0, cap(row), expected[4])
		assert.Equal(t, 0, len(row), expected[5])
	}

	// Run Test
	x = 2
	y = 2
	m1, err = Create(x, y)

	// Assertions
	assert.IsType(t, [][]int{}, m1, expected[0])
	assert.Equal(t, [][]int{{0, 0}, {0, 0}}, m1, expected[0])
	assert.NoError(t, err, expected[1])
	assert.Equal(t, x, cap(m1), expected[2])
	assert.Equal(t, x, len(m1), expected[3])
	for _, row := range m1 {
		assert.IsType(t, []int{0, 0}, row, "Every row should be a slice of integers.")
		assert.Equal(t, []int{0, 0}, row, "Every row should be an empty slice of integers.")
		assert.Equal(t, x, cap(row), expected[4])
		assert.Equal(t, x, len(row), expected[5])
	}

	m1 = append(m1, []int{0, 0})
}

func TestCreateError(t *testing.T) {
	t.Log("Number of Rows cannot be greater than maxRows.")
}

///////////////////////////////////////
// Arithmetic
//////////////////////////////////////

// Add()
func TestAddSuccess(t *testing.T) {

	m1 := [][]int{{1, 1}, {1, 1}}
	m2 := [][]int{{2, 2}, {2, 2}}
	m3_expected := [][]int{{3, 3}, {3, 3}}

	t.Log(
		"\nPredefined:",
		fmt.Sprintf("\n- m1 is %v.", m1),
		fmt.Sprintf("\n- m2 is %v.", m2),
		"\nExpected Result:",
		fmt.Sprintf("\n- m3 is %v.", m3_expected),
	)

	m3, err := Add(m1, m2)

	assert.NoError(t, err, "It should not return an error.")
	assert.IsType(t, [][]int{}, m3, "It should return a matrix.")
	assert.Equal(t, m3_expected, m3, "It should return the sum of 2 matrices.")
}

func TestAddErrorInequivalentSize(t *testing.T) {
	m1 := [][]int{{1, 1}, {1, 1}}
	m2 := [][]int{{2, 2}, {2, 0}, {1, 1}}

	t.Log(
		"\nPredefined:",
		fmt.Sprintf("\n- m1 is %v.", m1),
		fmt.Sprintf("\n- m2 is %v.", m2),
		"\nExpected Result:",
		fmt.Sprintf("\n- An error is returned."),
	)

	_, err := Add(m1, m2)

	assert.Error(t, err, "It should return an error.")
}

// Subtract()
func TestSubtractSuccess(t *testing.T) {
	minuend := [][]int{{5, 5}, {5, 5}}
	subtrahend := [][]int{{1, 1}, {1, 1}}

	difference, err := Subtract(minuend, subtrahend)
	assert.NoError(t, err, "Subtract() should not return an error")

	expectedM3 := [][]int{{4, 4}, {4, 4}}
	assert.Equal(t, expectedM3, difference, fmt.Sprintf("Subtract() should output the difference between the minuend (m1 %v) and the subtrahend (m2 %v). Said output should be (m3 %v).", minuend, subtrahend, difference))
}

func TestSubtractErrorInequivalentSize(t *testing.T) {
	minuend := [][]int{{5, 5}, {5, 5}}
	subtrahend := [][]int{{1, 1}}

	difference, err := Subtract(minuend, subtrahend)
	assert.Error(t, err, "Subtract() should return an error.")
	assert.Nil(t, difference, "Subtract() should return a difference of nil.")
}

// Multiply()
func TestMultiplySuccess(t *testing.T) {
	t.Log("Multiply() should return a matrix equal to the product of the multiplicand and the multiplier.")

	var testCases = []struct {
		multiplicand    [][]int
		multiplier      [][]int
		expectedProduct [][]int
	}{
		{
			[][]int{{1, 1}, {1, 1}},
			[][]int{{2, 2}, {2, 2}},
			[][]int{{4, 4}, {4, 4}},
		},
		{
			[][]int{{1, 2}, {3, 4}},
			[][]int{{5, 6}, {7, 8}},
			[][]int{{19, 22}, {43, 50}},
		},
	}

	for _, tc := range testCases {
		product, err := Multiply(tc.multiplicand, tc.multiplier)
		assert.NoError(t, err, "Multiply() should not return an error.")
		assert.Equal(t, tc.expectedProduct, product, fmt.Sprintf("Multiply() should return a matrix equal to the product of the multiplicand (%v) and the multiplier (%v). Said product should be (%v).", tc.multiplicand, tc.multiplier, tc.expectedProduct))
	}

}

///////////////////////////////////////
// Metadata
//////////////////////////////////////

// Size()
func TestSizeSuccess(t *testing.T) {
	m1 := [][]int{{0, 0}, {0, 0}}
	x, y := Size(m1)
	assert.Equal(t, len(m1[0]), x, "It should return an x value that is equal to number of columns in the matrix.")
	assert.Equal(t, len(m1), y, "It should return an y value that is equal to number of rows in the matrix.")
}

// CompareSize()
func TestCompareSizeTrue(t *testing.T) {
	m1 := [][]int{{0, 0}, {0, 0}}
	ok := CompareSize(m1, m1)
	assert.True(t, ok, "It should return true.")
}

func TestCompareSizeFalseX(t *testing.T) {
	m1 := [][]int{{0, 0}, {0, 0}}
	m2 := [][]int{{0}, {0}}
	ok := CompareSize(m1, m2)
	assert.False(t, ok, "It should return false.")
}

func TestCompareSizeFalseY(t *testing.T) {
	m1 := [][]int{{0, 0}, {0, 0}}
	m2 := [][]int{{0, 0}}
	ok := CompareSize(m1, m2)
	assert.False(t, ok, "It should return false.")
}

// IsValid()
func TestIsValidTrue(t *testing.T) {
	m1 := [][]int{{0, 0}, {0, 0}}
	ok := IsValid(m1)
	assert.True(t, ok, "It should return true.")
}

func TestIsValidFalse(t *testing.T) {
	m1 := [][]int{{0, 0}, {0}}
	ok := IsValid(m1)
	assert.False(t, ok, "It should return false.")
}

// IsSquare()
func TestIsSquareTrue(t *testing.T) {
	m1 := [][]int{{0, 0}, {0, 0}}
	ok := IsSquare(m1)
	assert.True(t, ok, "It should return true.")
}

func TestIsSquareFalse(t *testing.T) {
	m1 := [][]int{{0}, {0}}
	ok := IsSquare(m1)
	assert.False(t, ok, "It should return false.")

	m1 = [][]int{{0, 0}}
	ok = IsSquare(m1)
	assert.False(t, ok, "It should return false.")
}
