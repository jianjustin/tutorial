package main

import "fmt"

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	a, b := vals()
	fmt.Printf("a = %d, b = %d\n", a, b)

	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("total = %d\n", total)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

func plus(a int, b int) int {
	return a + b
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
