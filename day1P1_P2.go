package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)
const FILE = "./input.txt"


func main() {
	file, err := os.ReadFile(FILE)	
	check(err)
	inputArray := strings.Split(string(file), "\n")
	listA, listB:= splitIntoSortedArrays(inputArray)	
	total := partOne(listA, listB)
	fmt.Println(total)
}

func partOne(listA []int, listB []int)(int){
	var total int = 0 
	for i:= range listA {
		if listA[i] > listB[i] {
			total += listA[i] - listB[i]
			} else {
				total += listB[i] - listA[i]				
		}
	}
	return total
}

func check(e error){
	if e != nil {
		panic(e)
	}
}


func splitIntoSortedArrays(inputArray []string)([]int, []int){
	listA := make([]int, len(inputArray))
	listB := make([]int, len(inputArray))
	
	for i, item:= range inputArray {
		if(len(item) > 0){
			listA[i], _ = strconv.Atoi(strings.Fields(item)[0])							
			listB[i], _ = strconv.Atoi(strings.Fields(item)[1])		
		}
	}

	slices.Sort(listA)
	slices.Sort(listB)
	return listA, listB
}
