package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Return true si i est dans tab
func isin(i int, tab []int) bool {
	for j := len(tab) - 1; j > 0; j = j - 1 {
		if tab[j] == i {
			return true
		}
	}
	return false
}

func pgdc(tab1 []int, tab2 []int) int {
	for i := len(tab1) - 1; i > 0; i-- {
		if isin(tab1[i], tab2) {
			return tab1[i]
		}
	}
	return 1
}

func main() {

	var facteurs []int
	facteurs = make([]int, 0)
	var facteurs2 []int
	facteurs2 = make([]int, 0)

	grosEntierS := os.Args[1]
	grosEntierS2 := os.Args[2]

	grosEntier, err := strconv.Atoi(grosEntierS)
	grosEntier2, err2 := strconv.Atoi(grosEntierS2)

	if err == nil {
		for i := 1; i < grosEntier; i++ {
			if math.Mod(float64(grosEntier), float64(i)) == 0 {
				facteurs = append(facteurs, i)

			}
		}
	}

	if err2 == nil {
		for i := 1; i < grosEntier2; i++ {
			if math.Mod(float64(grosEntier2), float64(i)) == 0 {
				facteurs2 = append(facteurs2, i)
			}
		}
	}
	f := pgdc(facteurs, facteurs2)

	if f != 1 {
		fmt.Println("Pas fini !!")
	}

	fmt.Println(f, grosEntier/f, grosEntier2/f)
}
