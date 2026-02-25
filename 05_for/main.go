package main

import "fmt"

func main(){
	x := 5
	for i := 0; i <= x ; i++{
		fmt.Println(i)
	}
	j := 3
	for j < 5{
		fmt.Println(j)
		j = j+1
	}

	for i := range 3{
		fmt.Println("range", i)
	}
	for{
		fmt.Println("loop")
		break
	}
	for n := range 6{
		if n %2 ==0{
			continue
		}
		fmt.Println(n)
	}
}