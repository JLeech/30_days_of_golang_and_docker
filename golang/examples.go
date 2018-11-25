package main
import "fmt"
// ===========================================================================================================
// defer_evaluation

// func pp(value string) string {
// 	fmt.Println("pp: " + value)
// 	return value
// }

// func main() {
// 	defer fmt.Println(pp("first"))
// 	fmt.Println("second")
// 	defer fmt.Println(pp("third"))

// }

// -> pp: first
// -> second
// -> pp: third
// -> third
// -> first
// ===========================================================================================================

// closure_examples
// func main() {
// 	summer := func(value int) (func(int) int){
// 	  return func(plusValue int)(int){
// 	    return plusValue+value
// 	  }
// 	}
// 	sum3 := summer(3)
// 	fmt.Println(summer(3)(10))
// 	fmt.Println(sum3(10))
// }

// type newStr struct {
// 	Id 	int
// 	Id2 int
// }
// func main(){
// var obj newStr = newStr{
// 	Id: 1,
// }
// ss := newStr{1,2}

// fmt.Println(obj)
// fmt.Println(ss)
// }

func main(){
	a := new([]int)
	fmt.Println(&a)
	*a = append(*a, 3)
	fmt.Println(&a)
	fmt.Println(*a)

}