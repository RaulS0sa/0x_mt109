package main

import (
	"fmt"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestMatrix1(t *testing.T) {
	fmt.Println("running test_1.")
	var InputMatrix = [][]int{
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1}}
	var OutputMatrix = [][]int{
		{1, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1}}
	resultingMAtrix := PixelRemover(InputMatrix)
	for i, subArray := range resultingMAtrix {
		for j, _ := range subArray {
			if InputMatrix[i][j] != OutputMatrix[i][j] {
				t.Fatalf("Resulting Matrix mismatch")
			}
		}
	}

}
func TestMatrix2(t *testing.T) {
	fmt.Println("running test_2.")
	var InputMatrix = [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1}}
	var OutputMatrix = [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1}}
	resultingMAtrix := PixelRemover(InputMatrix)
	for i, subArray := range resultingMAtrix {
		for j, _ := range subArray {
			if InputMatrix[i][j] != OutputMatrix[i][j] {
				t.Fatalf("Resulting Matrix mismatch")
			}
		}
	}

}
func TestMatrix3(t *testing.T) {
	fmt.Println("running test_3.")
	var InputMatrix = [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1}}
	var OutputMatrix = [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1}}
	resultingMAtrix := PixelRemover(InputMatrix)
	for i, subArray := range resultingMAtrix {
		for j, _ := range subArray {
			if InputMatrix[i][j] != OutputMatrix[i][j] {
				t.Fatalf("Resulting Matrix mismatch")
			}
		}
	}

}

func TestMatrix4(t *testing.T) {
	fmt.Println("running test_4.")
	var InputMatrix = [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 1}}
	var OutputMatrix = [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 1}}
	resultingMAtrix := PixelRemover(InputMatrix)
	for i, subArray := range resultingMAtrix {
		for j, _ := range subArray {
			if InputMatrix[i][j] != OutputMatrix[i][j] {
				t.Fatalf("Resulting Matrix mismatch")
			}
		}
	}

}
