package main

import("fmt")

func main(){
	var a [5]int
	fmt.Println("array", a)
	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[4])

	fmt.Println("Length", len(a))

	 b := [5]int{1,2,3,4,5}
	 fmt.Println("dcl",b)

	 c := [...]int{0,1,2,3,4,5,6}
	 fmt.Println("dcl",c)
	 b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)
}