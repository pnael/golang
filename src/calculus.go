package main

import "fmt"

func main() {
     var x,y int
     var reste float32

     fmt.Print("x=")
     _,err := fmt.Scanf("%d",&x)
     if err != nil {
     	fmt.Println(err)
	}
	
     fmt.Print("y=")
     _ , err = fmt.Scanf("%d",&y)
     if err != nil {
     	fmt.Println(err)
	}

     reste = float32(x % y)
     fmt.Println(reste)
    
     }