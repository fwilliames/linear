package matrix

import (
    "fmt"
    "linear/vector"
)

// Matrix represents a 2D matrix.
type Matrix struct {
    Columns []*vector.Vector // Columns represented as vectors
}

// NewMatrix creates a new Matrix with the given columns.
func NewMatrix(columns ...*vector.Vector) *Matrix {
    return &Matrix{
        Columns: columns,
    }
}

// Print prints the matrix.
func (m *Matrix) Print() {
    numRows := len(m.Columns[0].Data)
    numCols := len(m.Columns)

    for i := 0; i < numRows; i++ {
		fmt.Print("|")
        for j := 0; j < numCols; j++ {
            fmt.Printf("%d ", m.Columns[j].Data[i])
        }
		fmt.Print("|")
        fmt.Println() // Move to the next line after printing each row
    }
}

