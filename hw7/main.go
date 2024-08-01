package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	maxvalue := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > maxvalue {
			maxvalue = a[i]
		}
	}
	println(maxvalue)

	b := [5]int{1, 2, 3, 4, 5}
	minvalue := b[0]
	for i := 0; i > len(b); i++ {
		if b[i] < minvalue {
			minvalue = b[i]
		}
	}
	println(minvalue)

	c := [6]int{-1, -2, 3, 4, 5}
	for i := 0; i < len(c); i++ {
		if c[i] > 0 {
			println(c[i])
		}
	}
	d := [5]int{1, 2, 3, 4, -2}
	sum := 0
	for summ := range d {
		sum += summ
	}
	fmt.Printf("all in: %d\n", sum)
	g := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	grit := 0
	for gritt := range g {
		grit += gritt
	}
	fmt.Printf("middle numb : %.2f\n", float64(grit)/float64(len(g)))

	h := [5]int{1, 2, 3, 4, 5}
	h2 := 2
	for hhh := range h {
		h[hhh] *= h2
	}
	fmt.Println("all numbers vas  on  ", h)

	l := []int{1, 2, 3, 4, 5}
	l2 := 2
	l3 := []int{}
	for i, l3 := range l {
		if l3 == l2 {
			l3 = append(l3, i)
		}
	}
	fmt.Printf("Индексы числа %d в массиве: %v\n", l2, l3)
}
