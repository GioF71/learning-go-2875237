package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")

	reader := bufio.NewReader(os.Stdin)
	var v1 float64 = 0.0
	var v2 float64 = 0.0
	
	fmt.Printf("v1 %f v2 %f\n", v1, v2)

	fmt.Print("Enter v1: ")
	s1,_ := reader.ReadString('\n')
	v1,_ = strconv.ParseFloat(strings.TrimSpace(s1), 64)

	fmt.Print("Enter v2: ")
	s2,_ := reader.ReadString('\n')
	v2,_ = strconv.ParseFloat(strings.TrimSpace(s2), 64)

	fmt.Printf("%f + %f = %f", v1, v2, v1 + v2)
}
