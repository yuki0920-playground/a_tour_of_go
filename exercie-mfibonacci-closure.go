package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a, b := 1, 0
    return func() int {
				// NOTE: 初期値 0, 次に 1
        a, b = b, a+b
				fmt.Println(a, b)

        return a
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
