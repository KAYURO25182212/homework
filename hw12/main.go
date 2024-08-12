package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// 1.	Создание и вывод map
	// Описание: Создайте map для хранения строк и их длин, добавьте несколько элементов и выведите содержимое.
	//1
	strilen1 := make(map[string]int)
	strings := []string{"sun", "moon", "skyfall"}
	for _, s := range strings {
		strilen1[s] = len(s)
	}
	for str, length := range strilen1 {
		fmt.Printf("Строка: %s, Длина: %d\n", str, length)
	}
	//2
	strilen2 := map[string]int{
		"sea":   4,
		"earth": 6,
		"space": 6,
	}
	prover1 := []string{"sea", "elephan"}
	for _, key := range prover1 {
		if AproveKey(strilen2, key) {
			fmt.Printf("Ключ '%s' существует в map.\n", key)
		} else {
			fmt.Printf("Ключ '%s' не найден в map.\n", key)
		}
	}
	//3
	strilen3 := map[string]int{
		"snake":    5,
		"lion":   6,
		"horse": 7,
	}
	massage1(strilen3, "lion2", 7)
	for str, length := range strilen3 {
		fmt.Printf("Строка: %s, Длина: %d\n", str, length)
	}
	//4
	mMap1 := map[string]int{
		"ine": 1,
		"two":  2,
		"three":  3,
	}
	fmt.Println("До удаления:", mMap1)
	delet1(mMap1, "three")
	fmt.Println("После удаления:", mMap1)
}