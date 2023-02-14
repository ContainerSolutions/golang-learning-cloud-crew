package main

import (
	"fmt"
)


func Hello() string {
	return "Hello, world"
}

func Hey() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello()) // Println expects any interface and formats appropriately
}
