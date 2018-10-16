go run file.go - compile + run code
gofmt -r  - format and rewrite file 


### Arrays and slice internals

**Array**
```golang
var a [4]int
b := [...]int{1, 2}
```
Size and object type are constants
Default values
byte: 0
To open array element-wide(kek), use '...'

**Slice**
Build on top of array
```
sl := []int{1,2}

length := 0
capacity := 0
s := make([]int, length, capacity)
```
slice structure: *ptr* to array, *length* - length of segment, *capacity* - maximum length of segment
use *append*:
```
x := [1]int{1}
x = X.append(x,2)
```
or *copy*:
```
longer := make([]int, len(x), (cap(x)+1)*2)
copy(longer, x)
x = longer
```
to grow. All this functions *COPY* the underlying array. 
If any, even one-element slice is connected with array, the whole array exists.

### Functions
Pass parameters by value (except maps).
Any function can be stored as object.

```
func function1(param1 int) (out1 int)
func function1(param1, param2 int) (out1, out2) // named parameters and named outputs
func function1(all_params ...int) // variadic
```
In variadic functions all input params should be the same type.
```
func variadicFunc(inputs ...int){
	fmt.Println(inputs)
}
x := []int{1,2,3}
variadicFunc(x...) // should use ... to open
```

####Anonymous functions
```
func(in string) {return 10}("nobody") // "nobody" is input

```

**Function with special signature can be stored as type for golang type checker**
```
type funcWithStrInput func(string) (int, string)
```
Funny function closure(*closure_example*)

#### Defer function execution**
```
defer fmt.Println("I am defer")
```
Defer functions are executed **after** function where they were defined(host function), but before panic.
Input of such functions is calculated **during** host function evaluation
Defer functions are evaluated as stack. (*defer_evaluation*)

### Structures
```
type newStr struct {
	Id 	int,
	Id2 int
}
var obj newStr = newStr{
	Id: 1
}
x := newStr{1,2}
```
Structure can include other structures
```
type inner struct{
	id int 
	value string
}
type host struct{
	id int
	inner
}
```
Inner inside host should be constructed as  ```inner : inner{1, "val"}```
Host fields with same names have priority, but host.inner.id 

#### Functions of structures/types

```
func (p *Person) UpdateName(name string) {
	p.Name = name }
```
By ref/value from function (p \*Person), not from func call (UpdateName(person1))

Host functions have priority


### Package structure
if smt starts with Capital letter - it is exported(public), otherwise private

### Interfaces

cast interface to inner struct: interface.(structName). It is possible to cast interface to interface.