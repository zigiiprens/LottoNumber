package main

import (
	"fmt"
	"testing"
        "math"
        "gonum.org/v1/gonum/stat/distuv"
)



func within(a, b, maxDiff float64) bool {
	absDiff := math.Abs(a - b)
	return absDiff < maxDiff
}

func TestLottonumber(t *testing.T) {
	n := 10
	x := lotonumber(n)
	total := 0
	c := 0
	for _, v := range x {
		for _, w := range v { 
			total++                   
			if (w > 48) {t.Error("Error with number better than 48")}
			if (w < 0) {t.Error("Error with number less than 0")}
			if (len(v)!=7) {t.Error("Error with lotto number lenght")}
                        if (within(float64(c),float64(w),0.02) == true) {t.Error("Error with maxDifference")}
                        c = w
		}		
	}
	s := distuv.Uniform{Min : 0, Max : 49}
        std := s.StdDev()
        skn := s.Skewness()
        kurt := s.ExKurtosis()
	fmt.Println(std, skn, kurt)
	fmt.Println("Total Numbers generated is %d",total)
}
