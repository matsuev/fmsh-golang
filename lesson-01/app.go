package main

import "fmt"

func main() {
	result, err := Divider(20, 12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func MyFunction(a, b int64) int64 {
	return a * b
}

func Divider(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}
