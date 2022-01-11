package main

import "fmt"

const nilai string = "string paten"
var out_var string = "variabel di luar"
var (
	a = 1
	b = 2
	c = 3
)

func main() {

	var kata string = "Hello world"
	var txt string;
	
	txt += "Hello"
	txt += " World"
	
	x := 10
	
	fmt.Println(kata)
	fmt.Println(txt)
	fmt.Println("variabel x", x)
	fmt.Println(out_var)
	fmt.Println(nilai)
	fmt.Println(a,b,c)
}