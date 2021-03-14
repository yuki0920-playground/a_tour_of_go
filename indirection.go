package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// vをレシーバとするメソッド
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 普通の関数
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v1 := Vertex{3, 4}
	v1.Scale(2)
	ScaleFunc(&v1, 10)

	v2 := Vertex{3, 4}
	// v2.Scalae(3) と同じ
	(&v2).Scale(3)
	ScaleFunc(&v2, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v1, v2, p)
}
