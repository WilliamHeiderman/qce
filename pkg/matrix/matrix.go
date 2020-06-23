package matrix

import (
	"errors"
	"fmt"
	"log"
)

///////////////////////////////////////
// Arithmetic
//////////////////////////////////////

// Returns 1 matrix. Total number of rows is equal to y arg, total number of columns is equal to x arg, predefined rows are defined by rows arg.
func Create(x int, y int, rows ...[]int) ([][]int, error) {
	log.Println(fmt.Sprintf("Creating a %vx%v matrix with predefined rows of %v", x, y, rows))
	m1 := make([][]int, y, y)

	i := 0
	for i < y {
		log.Println("i ::: ", i)
		row := make([]int, x, x)
		if i < len(rows) {
			m1[i] = rows[i]
		} else {
			m1[i] = row
		}
		i++
	}

	return m1, nil
}

// Add() returns a matrix equal to the sum of the augend and the addend.
func Add(augend [][]int, addend [][]int) ([][]int, error) {
	log.Println(fmt.Sprintf("Add: %v + %v", augend, addend))
	if !CompareSize(augend, addend) {
		return nil, errors.New("Both matrices must be the same size.")
	}

	m3 := [][]int{}
	for i := 0; i < len(augend); i++ {
		var temp []int
		for j := 0; j < len(augend[i]); j++ {
			temp = append(temp, augend[i][j]+addend[i][j])
		}
		m3 = append(m3, temp)
	}
	return m3, nil
}

// Subtract() returns a matrix equal to the difference of the minuend and the subtrahend.
func Subtract(minuend [][]int, subtrahend [][]int) (sum [][]int, err error) {
	log.Println(fmt.Sprintf("Subtract: %v - %v", minuend, subtrahend))
	if !CompareSize(minuend, subtrahend) {
		return nil, errors.New("Both matrices must be the same size.")
	}

	for i := 0; i < len(minuend); i++ {
		sum = append(sum, []int{})
		for j := 0; j < len(minuend[i]); j++ {
			sum[i] = append(sum[i], minuend[i][j]-subtrahend[i][j])
		}
	}
	return
}

// Multiply returns a matrix equal to the product of the multiplicand and the multiplier.
func Multiply(multiplicand [][]int, multiplier [][]int) (product [][]int, err error) {
	log.Printf("\nMultiply: %v x %v\n", multiplicand, multiplier)

	columns, _ := Size(multiplicand)
	_, rows := Size(multiplier)
	if columns != rows {
		return nil, errors.New("The number of columns in the multiplicand must equal the number of rows in the multiplier.")
	}

	if product, err = Create(rows, columns); err != nil {
		return nil, err
	}

	for i := 0; i < len(multiplicand); i++ {
		for j := 0; j < len(multiplicand[i]); j++ {
			fmt.Printf("\nc[%v][%v] = ", i, j)
			for k := 0; k < len(multiplicand[i]); k++ {
				product[i][j] += multiplicand[i][k] * multiplier[k][j]
				// fmt.Printf("a[%v][%v] x b[%v][%v]%v", i, k, k, j, func(k int, l int) string {
				// 	if k < l-1 {
				// 		return " + "
				// 	}
				// 	return ""
				// }(k, len(multiplicand[i])))
			}
		}
	}
	return
}

///////////////////////////////////////
// Metadata
//////////////////////////////////////

// Returns size of matrix as x, y
func Size(m [][]int) (x int, y int) {
	return len(m[0]), len(m)
}

// Returns true if m1 and m2 are the same size.
func CompareSize(m1 [][]int, m2 [][]int) bool {
	m1SizeX, m1SizeY := Size(m1)
	m2SizeX, m2SizeY := Size(m2)
	if m1SizeX == m2SizeX && m1SizeY == m2SizeY {
		return true
	}
	return false
}

// Returns true when all the rows have the same length.
func IsValid(m [][]int) bool {
	var xLength int = len(m[0])
	for _, x := range m {
		if len(x) != xLength {
			return false
		}
	}
	return true
}

// Returns true when n columns equals n rows.
func IsSquare(m [][]int) bool {
	var yLength int = len(m)
	for _, x := range m {
		if len(x) != yLength {
			return false
		}
	}
	return true
}
