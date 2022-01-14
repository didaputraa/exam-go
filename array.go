package main

import "fmt"

func main() {

	var x[] string
	var y[5] string

	x = append(x, "Hello")
	x = append(x, "World")

	y[0] = "h"
	y[1] = "e"
	y[2] = "l"
	y[3] = "l"
	y[4] = "o"
	
	map_x := make(map[int]string)
	map_x[0] = "maps 1"
	map_x[1] = "maps 2"
	
	fmt.Println("nilai var x:", len(x))
	fmt.Println("nilai var y:", len(y))
	
	for i:= 0; i<len(x); i++ {
		fmt.Print(x[i], " ")
	}
	
	fmt.Println()
	
	for i:= 0; i<len(y); i++ {
		fmt.Print(y[i]);
	}
	
	delete(map_x, 0)
	
	x_value,x_index := map_x[1]
	
	fmt.Println()
	fmt.Println(len(map_x), map_x[1])
	fmt.Println(x_index, x_value)
	
	z := map[int]map[int]string {
		0 : map[int]string {
			  0: "I",
			  1: "Love",
			  2: "You",
		  },
		1 : map[int]string{},
	}
	z2 := []string{"a", "b", "c", "d",}
	
	fmt.Println(z[0][0], z[0][1], z[0][2])
	fmt.Println(z[0], z[1])
	
	for i:=0; i<len(z2); i++ {
		fmt.Print(z2[i])
	}
}