package main
import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
	x := 42

	if x > 18 {
		fmt.Println("Allowed ")
	}
	// From https://go.dev/doc/tutorial/getting-started#install
	fmt.Println(quote.Go())

}
