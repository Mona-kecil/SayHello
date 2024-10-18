package SayHello

import (
	"fmt"
	"strings"
)

func Hello(name string) {
	fmt.Println("Hello " + name)
}

func HelloFilter(name string) {
	if strings.Contains(strings.ToLower(name), "anjing") {
		fmt.Println("Hello ...")
	} else {
		fmt.Println("Hello " + name)
	}
}
