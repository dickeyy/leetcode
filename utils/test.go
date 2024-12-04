package utils

import (
	"fmt"
	"reflect"
)

const (
	colorReset = "\033[0m"
	darkGray   = "\033[90m"
	red        = "\033[31m"
	green      = "\033[32m"
)

func AssertEqual(expected, actual interface{}, testNum int) {
	// Check if types match
	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(actual)

	if expectedType != actualType {
		fmt.Printf("%sTest %d Failed%s %s\n    expected: %v but got: %v%s\n", red, testNum, colorReset, darkGray, expectedType, actualType, colorReset)
		return
	}

	// Deep equality check
	if !reflect.DeepEqual(expected, actual) {
		// println("❌ Test Failed!")
		// fmt.Printf("Expected: %v\nActual: %v\n", expected, actual)
		fmt.Printf("%sTest %d Failed%s %s\n    expected: %v but got: %v%s\n", red, testNum, colorReset, darkGray, expected, actual, colorReset)
		return
	}

	fmt.Printf("%sTest %d Passed%s\n", green, testNum, colorReset)
}
