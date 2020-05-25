package diracMatrix

// import (
// 	"fmt"
// )

// // type Interface interface {
// // }

// type Matrix struct {
// 	Size  int
// 	Value [][]int
// 	Types struct {
// 		Dirac diracMatrix
// 	}
// }

// // Prints matrix in human readable format. (for debugging)
// func (m *Matrix) Print() {
// 	for _, row := range m.Value {
// 		fmt.Println(row)
// 	}
// }

// //////////////////////////////////
// // Dirac
// //////////////////////////////////
// /*
// # Sources:
// 	- https://mathworld.wolfram.com/DiracMatrices.html
// # Dirac Matrices (aka Gamma Matrices)
// 	- The Dirac matrices are a class of 4×4 matrices which arise in quantum electrodynamics. There are a variety of different symbols used, and Dirac matrices are also known as gamma matrices or Dirac gamma matrices.
// 	- The Dirac matrices alpha_n may be implemented in a future version of the Wolfram Language as DiracGammaMatrix[n], where n=1, 2, 3, 4, or 5.
// 	- The Dirac matrices are defined as the 4×4 matrices.
// */

// type diracMatrix [4][4]int

// // Returns a matrix.
// func NewDiracMatrix() *Matrix {
// 	scale := 4
// 	v := make([][]int, scale, scale)
// 	for i := 0; i < scale; i++ {
// 		v[i] = make([]int, scale, scale)
// 	}
// 	return &Matrix{
// 		Size:  scale,
// 		Value: v,
// 	}
// }
