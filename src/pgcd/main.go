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

var debug bool = false

func main() {
	var f int
	if len(os.Args) == 4 {
		if os.Args[3] == "-d" {
			debug = true
		}
	}
	grosEntierS := os.Args[1]
	grosEntierS2 := os.Args[2]

	grosEntier, err := strconv.Atoi(grosEntierS)
	grosEntier2, err2 := strconv.Atoi(grosEntierS2)
	if err == nil && err2 == nil {
		f = getPGDC(grosEntier, grosEntier2, debug)
	}

	if f != 1 {
		a := grosEntier / f
		b := grosEntier2 / f
		fmt.Println(f, a, b)
		f = getPGDC(a, b, debug)
		fmt.Println(f, a/f, b/f)
	} else {
		fmt.Println(f, grosEntier/f, grosEntier2/f)
	}
}

func getPGDC(a int, b int, d bool) int {
	var facteurs []int
	facteurs = make([]int, 0)
	var facteurs2 []int
	facteurs2 = make([]int, 0)

	grosEntier := a
	grosEntier2 := b

	for i := 1; i <= grosEntier; i++ {
		if math.Mod(float64(grosEntier), float64(i)) == 0 {
			facteurs = append(facteurs, i)
		}
	}

	if d {
		fmt.Println(facteurs)
	}

	for i := 1; i <= grosEntier2; i++ {
		if math.Mod(float64(grosEntier2), float64(i)) == 0 {
			facteurs2 = append(facteurs2, i)
		}
	}
	if d {
		fmt.Println(facteurs2)
	}

	f := pgdc(facteurs, facteurs2)
	return f
}
