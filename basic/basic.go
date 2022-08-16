package basic

import (
	"fmt"
	"unsafe"

	geometry "github.com/itisaby/go-microservice/geometry"
	"rsc.io/quote"
)

func func1(x, y float64) (float64, float64) {
	area := x * y
	fmt.Println(area)
	perimeter := 2 * (x + y)
	fmt.Println(perimeter)
	fmt.Println(x)
	fmt.Println(y)
	return area, perimeter
}

func func2(x, y float64) (area float64) {
	area = x * y
	fmt.Println(area)
	return area
}

func func3() {
	var daysofMonth map[string]int
	daysofMonth = make(map[string]int)
	daysofMonth["January"] = 31
	daysofMonth["February"] = 28
	fmt.Println(daysofMonth)
}

func main() {
	x := 10
	var y int = 20
	name := "Arnab"
	fmt.Println("Hello World")
	fmt.Println(quote.Go())
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(name)
	fmt.Println(x, y, name)
	fmt.Printf("Type of name %T and size is %d", name, unsafe.Sizeof(name))
	fmt.Println()
	a, p := func1(10, 20)
	fmt.Println(a, p)
	a1 := func2(10, 20)
	fmt.Println(a1)
	h := geometry.Area(1, 20)
	d := geometry.Diagonal(10, 20)
	fmt.Println(h)
	fmt.Println(d)
	func3()

}
