go run file.go - compile + run code
gofmt -r  - format and rewrite file 

### For

```
MARK:
 for {}
```
MARK - mark of for. can call ```break MARK ```

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
[Interfaces/readtomorrow](https://www.tapirgames.com/blog/golang-interface-implementation)

### make/new var

make - for slices/maps/channels
```
s := make([]int, 10, 100)
````
new - allocate and zero memory, return pointer

### Garbage Collector
concurrent/tri-color/mark-sweep
**concurrent**:  ~no stop of the world 
**tri-color(white/gray/black)**: heap is a graph of objects. At the start of GC round all objects a white. Go throught root objects: global variables and call stack. mark all of them as gray. Select one grey object, go deeper by references and mark all white as gray. Repeate while no gray left(how they keep counter?). All white are unreachable and can be deleted.
**mark-sweep**: mark unreachable objects/get rid of them


### Coroutines
add **go** for function call to call it in separate goroutine

runtime.Goshed() - pass execution to other goroutine
runtime - package to control go runtime system and goroutines
Goshed() - yields the processor, allowing other goroutines to run

switching between goroutins is possible only on functions call

### Channels
pass data ownership to other goroutine
```golang
make(chan type, bufferize)

value := <-channel // read from channel
channel <- value //add value to channel
```
channels are not buffered by default
```channel <-chan int``` - only read channel
```channel chan<= int``` - only write channel

```select{case  val <- ch1}``` - switch for channels
select check channels for read/write.
select selects random case to evaluate

```ticker := time.NewTicker(time.Second)``` - new tick every second to ticker. 
get ticker channel: ticker.C 
```ticker.Stop()```
```time.Tick()``` - return channel. imposible to stop
```time.AfterFunc(1 * time.Second(), functionToCall)```- call function on tick

#### contexts
interface to control channels in context
```ctx, finish := context.WithCancel(context.Background())```
```finish()``` - function to stop context. push to ```context.Done()``` channel
```context.WithTimeout(context.Background(), workTime)```
workTime - time for context to live

#### sync
```sync.WaitGroup()``` - struct with counter inside
```waitgroup.Wait()``` - stop execution till counter in WaitGroup equal to zero
```waitgroup.Done()``` - decrement counter

```sync.Mutex()```
```mutex.Lock()``` - get mutex
```mutex.Unlock()``` - release mutex

```sync\atomic``` - package for atomic operations (build in)
```atomic.AddInt32()```

### JSON

```string, _ = json.Marshal(struct)```
```struct, _ = json.Unmarshal(string)```
Check Public/Private access
```
type X struct{
	val int \`json:"value, string" \` // change representation for json
}
```

### memory Pool
```sync.Pool```