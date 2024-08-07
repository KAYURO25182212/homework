package main

import (
	"fmt"
	"strings"
)

func main() {
	//1 Конкатенация строк. Напишите функцию, которая принимает две строки и возвращает их конкатенацию.

	fmt.Println(concat("Dark ", "Knight"))
	//2 Длина строки. Напишите функцию, которая принимает строку и возвращает её длину.
	fmt.Println(Len("Hello"))
	//3. Проверка подстроки. Напишите функцию, которая проверяет, содержит ли строка заданную подстроку.
	Podstroki()
	//4. Поиск подстроки. Напишите функцию, которая находит индекс первого вхождения подстроки в строке. Верните -1, если подстрока не найдена.
	Podstrok2i()
	//5. Замена подстроки. Напишите функцию, которая заменяет все вхождения подстроки в строке на другую подстроку.
	podstrok3i()
	//6. Удаление пробелов. Напишите функцию, которая удаляет пробелы в начале и в конце строки.
	spaces()
	//7. Преобразование регистра. Напишите функцию, которая преобразует строку к верхнему и нижнему регистру.
	Regist()
}

// 1
func concat(s1, s2 string) string {
	s3 := s1 + s2
	return s3
}

// 2
func Len(a string) int {
	return len(a)
}

// 3
func Podstroki() {
	str := "Hello world!"
	fmt.Println(strings.Contains(str, "world"))
}

// 4
func Podstrok2i() int {
	str2 := "Hi my nigga"
	fmt.Println(strings.Index(str2, "nigga"))
	return -1
}

// 5
func podstrok3i() {
	str3 := "Stan Exselsior"
	replaced := strings.Replace(str3, "Stan", "Marvel", -1)
	fmt.Println(replaced)
}

//6
func spaces(){
	spc := "Michel Jordan"
	fmt.Println(strings.TrimSpace(spc))
}

//7
func Regist(){
	reg := "The Shot"
	fmt.Println(strings.ToUpper(reg))
	fmt.Println(strings.ToLower(reg))
}