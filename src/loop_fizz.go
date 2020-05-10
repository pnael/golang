package main

import "fmt"

func main() {
     for i:=0; i <1000; i++ {
     	 switch {
     	 	case( i % 3 == 0 ):
	    	      fmt.Println("Fuzz")
		case( i % 5 ==0 ):
		      fmt.Println("FizzBuzz")
		default:
		      fmt.Println(i)
		}
    }
}

