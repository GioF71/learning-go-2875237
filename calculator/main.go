package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func get(prompt string, valueName string) float64 {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	var s1 string
	var v1 float64 = 0.0
	var err error = nil
	
	s1,err = reader.ReadString('\n')
	v1,err = strconv.ParseFloat(strings.TrimSpace(s1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	return v1
}



func main() {
	fmt.Println("A simple calculator")

	//reader := bufio.NewReader(os.Stdin)
	//var v1 float64 = 0.0
	//var v2 float64 = 0.0
	//var err error = nil

	//fmt.Printf("v1 %f v2 %f\n", v1, v2)

	/*fmt.Print("Enter v1: ")
	s1,err := reader.ReadString('\n')
	v1,err = strconv.ParseFloat(strings.TrimSpace(s1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}*/
	v1 := get("Enter v1", "v1")
	v2 := get("Enter v2", "v2")

	/*fmt.Print("Enter v2: ")
	s2,err := reader.ReadString('\n')
	v2,err = strconv.ParseFloat(strings.TrimSpace(s2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}*/

	fmt.Printf("%f + %f = %f", v1, v2, v1 + v2)
}
