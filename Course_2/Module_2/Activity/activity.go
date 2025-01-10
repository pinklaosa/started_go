package main

import "fmt"

func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return ((a * (t * t)) / 2) + (vo * t) + so
	}
}

func main() {

	var a float64
	var vo float64
	var so float64
	var time float64

	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("vo: ")
	fmt.Scan(&vo)
	fmt.Print("so: ")
	fmt.Scan(&so)
	fmt.Print("time (secound) > ")
	fmt.Scan(&time)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Println(fn(time))

}

// s = ½ a t2 + vot + so
