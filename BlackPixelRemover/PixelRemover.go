package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		Method := os.Args[1]
		switch Method {
		case "test_1":
			fmt.Println("running test_1.")
			PixelRemover(test1())
		case "test_2":
			fmt.Println("running test_2.")
			PixelRemover(test2())
		case "test_3":
			fmt.Println("running test_3.")
			PixelRemover(test3())
		case "test_4":
			fmt.Println("running test_4.")
			PixelRemover(test4())
		default:
			fmt.Println("Nothing to test.")
		}
	} else {
		fmt.Println("Nothing to test.")
	}

}
func PixelRemover(matrix [][]int) [][]int {
	fmt.Println("Original Matrix.")
	for _, element := range matrix {
		fmt.Println(arrayToString(element, " "))
	}
	var VisitedCells [][]bool
	var xMaxSize = len(matrix)
	for _, element := range matrix {
		var tempInitArray []bool
		for i := 0; i < len(element); i++ {
			tempInitArray = append(tempInitArray, false)
			//fmt.Println(i)
			//tempInitArray.append(false)
		}

		VisitedCells = append(VisitedCells, tempInitArray)
	}
	for index, element := range matrix {
		for secondIndex, SecondElement := range element {
			if SecondElement == 1 && VisitedCells[index][secondIndex] == false {
				if index == 0 || index == xMaxSize-1 {
					ExploreAdjacentCells(index, secondIndex, VisitedCells, matrix, len(matrix), len(element))
				} else if secondIndex == 0 || secondIndex == (len(element)-1) {
					ExploreAdjacentCells(index, secondIndex, VisitedCells, matrix, len(matrix), len(element))
				}
			}
		}
	}
	for index, element := range matrix {
		for secondIndex, _ := range element {
			if VisitedCells[index][secondIndex] == false {
				matrix[index][secondIndex] = 0
			}
		}
	}
	fmt.Println("Result Matrix.")
	for _, element := range matrix {
		fmt.Println(arrayToString(element, " "))
	}
	return matrix
}
func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func ExploreAdjacentCells(x int, y int, VisitedCells [][]bool, matrix [][]int, xMaxSize int, yMaxSize int) {
	if x >= 0 && x < xMaxSize && y >= 0 && y < yMaxSize {
		if VisitedCells[x][y] == false && matrix[x][y] == 1 {
			VisitedCells[x][y] = true
			ExploreAdjacentCells(x+1, y, VisitedCells, matrix, xMaxSize, yMaxSize)
			ExploreAdjacentCells(x-1, y, VisitedCells, matrix, xMaxSize, yMaxSize)
			ExploreAdjacentCells(x, y+1, VisitedCells, matrix, xMaxSize, yMaxSize)
			ExploreAdjacentCells(x, y-1, VisitedCells, matrix, xMaxSize, yMaxSize)
		}
	}
}

func test1() [][]int {
	var a = [][]int{
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1}}
	return a
}
func test2() [][]int {
	var a = [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1}}
	return a
}
func test3() [][]int {
	var a = [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1}}
	return a
}
func test4() [][]int {
	var a = [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 1}}
	return a
}

/*Output
[1, 0, 0, 0, 0, 0]
[0, 0, 0, 1, 1, 1]
[0, 0, 0, 1, 1, 0]
[0, 0, 0, 0, 1, 0]
[0, 0, 0, 0, 0, 0]
[0, 0, 0, 0, 0, 1]*/
