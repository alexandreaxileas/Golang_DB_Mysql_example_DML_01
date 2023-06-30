package dbfullexercise

import (
	"math"
	"math/rand"
)

// fnRetGrade - Grade from 0 to 10
func fnRetGrade() float64 {
	rand.Float64()
	vGrade := math.Ceil((rand.Float64()*10)*10) / 10
	return vGrade
}
