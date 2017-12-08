package main

import "fmt"
import (
	"math/rand"
	"math"
)

func main() {
	fmt.Println("Hello World.")
	num := 0
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i, "----", rand.Intn(10))
		num += i
		if num % 3 == 0 {
			fmt.Println("======================")
		} else if num % 3 == 1 {
			fmt.Println("**********************")
		} else {
			fmt.Println("######################")
		}
	}
	fmt.Println(math.Pi)

	v := Person{"刘枫", "xxx"}
	v.Name = "张三"

	fmt.Println(v.Name)

	s := &v

	fmt.Println(s)

	a := make([]string, 10)
	fmt.Println(a)

	var pow = []string{"刘枫", "张飒", "李斯", "王屋"}

	for v := range pow {
		fmt.Printf("s = %s\n", pow[v])
	}



}

type Person struct {
	Name   string
	Avatar string
}