package main

import (
	"log"
)

func main() {
	// Setup Logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("QCE: ")

	// m1 := matrix.NewDiracMatrix()
	// m1.Print()
}
