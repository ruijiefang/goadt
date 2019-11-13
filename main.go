package main

import (
	"fmt"
)
import ("goadt/storage")

func testInsertion(i uint32) bool {
	var list = new(storage.LinkedList)
	for j := uint32(0); j < i; j++ {
		list.Add(j)
		if list.Size() != j + 1 {
			fmt.Printf("Error: Size mismatch. Actual size is %v whereas module claims %v\n", j, list.Size())
			return false
		}
		if !list.Exists(j) {
			fmt.Printf("Error: Cannot find inserted element [%v]\n", j)
			return false
		}
	}
	for j := uint32(0); j < i; j++ {
		if !list.Remove(j) {
			fmt.Printf("Error: Cannot remove element [%v]\n", j)
			return false
		}
		if list.Exists(j) {
			fmt.Printf("Error: Removed element [%v] exists in list.\n", j)
			return false
		}
	}
	return true
}

func main() {
	/*
	for i := uint32(1) ; i < 2000; i++ {
		fmt.Printf(" ========= \n")
		var j = testInsertion(i)
		fmt.Printf("^ Tested [%v] insertions (Return status: %v)\n", i, j)
		if !j {
			return
		}
	} */
	var vec []uint32 = make([]uint32, 0)
	vec = append(vec, 1)
	vec = append(vec, 2)
	vec = append(vec, 3)
	vec = append(vec, 4)
	vec = append(vec, 5)
	fmt.Printf("%v\n", vec)
	vec = vec[:len(vec) - 1]
	fmt.Printf("%v\n", vec)

}