package main

import (
	"fmt"
	"sort"
)

func main() {
	ClarkeWrite()
}
func ClarkeWrite() {
	savingsMatrix := SavingsMatrix(test0())
	limWeight := 30
	var RouteList [][]int
	remainingNodes := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0}
	demand := []int{0, 12, 12, 6, 16, 15, 10, 8}
	for i := 1; i < len(savingsMatrix); i++ {
		if len(remainingNodes) != 0 {
			_, Contains := remainingNodes[i]
			if Contains {
				CurrentRoute := []int{i}
				keys := make([]int, 0)
				for k, _ := range savingsMatrix[i] {
					keys = append(keys, k)
				}

				sort.Slice(keys, func(i, j int) bool {
					return keys[i] > keys[j]
				})
				runningAvailableWeight := demand[i]
				for _, KeyElement := range keys {
					if (runningAvailableWeight + demand[KeyElement]) <= limWeight {
						_, ContainsKey := remainingNodes[KeyElement]
						if ContainsKey {
							CurrentRoute = append(CurrentRoute, KeyElement)
							delete(remainingNodes, KeyElement)
							runningAvailableWeight = runningAvailableWeight + demand[KeyElement]
						}
					}
				}
				RouteList = append(RouteList, CurrentRoute)
			}

		} else {
			break
		}

	}
	for _, Route := range RouteList {
		fmt.Println(Route)
	}

}

func test0() [][]float32 {
	var costMAtrix = [][]float32{
		{4},
		{4, 5.66},
		{2.83, 6.32, 2.83},
		{4, 8, 5.66, 2.83},
		{5, 8.54, 8.06, 5.39, 3},
		{2, 4.47, 6, 4.47, 4.47, 4.12},
		{4.24, 3.16, 7.62, 7.07, 7.62, 7, 3.16}}
	return costMAtrix
}
func SavingsMatrix(distancesMAtrix [][]float32) map[int]map[int]float32 {
	var savingsMatrix = make(map[int]map[int]float32)
	for i := 1; i < len(distancesMAtrix); i++ {
		var tempArr = make(map[int]float32)
		for j := i + 1; j <= len(distancesMAtrix); j++ {
			tempArr[j] = (distancesMAtrix[i-1][0] + distancesMAtrix[j-1][0] - distancesMAtrix[j-1][i])
		}
		savingsMatrix[i] = tempArr
	}
	return savingsMatrix
}

/*Output
[1, 0, 0, 0, 0, 0]
[0, 0, 0, 1, 1, 1]
[0, 0, 0, 1, 1, 0]
[0, 0, 0, 0, 1, 0]
[0, 0, 0, 0, 0, 0]
[0, 0, 0, 0, 0, 1]*/
