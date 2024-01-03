package main
import "fmt"
import "rsc.io/quote"

func frequencies() {
	var f int
	fmt.Scanln(&f)
	//your code goes here
	switch {
	  case f >0 && f<20:
		fmt.Println("Infrasound")
	  case f>=20 && f<=20000:
		fmt.Println("Audible")
	  case f>20000:
		fmt.Println("Ultrasound")
	  default:
		fmt.Println("Wrong Input")
	}
  }

func main() {
	fmt.Println("Hello, World!")
	x := 42

	if x > 18 {
		fmt.Println("Allowed ")
	}
	// From https://go.dev/doc/tutorial/getting-started#install
	fmt.Println(quote.Go())
	frequencies()

}
