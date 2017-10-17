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
	maxDiff := 0.2
	n := 100000
	x := lottonumber(n)
	total := 0
	total1 := 0
	total2 := 0
	total3 := 0
	var data []float64

	for _, v := range x {
		for _, w := range v {
			data = append(data, float64(w))
			total++
			total1++
			total2++
			total3++
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
		}
	}
	_, stdd := stat.MeanStdDev(data, nil)
	sknn := stat.Skew(data, nil)
	kurtt := stat.ExKurtosis(data, nil)
	std := s.StdDev()
	skn := s.Skewness()
	kurt := s.ExKurtosis()

	fmt.Println("StdDev, Skewness, ExKurt calculated stat/", stdd, sknn, kurtt)
	fmt.Println("StdDev, Skewness, ExKurt from disturv/uniform", std, skn, kurt)
	fmt.Println(total, total1, total2, total2, total3)
	fmt.Println("The difference between the standart deviations are lower than",maxDiff, within(stdd,std,maxDiff))
	fmt.Println("The difference between the Skewnesses are lower than",maxDiff, within(sknn,skn,maxDiff))
	fmt.Println("The difference between the ExKurtosises are lower than",maxDiff, within(kurtt,kurt,maxDiff))
}
