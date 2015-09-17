package fibonacci

//import "fmt"

func fibSeries(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibSeries(n-2) + fibSeries(n-1)
	}
}

/*
func main() {

	fmt.Print("Enter a number: ")
	var x int
	fmt.Scanln("%d", &x)
	fmt.Println(fibSeries(x))

}
*/