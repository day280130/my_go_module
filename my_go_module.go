package mygomodule

import "fmt"

func SayHelloTo(target string) {
	fmt.Printf("Hello %s!", target)
}

func GenerateHiTo(target string) string {
	return "Hi " + target + "!"
}
