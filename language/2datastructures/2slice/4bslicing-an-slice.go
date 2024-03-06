// slice: slicing an slice

package main

import (
	"fmt"
)

func main() {

	var myarray [10]int

	myarray = [10]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	var myOriginalSlice []int = myarray[1:5]
	fmt.Println("myOriginalSlice", myOriginalSlice, len(myOriginalSlice), cap(myOriginalSlice))

	myslice := myOriginalSlice[1:3] // myOriginalSlice[i:j] i left boundary (inclusive). j is right boundary (excluding, i.e included till j-1)
	fmt.Println("myslice", myslice, len(myslice), cap(myslice))

	myslice1 := myOriginalSlice[1:] // right default is len. myOriginalSlice[1:len(myOriginalSlice)] or myOriginalSlice[1:4]
	fmt.Println("myslice1", myslice1, len(myslice1), cap(myslice1))

	myslice3 := myOriginalSlice[:3] // left default is 0 index
	fmt.Println("myslice3", myslice3, len(myslice3), cap(myslice3))

}
