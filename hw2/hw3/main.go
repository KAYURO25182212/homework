package main

/*import (
	"fmt"
)*/
func main() {
	x := 2
	y := 3
	a := 12
	b := 25
	word1 := "Hello"
	word2 := "Hello"
	PrintGreeting()
	PrintWeather()
	PrintTrafficLight()
	GetGrade()
	GetDiscount()
	GetTemperatureDescription()
	CheckNumber()
	CheckAge()
	CheckPassword()
	println(add(x, y))
	println(CompareStrings(word1, word2))
	println(Max(a, b))
	operation := func(a, b int) int {
		c := a - b
		if c < 0 {
			c = -c
		}
		return c
	}
	println(operation(a, b))
	concat := func(m, j string) string {
		c := m + j
		return c
	}
	println(concat("Hello \n", "What's up \n"))

	multiply := func(a, b int) int {
		a = 2
		b = 4
		c := a + b
		return c
	}
	println(multiply(2, 4))
	println(ApplyOperation(24, 23, apply))
}
func PrintGreeting() {
	var DayType string = "vecher"
	if DayType == "utro" {
		println("Dobroe", DayType)
	} else if DayType == "vecher" {
		println("dobriy", DayType)
	} else {

	}
}

func PrintWeather() {
	var weatherType string = "oblachno"
	switch weatherType {
	case "solnechno":
		println("Pogoda", weatherType)
	case "oblachno":
		println("Pogoda", weatherType)
	case "dozhdlivaya":
		println("Pogoda", weatherType)
	}

}

func PrintTrafficLight() {
	var lightColor string = "red"
	switch lightColor {
	case "red":
		println("Ready")
	case "yelow":
		println("Set")
	case "green":
		println("Go")
	}
}

func GetGrade() int {
	var score int = 99
	if score == 100 {
		println("Your score A+")
	} else if score > 90 && score < 100 {
		println("Your score A")
	} else if score > 75 && score < 90 {
		println("Your score B")
	} else if score > 50 && score < 75 {
		println("Your score B")
	} else {
		println("You failed test")
	}
	return score
}

func GetDiscount() int {
	var amount int = 2016
	if amount > 1000 {
		a := amount * 10 / 100
		println("You have 10% discount, your discount is ", a, "$")
	} else if 500 < amount && amount <= 1000 {
		b := amount * 5 / 100
		println("You have 5%, discount, your discount is ", b, "$")
	} else if amount < 500 {
		println("You dont have discount your discount is ", amount, "$")
	} else {
		println("error")
	}
	return amount
}

func GetTemperatureDescription() int {
	var temperature int = 25
	if temperature < 10 && temperature > 0 {
		println("Its too cold")
	} else if temperature > 10 && temperature <= 25 {
		println("its warm")
	} else if temperature > 25 {
		println("its to hot")
	} else {
		println("You cant live here")
	}
	return temperature
}

func CheckNumber() int {
	number := 25
	if number > 0 {
		println("Chislo polozhitelnoe")
	} else if number <= 0 {
		println("Chislo otricatelnoe")
	} else {
		println("Error")
	}
	return number
}

func CheckAge() {
	age := 18
	if age >= 18 {
		println("you can go")
	} else {
		println("Get out ")
	}
}

func CheckPassword() {
	password := "qwert"
	if password == "qwert1234" {
		println("Acces denied")
	} else {
		println("Wrong password")
	}
}

func add(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	c := x + y
	return c
}

func CompareStrings(word1, word2 string) string {
	if word1 == word2 {
		return "everything ok"
	} else {
		return "Something went wrong"
	}

}
func Max(a, b int) int {
	a = 20
	b = 25
	if a > b {
		return a
	} else if b > a {
		return b
	} else {

	}
	return 0

}

func ApplyOperation(a, b int, c func(a, b int) int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return c(a, b)

}
func apply(a, b int) int {
	c := a + b
	return c
}
