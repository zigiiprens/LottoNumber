package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
	"testing"
)

func within(a, b, maxDiff float64) bool {
	absDiff := math.Abs(a - b)
	return absDiff < maxDiff
}

func TestLottonumber(t *testing.T) {

	s := distuv.Uniform{Min: 0, Max: 49}
	n := 100
	x := lotonumber(n)
	total := 0
	total1 := 0
	total2 := 0
	total3 := 0
	total4 := 0
	c := 0
	data := make([]float64, 1e5)

	for _, v := range x {
		for i, w := range v {
			total++
			total1++
			total2++
			total3++
			total4++
			data[i] = float64(w)
			if w > 48 {

				t.Error("Error with number better than 48")
				total1--
			}
			if w < 0 {
				t.Error("Error with number less than 0")
				total2--
			}
			if len(v) != 7 {
				t.Error("Error with lotto number lenght")
				total3--
			}
			if within(float64(c), float64(w), 0.02) == true {
				t.Error("Error with maxDifference")
				total4--
			}
			c = w
		}
	}
	_, stdd := stat.MeanStdDev(data, nil)
	sknn := stat.Skew(data, nil)
	kurtt := stat.ExKurtosis(data, nil)
	fmt.Println("StdDev, Skewness, ExKurt calculated stat/", stdd, sknn, kurtt)

	std := s.StdDev()
	skn := s.Skewness()
	kurt := s.ExKurtosis()
	fmt.Println("StdDev, Skewness, ExKurt from disturv/uniform", std, skn, kurt)
	fmt.Println(total, total1, total2, total2, total3, total4)
}
