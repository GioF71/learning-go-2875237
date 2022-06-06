package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_from_reader(reader *bufio.Reader, valueName string) float64 {
	fmt.Printf("Enter %s: ", valueName)
	var inputString string
	var theValue float64 = 0.0
	var theError error = nil

	inputString, theError = reader.ReadString('\n')
	if theError != nil {
		fmt.Println(theError)
		panic(fmt.Sprintf("Cannot read value for [%s]", valueName))
	}
	theValue, theError = strconv.ParseFloat(strings.TrimSpace(inputString), 64)
	if theError != nil {
		fmt.Println(theError)
		panic(fmt.Sprintf("Value for [%s] must be a number", valueName))
	}
	return theValue
}

func get(valueName string) float64 {
	reader := bufio.NewReader(os.Stdin)
	result := get_from_reader(reader, valueName)
	return result
}

func main() {
	fmt.Println("A simple calculator")
	v1 := get("v1")
	v2 := get("v2")
	fmt.Printf("%f + %f = %f\n", v1, v2, v1+v2)
}
