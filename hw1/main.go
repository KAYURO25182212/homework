package main

import (
	"fmt"
)

func main() {
	dz()
	dz2()
	dz3()
	dz4()
	dz5()
	dz6()
	dz7()
	dz8()
	dz9()
	dz10()
	dz11()
	dz12()
	dz13()
	dz14()
	dz15()
}
func dz() {
	var a int = 10
	p := a * 2
	fmt.Println(p)
}
func dz2() {
	var a int = 3
	s := a
	fmt.Println(s)
}
func dz3() {
	var a int = 10
	var b int = 2
	S := a * b
	P := 2 * (a + b)
	fmt.Println(S)
	fmt.Println(P)
}
func dz4() {
	var d float32 = 3
	var p float32 = 3.14
	L := p * d
	fmt.Println(L)
}
func dz5() {
	var a int = 3
	v := a * 3
	s := 6 * a
	fmt.Println(v)
	fmt.Println(s)
}
func dz6() {
	var (
		a int = 2
		b int = 3
		c int = 4
	)
	v := a * b * c
	s := 2 * (a*b + b*c + a*c)
	fmt.Println(v)
	fmt.Println(s)
}
func dz7() {
	var r float32 = 2
	var p float32 = 3.14
	l := 2 * p * r
	s := p * r
	fmt.Println(l)
	fmt.Println(s)
}
func dz8() {
	var a int = 2
	var b int = 4
	fmt.Println((a * b) / 2)
}
func dz9() {
	var a int = 6
	var b int = 3
	fmt.Println((a * b) / 1 / 2)
}
func dz10() {
	var a float32 = 8
	var b float32 = 4
	fmt.Println(a*a + b*b)
	fmt.Println(a*a - b*b)
	fmt.Println((a * a) * (b * b))
	fmt.Println((a * a) / (b * b))
}
func dz11() {
	var L float64 = 835
	M := L / 100
	fmt.Println(M)
}
func dz12() {
	var kg float64 = 25356
	t := kg / 1000
	fmt.Println(t)
}
func dz13() {
	var bytes float64 = 23456
	k := bytes / 1024
	fmt.Println(k)
}
func dz14() {
	var a uint16 = 10
	var b uint16 = 2
	fmt.Println(a / b)
}
func dz15() {
	var a uint16 = 20
	var b uint16 = 2
	fmt.Println(a % b)
}
