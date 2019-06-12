package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

//This method means type T implement the interfacce I
//But we don't need to explicitely declare that it does so
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
