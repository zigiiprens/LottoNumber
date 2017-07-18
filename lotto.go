package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lottonumber(n int) (lot [][]int) {
	for j := 0; j < n; j++ {
		const min int = 0
		const max int = 49
		var z int
		var lottonumber []int

		for i := 0; i < 7; i++ {
			z = random(min, max)
			lottonumber = append(lottonumber, z)

			//fmt.Println(z)-
		}
		lot = append(lot, lottonumber)

	}
	return
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	var s int
	fmt.Println("Please Enter the n value")
	fmt.Scanf("%d", &s)
	fmt.Println(s)

	fmt.Println(lottonumber(s))

}
