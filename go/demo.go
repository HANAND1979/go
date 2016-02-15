// example1
//on branch md5
package main

//importing packages
import (
	//"crypto/md5"
	"fmt"
	//"io"
	"math"
	"math/rand"
	"time"
)

var hello string
var gdp int

//Basic arithmetic functions
func main() {
	fmt.Println("Hello Git World!")
	hello = "Hello World!"
	fmt.Println(hello)

	a, b := 8.1, 9.7 //short variable declaration
	fmt.Println(a * b)
	fmt.Println(math.Pi)
	fmt.Println(rand.Intn(7))

	//type conversions
	var x, y int = 7, 9
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("Type conversion", x, y, f, z)

	//Demo of function
	m, n := func_demo("Google")
	fmt.Println(m, n)

	//Demo of control statements
	defer func_for(3) //illustration of defer
	if_output := func_if(9, 9)
	fmt.Println(if_output)
	func_case()

	//Demo of MD5 hash
	//func_hashmd5("Wow!Password")
}

//demo of function
func func_demo(s string) (string, int) {

	a := "Hello" + " " + s
	b := len(s)
	return a, b
}

func func_for(counter int) {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += i
		defer print(sum, "\n") //use of defer inverts the output
	}

}

func func_if(p, q int) bool {

	if p > q {
		return true
	} else {
		return false
	}

}

func func_case() {
	fmt.Print("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("One day away")
	case today + 2:
		fmt.Println("Two days away")
	default:
		fmt.Println("Too far away")

	}
}

//This function demonstrates the usage of crypto function for MD5 hash
//func func_hashmd5(password string) string {

//	p := md5.New()
//	return io.WriteString(p, password)

//}
