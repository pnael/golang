package main

import "fmt"

func main() {
     sum :=0
     for i:=0; i <10; i++ {
     	 switch {
     	 	case( i % 3 == 0 ):
	   	      sum = sum  + i
		case( i % 5 ==0 ):
		      sum = sum + i 
		}
    }
    fmt.Println(sum)
}

