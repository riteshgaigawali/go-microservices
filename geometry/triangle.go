package geometry

import "fmt"

var str string = "The area of triangle is " //Variable name starts with small letters indicate private access specifier.

func Area(base, height float64) float64 { //Method name starts with capital letter indicates public access specifier.
	area := 0.5 * base * height
	fmt.Print(str)
	return area
}
