package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

type indexValuePair struct {
	row int
	col int
	value int
}

func main() {
	var twoD = strToMatrix("[[6,2,3],[4,9,1],[7,8,5],[4,6,8]]")
	var choices, cost = assignTask(twoD,"MAX-MIN-RESOURCE")
	for i := 0; i < len(choices); i++ {
		fmt.Printf("%d\n", choices[i])
	}
	fmt.Printf("\n")
	fmt.Printf("%d\n\n", cost)
}

func strToMatrix(input string) [][]int {
	var parsed [][]int
	json.Unmarshal([]byte(input), &parsed)
	return parsed
}

func assignTask(inputMatrix [][]int, input string) ([]indexValuePair, int) {
	var emptyArr []indexValuePair
	var emptyArr1 []int
	tempMatrix := inputMatrix
	choices, timespent := helper(tempMatrix, emptyArr, emptyArr1, input) // choices contains the row and col number of the original matrix
	timeCost := -1
	fmt.Println(timespent)
	for i := 0; i < len(timespent); i++ {
		if timespent[i] > timeCost {
			timeCost = timespent[i]
		}
	}
	return choices, timeCost
}

func helper(inputMatrix [][]int, result []indexValuePair, timespent []int, input string) ([]indexValuePair, []int) {
	if len(inputMatrix) == 1 {
		var incides  []indexValuePair
		inputArr := strings.Split(input,"-")

		if inputArr[0]=="MIN" && inputArr[2]=="TASK"{
			incides = getminIndicesByRow(inputMatrix)
		}else if inputArr[0]=="MAX" && inputArr[2]=="TASK"{
			incides = getmaxIndicesByRow(inputMatrix)
		}else if inputArr[0]=="MIN" && inputArr[2]=="RESOURCE"{
			incides = getminIndicesByCol(inputMatrix)
		}else if inputArr[0]=="MAX" && inputArr[2]=="RESOURCE"{
			incides = getmaxIndicesByCol(inputMatrix)
		}

		var valuePair indexValuePair
		if inputArr[1]=="MIN"{
			valuePair = getMinIndexValuePair(incides) 
		}else{
			valuePair = getMaxIndexValuePair(incides) 
		}
	
		result = append(result, valuePair)
		timespent = append(timespent, valuePair.value)
		return result, timespent
	}
	
	var incides []indexValuePair
	inputArr := strings.Split(input,"-")

	if inputArr[0]=="MIN" && inputArr[2]=="TASK"{
		incides = getminIndicesByRow(inputMatrix)
	}else if inputArr[0]=="MAX" && inputArr[2]=="TASK"{
		incides = getmaxIndicesByRow(inputMatrix)
	}else if inputArr[0]=="MIN" && inputArr[2]=="RESOURCE"{
		incides = getminIndicesByCol(inputMatrix)
	}else if inputArr[0]=="MAX" && inputArr[2]=="RESOURCE"{
		incides = getmaxIndicesByCol(inputMatrix)
	}
	

	var valuePair indexValuePair
	if inputArr[1]=="MIN"{
		valuePair = getMinIndexValuePair(incides) 
	}else{
		valuePair = getMaxIndexValuePair(incides) 
	}

	tempMatrix := shrinkMatrixRow(inputMatrix, valuePair.row)
	for i := 0; i < len(tempMatrix); i++ {
		tempMatrix[i][valuePair.col] += valuePair.value
	}
	result = append(result, valuePair)
	timespent = append(timespent, valuePair.value)
	return helper(tempMatrix, result, timespent, input)
}

func getminIndicesByRow(inputMatrix [][]int) []indexValuePair {
	result := make([]indexValuePair, len(inputMatrix))
	for i := 0; i < len(inputMatrix); i++ {
		var minofRow int = math.MaxInt16
		for j := 0; j < len(inputMatrix[i]); j++ {
			if inputMatrix[i][j] < minofRow {
				minofRow = inputMatrix[i][j]
				result[i].row = i
				result[i].col = j
				result[i].value = minofRow
			}
		}
	}
	return result
}

func getminIndicesByCol(inputMatrix [][]int) []indexValuePair {

	result := make([]indexValuePair, len(inputMatrix[0]))
	for j := 0; j < len(inputMatrix[0]); j++ {
		var minofCol int = math.MaxInt16
		for i := 0; i < len(inputMatrix); i++ {
			if inputMatrix[i][j] < minofCol {
				minofCol = inputMatrix[i][j]
				result[j].row = i
				result[j].col = j
				result[j].value = minofCol
			}
		}
	}
	return result
}

func getmaxIndicesByRow(inputMatrix [][]int) []indexValuePair {
	result := make([]indexValuePair, len(inputMatrix))
	for i := 0; i < len(inputMatrix); i++ {
		var maxofRow int = -1
		for j := 0; j < len(inputMatrix[i]); j++ {
			if inputMatrix[i][j] > maxofRow {
				maxofRow = inputMatrix[i][j]
				result[i].row = i
				result[i].col = j
				result[i].value = maxofRow
			}
		}
	}
	return result
}

func getmaxIndicesByCol(inputMatrix [][]int) []indexValuePair {
	result := make([]indexValuePair, len(inputMatrix[0]))
	for j := 0; j < len(inputMatrix[0]); j++ {
		var maxofCol int = -1
		for i := 0; i < len(inputMatrix); i++ {
			if inputMatrix[i][j] > maxofCol {
				maxofCol = inputMatrix[i][j]
				result[j].row = i
				result[j].col = j
				result[j].value = maxofCol
			}
		}
	}
	return result
}

func getMaxIndexValuePair(inputMatrix []indexValuePair) indexValuePair {
	var max int = -1
	var maxPair indexValuePair
	for i := 0; i < len(inputMatrix); i++ {
		if inputMatrix[i].value > max {
			max = inputMatrix[i].value
			maxPair = inputMatrix[i]
		}
	}
	return maxPair
}

func getMinIndexValuePair(inputMatrix []indexValuePair) indexValuePair {
	var min int = math.MaxInt16
	var minPair indexValuePair
	for i := 0; i < len(inputMatrix); i++ {
		if inputMatrix[i].value < min {
			min = inputMatrix[i].value
			minPair = inputMatrix[i]
		}
	}
	return minPair
}

func shrinkMatrixRow(inputMatrix [][]int, rowRemoved int) [][]int {
	result := make([][]int, len(inputMatrix)-1)
	for c := range result {
		result[c] = make([]int, len(inputMatrix[c]))
	}
	if len(inputMatrix) == 1 {
		return inputMatrix
	}

	newRow := 0
	for OriRow := 0; OriRow < len(inputMatrix); OriRow++ {
		if OriRow != rowRemoved {
			result[newRow] = inputMatrix[OriRow]
			newRow++
		}
	}
	return result
}
