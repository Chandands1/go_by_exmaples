package main

import ("fmt"
"slices")

func main(){
	var s []string
	fmt.Println(s, len(s), s == nil)

	s = make([]string, 4,5)

	fmt.Println(s, len(s), s == nil,cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2],s[3] = "c", "d"

	fmt.Println(s)
	s = append(s, "d")
	s = append(s, "e")

	fmt.Println(s)

	c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

	l := s[2:5]
    fmt.Println("sl1:", l)
	l = s[:5]
    fmt.Println("sl2:", l)

	l = s[2:]
    fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

	 t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }
}