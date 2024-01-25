package mygomodule

import "fmt"

func SayHelloTo(target string, selfName string) {
	fmt.Printf("Hello %s! I'm %s", target, selfName)
}

func GenerateHiTo(target string) string {
	return "Hi " + target + "!"
}
