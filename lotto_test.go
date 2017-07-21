package main

import (
	"fmt"
	"math"
	"testing"
	//"github.com/gonum/stat"
	//"github.com/gonum/stat/distuv"
)

func TestLottonumber(t *testing.T) {
	n := 100000
	x := lotonumber(n)
	var sum, kare, ort int
	var sapma float64
	c := 0
	for _, v := range x {
		for _, w := range v {
			sum = sum + w
			kare = kare + (w ^ 2)

			c++
			if (w > 48) || (w < 0) || (len(v) != 7) {
				t.Error("Error")
			}

		}

	}
	ort = ((sum / (n * 7)) ^ 2) * (n * 7)
	sapma = math.Sqrt(math.Abs(float64(((kare - ort) / (n * 7)))))
	fmt.Println("Standart Devition is:", sapma)
	fmt.Println("Numbers of random numbers:", c)
}
