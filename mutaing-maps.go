package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	m["Question"] = 52
	fmt.Println("The value:", m["Question"])

	v_a, ok := m["Answer"]
	fmt.Println("The value:", v_a, "Present?", ok)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v_a, ok = m["Answer"]
	fmt.Println("The value:", v_a, "Present?", ok)

	v_q, ok := m["Question"]
	fmt.Println("The value:", v_q, "Present?", ok)
}
