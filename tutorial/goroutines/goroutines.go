package goroutines

import "fmt"

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, ":", s)
	}
}
