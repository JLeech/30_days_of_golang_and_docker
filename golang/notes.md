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
Pass parameters by value.
