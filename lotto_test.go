package main

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
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

	if !within(stdd, std, maxDiff) {
		t.Error("Error with standart deviations difference", stdd)
	}
	if !within(sknn, skn, maxDiff) {
		t.Error("Error with Skewnesses difference", sknn)
	}
	if !within(kurtt, kurt, maxDiff) {
		t.Error("Error with ExKurtosises difference", kurtt)
	}
}

func BenchmarkLottonumber(b *testing.B){
	for n:=0; n < b.N; n++{
		lottonumber(100000)
	}
}
