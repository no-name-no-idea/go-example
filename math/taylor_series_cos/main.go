package main

import (
	"fmt"
	"math"
)

func main() {
	PI := 3.14159265358979

	repeat := 10000
	var f float64
	var x float64
	var term float64
	var result float64

	for n := 0; n <= 180; n += 5 {
		x = float64(n) * PI / 180
		f = math.Cos(x)
		term = 1
		result = 1
		for i := 1; i <= repeat; i = i + 1 {
			term *= x * x / float64((2*i)*(2*i-1))
			if i%2 == 1 {
				result -= term
			} else {
				result += term
			}
		}
		fmt.Printf("x=%3d : cos(x)=%.18f , 근사값 = %.18f\n", n, f, result)
	}
}
