package main

import (
	"fmt"
	"time"
)

func main(){
	day := 6
	switch day{
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
		case 3:
		fmt.Println("Wednesday")
		case 4:
		fmt.Println("Thursday")
		case 5:
		fmt.Println("Friday")
		default:
			fmt.Println("Weekend")
	}

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("Weekends")
	default:
		fmt.Println("Weekdays")
	}
}