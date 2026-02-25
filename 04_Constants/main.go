package main

import ("fmt"
"math")

const s string = "constant"

func main(){
	fmt.Println(s)
	const n = 5000
	const d = 3e7 / n
	fmt.Println(d)

	fmt.Println(math.Sin(d))

}