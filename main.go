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
	value float32
}

func main() {
	var twoD = strToMatrix("[[6.3,2.8,3.7],[4.2,9.4,1.8],[7.2,8.3,5.5],[4.7,6.4,8.1]]")
	var choices, cost = assignTask(twoD,"MAX-MIN-RESOURCE")
	for i := 0; i < len(choices); i++ {
		fmt.Printf("%f\n", choices[i].value)
	}
	fmt.Printf("\n")
	fmt.Printf("%f\n\n", cost)
}

func strToMatrix(input string) [][]float32 {
	var parsed [][]float32
	json.Unmarshal([]byte(input), &parsed)
	return parsed
}

func assignTask(inputMatrix [][]float32, input string) ([]indexValuePair, float32) {
	var emptyArr []indexValuePair
	var emptyArr1 []float32
	tempMatrix := inputMatrix
	choices, timespent := helper(tempMatrix, emptyArr, emptyArr1, input) // choices contains the row and col number of the original matrix
	var timeCost float32
	timeCost = -1
	fmt.Println(timespent)
	for i := 0; i < len(timespent); i++ {
		if timespent[i] > timeCost {
			timeCost = timespent[i]
		}
	}
	return choices, timeCost
}

func helper(inputMatrix [][]float32, result []indexValuePair, timespent []float32, input string) ([]indexValuePair, []float32) {
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

func getminIndicesByRow(inputMatrix [][]float32) []indexValuePair {
	result := make([]indexValuePair, len(inputMatrix))
	for i := 0; i < len(inputMatrix); i++ {
		var minofRow float32 = math.MaxInt16
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

func getminIndicesByCol(inputMatrix [][]float32) []indexValuePair {

	result := make([]indexValuePair, len(inputMatrix[0]))
	for j := 0; j < len(inputMatrix[0]); j++ {
		var minofCol float32 = math.MaxInt16
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

func getmaxIndicesByRow(inputMatrix [][]float32) []indexValuePair {
	result := make([]indexValuePair, len(inputMatrix))
	for i := 0; i < len(inputMatrix); i++ {
		var maxofRow float32 = -1.0
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

func getmaxIndicesByCol(inputMatrix [][]float32) []indexValuePair {
	result := make([]indexValuePair, len(inputMatrix[0]))
	for j := 0; j < len(inputMatrix[0]); j++ {
		var maxofCol float32 = -1.0
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
	var max float32 = -1
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
	var min float32 = math.MaxInt16
	var minPair indexValuePair
	for i := 0; i < len(inputMatrix); i++ {
		if inputMatrix[i].value < min {
			min = inputMatrix[i].value
			minPair = inputMatrix[i]
		}
	}
	return minPair
}

func shrinkMatrixRow(inputMatrix [][]float32, rowRemoved int) [][]float32 {
	result := make([][]float32, len(inputMatrix)-1)
	for c := range result {
		result[c] = make([]float32, len(inputMatrix[c]))
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
