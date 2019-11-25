package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type indexValuePair struct {
	row int
	col int
	value float64
}

func main() {
	ETC := ETCgenerator(100, 10, "low", "low")
	ETC1 := deepcopy(ETC)
	ETC2 := deepcopy(ETC)
	ETC3 := deepcopy(ETC)
	ETC4 := deepcopy(ETC)
	ETC5 := deepcopy(ETC)

	fmt.Printf("%.2f\n\n",ETC)

	var _, cost = assignTask(ETC,"MIN-MIN-RESOURCE")
	fmt.Printf("%.2f\n\n", cost)

	 _, cost = assignTask(ETC1,"MAX-MIN-RESOURCE")
	fmt.Printf("%.2f\n\n", cost)

	_, cost = assignTask(ETC2,"MIN-MAX-RESOURCE")
	fmt.Printf("%.2f\n\n", cost)

	 _, cost = assignTask(ETC3,"MIN-MIN-TASK")
	fmt.Printf("%.2f\n\n", cost)

	 _, cost = assignTask(ETC4,"MAX-MIN-TASK")
	fmt.Printf("%.2f\n\n", cost)

	 _, cost = assignTask(ETC5,"MIN-MAX-TASK")
	fmt.Printf("%.2f\n\n", cost)


}


func deepcopy(inputMatrix [][]float64) [][]float64 {
	result := make([][]float64, len(inputMatrix))
	for i := range result {
		result[i] = make([]float64, len(inputMatrix[i]))
	}
	for i := 0; i < len(inputMatrix); i++ {
		for j := 0; j < len(inputMatrix[i]); j++ {
			result[i][j] = inputMatrix[i][j]
		}
	}
	return result
}

// ETCgenerator : generate an ETC matrix based on # tasks, resources, heterogenety of task and resource
func ETCgenerator(task int, resource int, taskHetero string, resourceHetero string) [][]float64 {
	result := make([][]float64, task)
	for i := range result {
		result[i] = make([]float64, resource)
	}
	var taskBound float64
	var resourceBound float64

	if taskHetero == "hi" {
		taskBound = 3000
	} else {
		taskBound = 100
	}

	if resourceHetero == "hi" {
		resourceBound = 1000
	} else {
		resourceBound = 10
	}

	for i := range result {
		result[i][0] = rand.Float64()*(taskBound-1.0) + 1.0
	}

	start := 1

	for i := 0; i < task; i++ {
		start = 1
		for j := 0; j < resource; j++ {
			if j == (resource - 1) {
				start = resource - 1
			}
			result[i][start] = result[i][0] * (rand.Float64()*(resourceBound-1.0) + 1.0)
			start++
		}
	}

	for i := 0; i < task; i++ {
		result[i][0] = result[i][0] * (rand.Float64()*(resourceBound-1.0) + 1.0)
	}

	return result
}


func assignTask(inputMatrix [][]float64, input string) ([]indexValuePair, float64) {
	var emptyArr []indexValuePair
	var emptyArr1 []float64
	tempMatrix := inputMatrix
	choices, timespent := helper(tempMatrix, emptyArr, emptyArr1, input) // choices contains the row and col number of the original matrix
	var timeCost float64
	timeCost = -1
	fmt.Println(timespent)
	for i := 0; i < len(timespent); i++ {
		if timespent[i] > timeCost {
			timeCost = timespent[i]
		}
	}
	return choices, timeCost
}

func helper(inputMatrix [][]float64, result []indexValuePair, timespent []float64, input string) ([]indexValuePair, []float64) {


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

func getminIndicesByRow(inputMatrix [][]float64) []indexValuePair {
	result := make([]indexValuePair, len(inputMatrix))
	for i := 0; i < len(inputMatrix); i++ {
		var minofRow float64 = math.MaxFloat64
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

func getminIndicesByCol(inputMatrix [][]float64) []indexValuePair {

	result := make([]indexValuePair, len(inputMatrix[0]))
	for j := 0; j < len(inputMatrix[0]); j++ {
		var minofCol float64 = math.MaxFloat64
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

func getmaxIndicesByRow(inputMatrix [][]float64) []indexValuePair {
	result := make([]indexValuePair, len(inputMatrix))
	for i := 0; i < len(inputMatrix); i++ {
		var maxofRow float64 = -math.MaxFloat64
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

func getmaxIndicesByCol(inputMatrix [][]float64) []indexValuePair {
	result := make([]indexValuePair, len(inputMatrix[0]))
	for j := 0; j < len(inputMatrix[0]); j++ {
		var maxofCol float64 = -math.MaxFloat64
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
	var max float64 = -math.MaxFloat64
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
	var min float64 = math.MaxFloat64
	var minPair indexValuePair
	for i := 0; i < len(inputMatrix); i++ {
		if inputMatrix[i].value < min {
			min = inputMatrix[i].value
			minPair = inputMatrix[i]
		}
	}
	return minPair
}

func shrinkMatrixRow(inputMatrix [][]float64, rowRemoved int) [][]float64 {
	result := make([][]float64, len(inputMatrix)-1)
	for c := range result {
		result[c] = make([]float64, len(inputMatrix[c]))
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
