package main

import("fmt")

func plus(a int, b int)int{
	return a+b
}

func plusplus(a,b,c int)int{
	return a+b+c
}

func main(){
	res := plus(2,3)
	fmt.Println(res)

	res = plusplus(2,3,4)
	fmt.Println(res)

}