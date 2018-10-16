package main
import "fmt"

// func x(in int) int{return 10}


type strFuncType func(int) int

func pp(value string) string {
	fmt.Println("pp: " + value)
	return value
}

func main() {
	// fmt.Println("hello world!")
	// summer := func(x int) (func(int) int){
	// 	return func(plusValue int)(int){
	// 		return plusValue+x
	// 	}
	// }
	// sum3 := summer(3)
	// fmt.Println(sum3(10))

	defer fmt.Println(pp("first"))
	fmt.Println("main")
	defer fmt.Println(pp("third"))

}


// prefixer := func(prefix string) strFuncType {
// 	return func(in string) {
// 		fmt.Printf("[%s] %s\n", prefix, in)
// 	}
// }
// successLogger := prefixer("SUCCESS")
// successLogger("expected behaviour")