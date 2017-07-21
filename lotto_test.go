package main

import (
	"fmt"
	"testing"
)

func TestLottonumber(t *testing.T) {

	x := lotonumber(143000)
	c := 0
	for _, v := range x {
		for _, w := range v {
			c++
			if (w > 48) || (w < 0) || (len(v) != 7) {
				t.Error("Error")
			}
		}
	}
	fmt.Println(c)
}
