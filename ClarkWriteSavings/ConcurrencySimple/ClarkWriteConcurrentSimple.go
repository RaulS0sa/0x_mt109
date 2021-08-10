package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	ClarkeWrite()
}

type entry struct {
	val float32
	key int
}
type entries []entry

func (s entries) Len() int           { return len(s) }
func (s entries) Less(i, j int) bool { return s[i].val < s[j].val }
func (s entries) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func ClarkeWrite() {
	savingsMatrix := SavingsMatrix(test0())
	limWeight := 30
	var RouteList [][]int
	remainingNodes := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0}
	demand := []int{0, 12, 12, 6, 16, 15, 10, 8}
	//for i := 1; i < len(savingsMatrix); i++ {
	for i := 1; i < len(savingsMatrix); i++ {
		if len(remainingNodes) != 0 {
			_, Contains := remainingNodes[i]
			delete(remainingNodes, i)
			if Contains {
				CurrentRoute := []int{i}
				var sortedPoints entries
				for k, v := range savingsMatrix[i] {
					sortedPoints = append(sortedPoints, entry{val: v, key: k})
				}
				sort.Sort(sort.Reverse(sortedPoints))

				runningAvailableWeight := demand[i]
				for _, KeyElement := range sortedPoints {
					//fmt.Println(KeyElement.val)
					loadCapacityContraint := (runningAvailableWeight + demand[KeyElement.key]) <= limWeight
					if loadCapacityContraint {
						_, ContainsKey := remainingNodes[KeyElement.key]
						if ContainsKey {
							CurrentRoute = append(CurrentRoute, KeyElement.key)
							delete(remainingNodes, KeyElement.key)
							runningAvailableWeight = runningAvailableWeight + demand[KeyElement.key]
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
	var wg sync.WaitGroup
	for i := 1; i < len(distancesMAtrix); i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			var tempArr = make(map[int]float32)
			for j := i + 1; j <= len(distancesMAtrix); j++ {
				tempArr[j] = (distancesMAtrix[i-1][0] + distancesMAtrix[j-1][0] - distancesMAtrix[j-1][i])
			}
			savingsMatrix[i] = tempArr
		}(i, &wg)
	}
	wg.Wait()
	return savingsMatrix
}
