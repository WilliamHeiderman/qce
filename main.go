package main

import (
	"fmt"
	"log"

	"github.com/WilliamHeiderman/qce/pkg/matrix"
)

func main() {
	// Setup Logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("QCE: ")

	// m1 := matrix.NewDiracMatrix()
	// m1.Print()

	// product, err := matrix.Multiply([][]int{{1, 1}, {1, 1}}, [][]int{{2, 2}, {2, 2}})
	product, err := matrix.Multiply([][]int{{1, 2}, {3, 4}}, [][]int{{5, 6}, {7, 8}})
	fmt.Println(product, err)
}
