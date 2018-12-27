package main

import "fmt"

func addUn(xVal *int) {
	*xVal = *xVal + 1
}

func swapInt(i, j *int) {
	*i, *j = *j, *i
}

func main() {
	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	addUn(&i)
	fmt.Println(i)

	fmt.Println(i, j)
	swapInt(&i, &j)
	fmt.Println(i, j)
}
