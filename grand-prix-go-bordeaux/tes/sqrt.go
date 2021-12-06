package piscine

import (
	"fmt"
	"math"
)

// Sqrt oui
func Sqrt(nb int) int {

	var x float64
	x = math.Sqrt(4)
	fmt.Println(x)
}
func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
