package main

import "fmt"

func main(){
	//创建切片
	s := make([]string, 3)
	fmt.Println("emp:", s)
	
	//初始化切片
	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[1:])

	//切片追加元素
	s = append(s, "d", "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c,s)
	fmt.Println("cpy:",c)

	//内部切片
	l := s[2:5]
    fmt.Println("sl1:", l)
    l = s[:5]
    fmt.Println("sl2:", l)
    l = s[2:]
    fmt.Println("sl3:", l)

	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}