package main

import "fmt"

func main() {
	println(temp(2))
	println(agerunner(19))
	println(speedtest(118))
	println(Scorem(25))
	println(IBM(75))
	var un UnaryOperation = Una
	println(un(2))
	var doub UnaryOperation = double
	println(doub(3))
	var ch check
	ch = Check
	println(ch(4))
	var pl check
	pl = pol
	println(pl(-3))
	count(5)
	var z Operation
	z[0] = plus
	z[1] = minus
	z[2] = equel
	println(plus(5, 5), minus(4, 2), equel(3, 3))
	println(Battery(25))
	println(wieght(75))
	println(grade(100))
	println(Blood(70, 70))
}

type temperature float64

func temp(a temperature) string {
	if a > 0 {
		println(a)
		return "temperatura Bolshe 0"
	} else if a < 0 {
		println(a)
		return "Temperatura menshe 0"
	} else {

	}
	return "oshibka"
}

type age int

func agerunner(a age) string {
	if a >= 13 && a <= 19 {
		println(a)
		return "Podrostok"
	} else if a > 19 {
		println(a)
		return "Nepodrostok"
	} else if a < 13 {
		println(a)
		return "rebenok"
	} else {

	}
	return "error"
}

type speed float64

func speedtest(a speed) string {
	if a >= 0 && a <= 120 {
		println(a, "km/h")
		return "skorost dopustima"
	} else if a > 120 {
		return "Slishkokm bistro"
	} else {

	}
	return "oshibka"
}

type Score int

func Scorem(a Score) string {
	if a == 0 {
		println(a)
		return "Chislo Nulevoe"
	} else if a > 0 {
		println(a)
		return "chislo polozhitelnoe "
	} else if a < 0 {
		println(a)
		return "chislo otricatelnoe"
	} else {

	}
	return "oshibka"
}

type BMI float64

func IBM(a BMI) string {
	if a > 85 && a < 100 {
		println(a)
		return "Overweight"
	} else if a <= 85 && a >= 60 {
		println(a)
		return "Normal weight"
	} else if a < 60 {
		println(a)
		return "Underweight"
	} else if a > 100 {
		println(a)
		return "Obesity"
	} else {

	}
	return "Error"
}

type UnaryOperation func(int) int

func Una(a int) int {
	return a * a
}

func double(a int) int {
	return a * 2
}

type check func(int) bool

func Check(a int) bool {
	if a%2 == 1 {
		return false
	} else {
		return true
	}
}
func pol(a int) bool {
	if a >= 0 {
		return true
	} else {
		return false
	}
}

type Operation [3]func(int, int)int
func plus(a,b int) int {
	return a + b
}

func minus(a,b int) int {
	return a - b
}

func equel(a,b int) int {
	return a * b
}

type Count int

func count(a Count) {
	for ; a >= 0; a-- {
		fmt.Println(a)
	}
}

type BatteryLevel int 
func Battery(a BatteryLevel) string {
	if a > 0 && a <=19 {
		println(a,"%")
		return "Battery low"
	} else if a >=20 && a <= 60 {
		println(a,"%")
		return "Battery medium"
	} else if a> 61 && a <=100 {
		println(a,"%")
		return "Battery High"
	} else {
		return "Error"
	}
}
type Weight float64 
func wieght(a Weight) string{
	if a > 0 && a <= 30 {
		println(a)
		return "Light"
	} else if a >= 31 && a <=75 {
		println(a)
		return "Medium"
	} else if a >= 76 {
		println(a)
		return "Heavy"
	} else {
		return "error"
	}
}
type Grade int
func grade(a Grade) bool {
	if a >= 50 {
		return true
	} else {
		return false
	}
}

type BloodPressure float64 
func Blood(a BMI, b BloodPressure) string {
	if a > 85 && b > 85  {
		return "Unhealthy"
	} else if a <= 85 && a >= 60 && b <= 85 && b >= 60{
		return "Healthy"
	} else if a < 60 {
		return "At risk"
	}else {
		return "error"
	}
}