package vector

import (
	"fmt"
	"math"
)

// Vector representa um vetor 1D.
type Vector struct {
	Data []int64
}

// NewVector cria um novo Vector com os dados fornecidos.
func NewVector(data ...int64) *Vector {
	return &Vector{
		Data: data,
	}
}

// Print imprime os elementos do vetor.
func (v *Vector) Print() {
	fmt.Println("Vector:", v.Data)
}

//isSameLength checks if the vector has the same length as another vector.
//
//This function takes a vector to compare and return a bool
func (v *Vector) isSameLength(other *Vector) bool {
	return len(v.Data) == len(other.Data)
}

func round(number float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	rounded := math.Round(number*shift) / shift
	return rounded
}

// Add adds another vector to this vector.
func (v *Vector) Add(other *Vector) *Vector {
	if !(v.isSameLength(other)) {
		return nil
	}

	result := make([]int64, len(v.Data))
	for i := range v.Data {
		result[i] = v.Data[i] + other.Data[i]
	}
	return NewVector(result...)
}

func (v *Vector) Subtract(other *Vector) *Vector {
	if !(v.isSameLength(other)) {
		return nil
	}

	result := make([]int64, len(v.Data))
	for i := range v.Data {
		result[i] = v.Data[i] - other.Data[i]
	}

	return NewVector(result...)
}

func (v *Vector) MultiScalar(scalar int64) *Vector {
	result := make([]int64, len(v.Data))

	for i := range v.Data {
		result[i] = v.Data[i] * scalar
	}

	return NewVector(result...)
}

func (v *Vector) DotProduct(other *Vector) int64 {
	if !(v.isSameLength(other)) {
		return 0
	}

	var result int64

	for i := range v.Data {
		result = result + v.Data[i]*other.Data[i]
	}

	return result
}

func (v *Vector) Norm2() float64 {
	var result float64

	for i := range v.Data {
		result = result + (float64(v.Data[i]) * float64(v.Data[i]))
	}

	return round(math.Pow(result, 0.5), 2)
}

func (v *Vector) IsNormalized() bool {
	var norm float64 = 1
	return v.Norm2() == norm
}

func (v *Vector) AreOrthogonals(other *Vector) bool {
	return v.DotProduct(other) == 0
}

/* trocar o tipo do vector para float 64
func (v *Vector) ToNormalizer() *Vector {
	norm := v.Norm2()

	result := make([]int64, len(v.Data))

	for i := range v.Data {
		result[i] = v.Data[i] / norm
	}

	return NewVector(result...)
}*/
