package main

import (
	"fmt"
	"math"
	"time"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println(true && false)

	const n = 500000000
	fmt.Println(math.Sin(n))

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good afternoon.")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d: ", twoD)

	var s []string
	s = append(s, "a", "b", "c")
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	twoDSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoDSlice)
	fmt.Println("2d[2][1]: ", twoDSlice[2][1])
}
