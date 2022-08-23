package main

import (
	"fmt"
	"os"
)

func main() {

	var (
		sortYes                 []int
		lowNumIndRec, lowNumRec int
		sortMe                  []int
	)

	sortMe = createToSortSlice()

	fmt.Printf("Unsorted slice: %T %v\n", sortMe, sortMe)

	writeSliceToFile(sortMe)

	sortMeLength := len(sortMe)

	for i := 0; i < sortMeLength; i++ {
		lowNumIndRec, lowNumRec = findLowestNum(sortMe) // find lowest
		sortYes = append(sortYes, lowNumRec)            // add lowest to new slice
		sortMe = popFromSlice(sortMe, lowNumIndRec)     // delete lowest from initial slice
	}

	fmt.Printf("Sorted slice: %T %v\n", sortYes, sortYes)

}

func findLowestNum(sliceToSearch []int) (int, int) {
	// finds the lowest number in slice
	lowNum := sliceToSearch[0]
	lowInd := 0
	for i := 0; i < len(sliceToSearch); i++ {
		if sliceToSearch[i] < lowNum {
			lowNum = sliceToSearch[i]
			lowInd = i
		}
	}
	return lowInd, lowNum
}

func popFromSlice(sliceToPop []int, indToPop int) []int {
	// removes element from slice

	// move slice data with a loop
	//for i := lowestInd + 1; i < len(sortMe); i++ {
	//	sortMe[i-1] = sortMe[i]
	//}

	// move slice data with a copy
	copy(sliceToPop[indToPop:], sliceToPop[indToPop+1:])
	// remove last element from slice
	sliceToPop = sliceToPop[:len(sliceToPop)-1]
	return sliceToPop
}

func createToSortSlice() []int {
	var sliceCreate int
	var sliceSize int
	var toSortSlice []int

	fmt.Println("Enter size of array to sort:")
	fmt.Scan(&sliceSize)
	fmt.Println("Enter numbers pressing Enter after each:")
	for i := 0; i < sliceSize; i++ {
		fmt.Print("Element number ", i+1, ": ")
		_, err := fmt.Scan(&sliceCreate)
		if err == nil {
			toSortSlice = append(toSortSlice, sliceCreate)
		} else {
			fmt.Println("Wrong data. Only integers allowed. Enter numbers pressing Enter after each:")
			i--
		}
	}
	return toSortSlice
}

func writeSliceToFile(sliceToWrite []int) {
	var fileNameUnsorted string
	fmt.Println("Enter file name to save unsorted numbers:")
	fmt.Scan(&fileNameUnsorted)
	z, err := os.Create(fileNameUnsorted + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	defer z.Close()
	for _, value := range sliceToWrite {
		fmt.Fprintln(z, value) // print values to z, one per line
	}
}
