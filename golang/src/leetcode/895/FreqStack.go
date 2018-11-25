package main

import (
    "fmt"
)

// type Elem struct {
//     counter int
//     positions chan int
// }

type FreqStack struct {
    heapSize int
    elems [10]int
}

func (this *FreqStack) print(){
    fmt.Println("Heap: ", this.heapSize)
    fmt.Println(this.elems)
}

func Constructor() FreqStack {
    return FreqStack{}

}


func (this *FreqStack) Push(x int)  {
    this.elems[this.heapSize] = x
    curPos := this.heapSize
    parentPos := (this.heapSize-1)/2
    for {
        if x <= this.elems[parentPos]{
            break
        }
        this.elems[curPos], this.elems[parentPos] = this.elems[parentPos], this.elems[curPos]
        curPos = parentPos
        parentPos = (parentPos-1)/2
    }
    this.heapSize += 1
}


func (this *FreqStack) Pop() int {
    result := this.elems[0]
    this.elems[0] = this.elems[this.heapSize-1]
    this.elems[this.heapSize-1] = 0
    this.heapSize -= 1
    this.Heapify(0)
    return result
}

func (this *FreqStack) Heapify(pos int) {
    leftPos := getLeftPos(pos)
    rightPos := getRightPos(pos)
    largestPos := pos
    
    if this.elems[leftPos] > this.elems[largestPos]{
        largestPos = leftPos
    }
    if this.elems[rightPos] > this.elems[largestPos]{
        largestPos = rightPos
    }
    if largestPos != pos{
         this.elems[largestPos], this.elems[pos] = this.elems[pos], this.elems[largestPos]
         this.Heapify(largestPos)
    }

}
func getLeftPos(pos int) int {
    return pos*2 + 1
}
func getRightPos(pos int) int {
    return pos*2 + 2
}

func (this *FreqStack) setElems(elems [10]int, size int){
    this.elems = elems
    this.heapSize = size
}

func main() {
    
    stack := FreqStack{}
    // stack.Push(1)
    // stack.Push(10)
    // stack.print()
    // stack.Push(11)
    // stack.Push(12)
    // stack.print()
    // stack.Push(3)
    // stack.Push(2)
    // stack.Push(20)
    stack.setElems([10]int{20,11,12,1,3,2,10,0,0,0}, 7)

    stack.print()
    fmt.Println(stack.Pop())
    stack.print()

    fmt.Println(stack.Pop())
    stack.print()
    // fmt.Println((6-1)/2)
}