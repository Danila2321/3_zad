package main

import (
	"fmt"
	"math"
)

func main() {
	b := 2.5
	b3 := b * b * b

	f := func(x float64) float64 {
		sum := b3 + x*x*x
		return (1 + math.Pow(math.Sin(sum), 2)) / math.Cbrt(sum)
	}

	fmt.Println("Задача А")
	fmt.Println("   x    |     y")
	fmt.Printf("%6.2f  | %8.6f\n", 1.28, f(1.28))
	fmt.Printf("%6.2f  | %8.6f\n", 1.68, f(1.68))
	fmt.Printf("%6.2f  | %8.6f\n", 2.08, f(2.08))
	fmt.Printf("%6.2f  | %8.6f\n", 2.48, f(2.48))
	fmt.Printf("%6.2f  | %8.6f\n", 2.88, f(2.88))
	fmt.Printf("%6.2f  | %8.6f\n", 3.28, f(3.28))

	fmt.Println("\nЗадача Б")
	fmt.Println("   x    |     y")
	fmt.Printf("%5.1f   | %8.6f\n", 1.1, f(1.1))
	fmt.Printf("%5.1f   | %8.6f\n", 2.4, f(2.4))
	fmt.Printf("%5.1f   | %8.6f\n", 3.6, f(3.6))
	fmt.Printf("%5.1f   | %8.6f\n", 1.7, f(1.7))
	fmt.Printf("%5.1f   | %8.6f\n", 3.9, f(3.9))
}
