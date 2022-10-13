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
    var k *int // K is a point to an int
    k = &i
    fmt.Println("K is",k)
    fmt.Println("K point to value",*k)
	p := &i // point to i
	fmt.Println(p)
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

    l := new(int) // l is now a pointer to an int
    *l = 4

    fmt.Println("l is",l)
    fmt.Println("l point to a variable with the value",*l)
}
