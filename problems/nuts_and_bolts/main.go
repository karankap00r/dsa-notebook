package main

import (
	"errors"
	"fmt"
)

func main() {
	var count int
	fmt.Print("Enter total count of nuts and bolts: ")
	fmt.Scanf("%d", &count)

	fmt.Print("Enter sizes of the nuts: ")
	nuts := getIntArrayInput(count)

	fmt.Print("Enter sizes of the bolts: ")
	bolts := getIntArrayInput(count)

	solver := DefaultNutsAndBoltsSolver{}
	err := solver.Solve(nuts, bolts)
	if err != nil {
		fmt.Print("The nuts and bolts solver could not solve")
		return
	}

	fmt.Println("Nuts array: ")
	printArray(nuts)

	fmt.Println()

	fmt.Println("Bolts array: ")
	printArray(bolts)
}

type Pair struct {
	NutSize, BoltSize int
}

type NutsAndBoltsSolver interface {
	Solve(nuts, bolts []int) error
}

type DefaultNutsAndBoltsSolver struct{}

func (solver DefaultNutsAndBoltsSolver) Solve(nuts, bolts []int) error {
	if len(nuts) != len(bolts) {
		return errors.New("Length of nuts and bolts did not match.")
	}

	if len(nuts) == 0 {
		return nil
	}

	nutsPivot := 0
	targetBoltIndex := -1

	smallerBolts := make([]int, 0)
	largerBolts := make([]int, 0)
	for idx, val := range bolts {
		if val < nuts[nutsPivot] {
			smallerBolts = append(smallerBolts, val)
		} else if val > nuts[nutsPivot] {
			largerBolts = append(largerBolts, val)
		} else {
			targetBoltIndex = idx
		}
	}

	pivotNut := nuts[nutsPivot]
	pivotBolt := bolts[targetBoltIndex]

	smallerNuts := make([]int, 0)
	largerNuts := make([]int, 0)
	for _, val := range nuts {
		if val < bolts[targetBoltIndex] {
			smallerNuts = append(smallerNuts, val)
		} else if val > bolts[targetBoltIndex] {
			largerNuts = append(largerNuts, val)
		}
	}

	leftErr := solver.Solve(smallerNuts, smallerBolts)
	if leftErr != nil {
		return leftErr
	}

	rightErr := solver.Solve(largerNuts, largerBolts)
	if rightErr != nil {
		return rightErr
	}

	index := 0
	for idx, _ := range smallerNuts {
		nuts[index] = smallerNuts[idx]
		bolts[index] = smallerBolts[idx]
		index++
	}

	nuts[index] = pivotNut
	bolts[index] = pivotBolt
	index++

	for idx, _ := range largerNuts {
		nuts[index] = largerNuts[idx]
		bolts[index] = largerBolts[idx]
		index++
	}

	return nil
}

/* util functions */

// printArray is used to print an array input
func printArray(arr []int) {
	if len(arr) == 0 {
		return
	}

	for _, val := range arr {
		fmt.Print(val)
		fmt.Print(" ")
	}
	fmt.Println()
}

func getIntArrayInput(size int) []int {
	if size == 0 {
		return make([]int, 0)
	}

	arr := make([]int, 0)
	for i := 0; i < size; i++ {
		var num int
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}
	return arr
}
