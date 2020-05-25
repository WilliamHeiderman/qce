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

// Returns sum of 2 matrices.
// Contraints: Matrices must be the same size.
func Add(m1 [][]int, m2 [][]int) ([][]int, error) {
	log.Println(fmt.Sprintf("Add: %v + %v", m1, m2))
	if !CompareSize(m1, m2) {
		return nil, errors.New("matrices must be the same size")
	}

	m3 := [][]int{}
	for i := 0; i < len(m1); i++ {
		var temp []int
		for j := 0; j < len(m1[i]); j++ {
			temp = append(temp, m1[i][j]+m2[i][j])
		}
		m3 = append(m3, temp)
	}
	return m3, nil
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
