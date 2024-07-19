package main

import "strings"

func main() {
	PrintGreeting()
	DisplaySeparator()
	NotifyStart()
	println(GetWelcomeMassage())
	println(GetPiValue())
	println(IsServerActive())
	PrintNumber(25)
	GreetUser("kayuro")
	ToggleLight(true)
	println(add(1, 5))
	println(concat("kayumov ", "abubakr"))
	println(IsEven(5))
	adder := func(a, b int) int {
		c := a + b
		return c
	}
	println(adder(5, 6))
	concater := func(a, b string) string {
		return a + b
	}
	println(concater("Neka ", "Gad"))
	IsPositive := func(a int) bool {
		if a > 0 {
			return true
		} else {
			return false
		}
	}
	println(IsPositive)

	println(Calculate(8, 7, add))
	Execute(true, Excute)
	Apply(25, Aply)
	s := Multiplier(5)
	println(s(25))
	p := StringRepeater(5, "Nigga")
	println(p())
	r := ConditionalPrinter(true)
	r()
}
func PrintGreeting() {
	println("Hello, World")
}
func DisplaySeparator() {
	println("--------------------")
}
func NotifyStart() {
	println("Program Started")
}

func GetWelcomeMassage() string {

	return "welcome"
}
func GetPiValue() float32 {
	return 3.14
}
func IsServerActive() bool {
	return true
}
func PrintNumber(a int) {
	println(a)
}
func GreetUser(nick string) {
	println(nick, "welcome")
}
func ToggleLight(light bool) {
	if light {
		println("light on")
	} else {
		println("light of")
	}
}
func add(a, b int) int {
	return a + b
}
func concat(a, b string) string {
	return a + b
}
func IsEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}
func Calculate(a, b int, add func(a, b int) int) int {
	d := add(a, b)
	return d
}
func Execute(b bool, c func(a bool)) {
	c(b)
}
func Excute(a bool) {
	println(a)
}
func Apply(a int, b func(c int) int) int {

	return b(a)
}
func Aply(c int) int {
	return c
}
func Multiplier(factor int) func(a int) int {
	return func(a int) int {
		return a * factor
	}
}
func StringRepeater(n int, a string) func() string {
	return func() string {
		return strings.Repeat(a, n)
	}
}
func ConditionalPrinter(condition bool) func() {
	return func() {
		if condition {
			println("Aproved")
		}
	}
}
