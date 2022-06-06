package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OperationType func(float64, float64) float64

func sum(v1 float64, v2 float64) float64 {
	return v1 + v2
}

func subtract(v1 float64, v2 float64) float64 {
	return v1 - v2
}

func multiply(v1 float64, v2 float64) float64 {
	return v1 * v2
}

func divide(v1 float64, v2 float64) float64 {
	return v1 / v2
}

func getUserInput(inputName string) string {
	reader := bufio.NewReader(os.Stdin)
	var inputString string
	var theError error = nil

	inputString, theError = reader.ReadString('\n')
	if theError != nil {
		fmt.Println(theError)
		panic(fmt.Sprintf("Cannot read user input for [%s]", inputName))
	}
	return strings.TrimSpace(inputString)
}

func getValue(valueName string) float64 {
	fmt.Printf("Enter %s: ", valueName)
	inputString := getUserInput(valueName)
	var theError error = nil
	var theValue float64 = 0.0
	theValue, theError = strconv.ParseFloat(strings.TrimSpace(inputString), 64)
	if theError != nil {
		fmt.Println(theError)
		panic(fmt.Sprintf("Value for [%s] must be a number", valueName))
	}
	return theValue
}

func getOperation(validOperationArray []string) string {
	valueName := "operation"
	fmt.Printf("Enter %s: ", valueName)
	inputString := getUserInput(valueName)
	if len(inputString) != 1 {
		panic(fmt.Sprintf("Operation length is %d, must be 1", len(inputString)))
	}
	success := false
	for _, validEntry := range validOperationArray {
		if strings.Compare(validEntry, inputString) == 0 {
			success = true
			break
		}
	}
	if !success {
		panic(fmt.Sprintf("Operation %s is not allowed", inputString))
	}
	return inputString
}

func main() {
	operations := []string{"+", "-", "*", "/"}

	opDescMap := make(map[string]string)
	opDescMap[operations[0]] = "sum"
	opDescMap[operations[1]] = "subtract"
	opDescMap[operations[2]] = "multiply"
	opDescMap[operations[3]] = "divide"

	fmt.Println(operations)

	opMap := make(map[string]OperationType)
	opMap[operations[0]] = sum
	opMap[operations[1]] = subtract
	opMap[operations[2]] = multiply
	opMap[operations[3]] = divide

	fmt.Println("A simple calculator")
	v1 := getValue("v1")
	v2 := getValue("v2")
	operation := getOperation(operations)
	fmt.Printf("Selected operation is: [%s]\n", opDescMap[operation])
	opType := opMap[operation]
	fmt.Printf("%f + %f = %f\n", v1, v2, opType(v1, v2))
}
