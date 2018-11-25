package main

import (
	"fmt"
)

func arrayChangeReference(a *[3]int){
	// a = real_array[1,2,3]
	a[0] = 10 // real_array[0] = 10
	a = &[3]int{99, 99, 99}  // a = *real_array[99,99,99]
	fmt.Println(a)
}

func arrayChangeCopy(a [3]int ){
	a[0] = 99
}

func arrayCallByCopyProof(){
	array := [3]int{1,2,3}
	fmt.Println(array)
	arrayChangeCopy(array)
	fmt.Println(array)

	array = [3]int{1,2,3} // array = real_array[1,2,3]
	arrayChangeReference(&array) // *real_array[1,2,3]
	fmt.Println(array)
}
// RESULT: whole array copy in function
// ===================================================================
func sliceChage(slice []int){
	slice[0] = 100 // real_array[0] = 100
	slice = append(slice, 99) // slice -> new_real_array[100,2,3,99]
}

func sliceNoArrayCopyProof(){
	slice := []int{1,2,3} // slice -> real_array[1,2,3]
	fmt.Println(slice)
	sliceChage(slice)
	fmt.Println(slice)
}
// RESULT: Copy of slice structure, not inner array. Same as map
// ===================================================================

type MyStruct struct{
	a int
	b []int
	c *[]int
	z [3]int
}

func changeStruct(my MyStruct){
	my.a = 66 // will not be changed
	my.b[0] = 77 
	(*my.c)[0] = 99
	my.z[0] = 88 // will not be changed
}

func myStructCopyProof(){
	my := MyStruct{1,[]int{1}, &[]int{1,1}, [3]int{6,6,6} }
	fmt.Println(my)
	fmt.Println(*my.c)
	changeStruct(my)
	fmt.Println(my)
	fmt.Println(*my.c)
}
// RESULT: MyStruct pass by value
// ===================================================================

func main(){
	//arrayCallByCopyProof()
	//sliceNoArrayCopyProof()
	myStructCopyProof()

}