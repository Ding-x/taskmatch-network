package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type indexValuePair struct {
	row int
	col int
	value float64
}

type TaskMatching struct {
	identifier string `json:"id"`       //docType is used to distinguish the various types of objects in state database
	Runtimes   [][]float64 `json:"runtimes"` //the fieldtags are needed to keep case from bouncing around
	VarMax int `json:"varmax"`
	VarMin int `json:"varmin"`
}

type Peer struct {
	identifier string `json:"id"`
	Status     string `json:"status"`
	CompletionTime    float64    `json:"CompletionTime"`
	Name       string `json:"name"`
}

type TaskMatchingSol struct {
	identifier string `json:"id"`
	CompletionTime    float64    `json:"CompletionTime"`
	Owner      string `json:"owner"`
	Algorithm  string `json:"alg"`
	Runtimes   [][]float64 `json:"runtimes"`
}

type Count struct {
	identifier string `json:"id"`
	Counter    int    `json:"count"`
}



// Position : this contains the currrent position and the point's fitness value
type Position struct {
	position []float64
	cost     float64
}

// Problem : defines the structure of a problem, including the number of tasks and
// resources
type Problem struct {
	nVar   int // number of tasks
	varmin int
	varMax int
}

// Particle : this is the particle struct
type Particle struct {
	position []float64
	velocity []float64
	pBest    []float64
	cost     float64
	bestCost float64
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "createTaskMatching" { //create a new taskmatching
		return t.createTaskMatching(stub, args)
	} else if function == "readTaskMatching" { //reads a taskmatching
		return t.readTaskMatching(stub, args)
	} else if function == "Initialize" { //initialize the network
		return t.Initialize(stub)
	} else if function == "calculateTaskMatching" { //calculate a taskmatching
		t.calculateTaskMatching(stub, args)

		if t.allPeersDone(stub) {
			return t.setBestSol(stub)
		} else {
			return shim.Success(nil)
		}
	}
	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) Initialize(stub shim.ChaincodeStubInterface) pb.Response {
	var err error

	p1 := &Peer{"p1", "waiting", -1, "Peer 1"}
	p1JSONasBytes, _ := json.Marshal(p1)

	err = stub.PutState("p1", p1JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p2 := &Peer{"p2", "waiting", -1, "Peer 2"}
	p2JSONasBytes, _ := json.Marshal(p2)

	err = stub.PutState("p2", p2JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p3 := &Peer{"p3", "waiting", -1, "Peer 3"}
	p3JSONasBytes, _ := json.Marshal(p3)

	err = stub.PutState("p3", p3JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p4 := &Peer{"p4", "waiting", -1, "Peer 4"}
	p4JSONasBytes, _ := json.Marshal(p4)

	err = stub.PutState("p4", p4JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p5 := &Peer{"p5", "waiting", -1, "Peer 5"}
	p5JSONasBytes, _ := json.Marshal(p5)

	err = stub.PutState("p5", p5JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p6 := &Peer{"p6", "waiting", -1, "Peer 6"}
	p6JSONasBytes, _ := json.Marshal(p6)

	err = stub.PutState("p6", p6JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p7 := &Peer{"p7", "waiting", -1, "Peer 7"}
	p7JSONasBytes, _ := json.Marshal(p7)

	err = stub.PutState("p7", p7JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}	

	count := &Count{"count", 0}
	countAsBytes, _ := json.Marshal(count)

	stub.PutState("count", countAsBytes)

	return shim.Success(nil)
}

func (t *SimpleChaincode) calculateTaskMatching(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string

	//Get the Task Math matrix
	TaskMatchAsBytes, _ := stub.GetState("work")
	tmpTM := TaskMatching{}

	json.Unmarshal(TaskMatchAsBytes, &tmpTM)

	//Convert matrix string to float matrix
	matrix := tmpTM.Runtimes

	//pass matrix to solution calculator
	var runtime float64

	runtime = Assign(matrix, args[0], tmpTM.VarMax, tmpTM.VarMin)

	//change Peer info
	PeerasBytes, err := stub.GetState(args[0])

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + args[0] + "\"}"
		return shim.Error(jsonResp)
	} else if PeerasBytes == nil {
		jsonResp = "{\"Error\":\"Peer does not exist: " + args[0] + "\"}"
		return shim.Error(jsonResp)
	}

	tmpPeer := Peer{}

	json.Unmarshal(PeerasBytes, &tmpPeer)

	tmpPeer.Status = "done"
	tmpPeer.CompletionTime = runtime

	PeerAsJSONbytes, _ := json.Marshal(tmpPeer)

	stub.PutState(args[0], PeerAsJSONbytes)

	return shim.Success(nil)
	//
}

func strToMatrix(input string) [][]float64 {
	var parsed [][]float64
	json.Unmarshal([]byte(input), &parsed)
	return parsed
}

func Assign(matrix [][]float64, peer string, varMax int, varMin int) float64 {
	rand.Seed(time.Now().UnixNano())
	var timeCost float64
	timeCost = -1
	if peer == "p1" {
		timeCost = assignTask(matrix,"MIN-MIN-TASK")
	} else if peer == "p2" {
		timeCost = assignTask(matrix,"MIN-MAX-TASK")
	} else if peer == "p3" {
		timeCost = assignTask(matrix,"MAX-MIN-TASK")
	} else if peer == "p4" {
		timeCost = assignTask(matrix,"MIN-MAX-RESOURCE")
	} else if peer == "p5" {
	 	timeCost = assignTask(matrix,"MAX-MIN-RESOURCE")
	} else if peer == "p6" {
		var newproblem = Problem{varMax, 0,varMin} 
		gbest, _  := pso(newproblem, matrix, 500, 100, 1.796180, 1.796180, 0.729844, 0.995)
		timeCost = gbest.cost
	}

	
	return  timeCost
}


func (t *SimpleChaincode) allPeersDone(stub shim.ChaincodeStubInterface) bool {
	peerArray := [6]string{"p1", "p2", "p3", "p4", "p5", "p6"}
	tmpPeer := Peer{}

	//loop over all of the peers
	for i := 0; i < len(peerArray); i++ {
		//check to see if any of the peers haven't finished

		//query chaincode to get the result
		PeerasBytes, _ := stub.GetState(peerArray[i])
		json.Unmarshal(PeerasBytes, &tmpPeer)

		if tmpPeer.Status != "done" {
			return false
		}
	}

	return true
}

//Method to set the best solution
func (t *SimpleChaincode) setBestSol(stub shim.ChaincodeStubInterface) pb.Response {
	peerArray := [6]string{"p1", "p2", "p3", "p4", "p5", "p6"}
	tmpPeer := Peer{}
	solPeer := Peer{}
	// var min float64 = math.MaxFloat64
	var min float64 = math.MaxFloat64

	//find which peer found the best solution and save their information
	for i := 0; i < len(peerArray); i++ {
		PeerasBytes, _ := stub.GetState(peerArray[i])
		json.Unmarshal(PeerasBytes, &tmpPeer)

		if tmpPeer.CompletionTime < min {
			min = tmpPeer.CompletionTime
			solPeer = tmpPeer
		}
	}

	//get the current matrix we were working on from the ledger
	taskMatchingAsBytes, _ := stub.GetState("work")
	tmpTM := TaskMatching{}

	json.Unmarshal(taskMatchingAsBytes, &tmpTM)

	//get the current count for how many solutions have been created.
	countAsBytes, _ := stub.GetState("count")
	tmpCount := Count{}

	json.Unmarshal(countAsBytes, &tmpCount)

	tmpCount.Counter += 1
	solNum := strconv.Itoa(tmpCount.Counter)

	var algName string

	//find which algorithm was used to calculate the solution.
	if solPeer.Name == "Peer 1" {
		algName = "MIN-MIN-TASK"
	} else if solPeer.Name == "Peer 2" {
		algName = "MIN-MAX-TASK"
	} else if solPeer.Name == "Peer 3" {
		algName = "MAX-MIN-TASK"
	} else if solPeer.Name == "Peer 4" {
		algName = "MIN-MAX-RESOURCE"
	} else if solPeer.Name == "Peer 5" {
		algName = "MAX-MIN-RESOURCE"
	} else if solPeer.Name == "Peer 6" {
		algName = "PSO"
	} 

	TMSol := TaskMatchingSol{solNum, solPeer.CompletionTime, solPeer.Name, algName, tmpTM.Runtimes}

	//update count and add TM sol
	countAsJSON, _ := json.Marshal(tmpCount)
	stub.PutState("count", countAsJSON)

	TMSolAsJSON, _ := json.Marshal(TMSol)
	stub.PutState(solNum, TMSolAsJSON)

	return shim.Success(nil)
}

// ============================================================
// createTaskMatching - create a taskmatching
// ============================================================
func (t *SimpleChaincode) createTaskMatching(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0       1
	//id   runtimes
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	fmt.Println("- creating TaskMatching")

	identifier := args[0]

	fmt.Println()

	arr := strings.Split(args[1], "|")

	row, _ := strconv.Atoi(arr[0])
	col, _ := strconv.Atoi(arr[1])

	runtimes := ETCgenerator(row, col, arr[2], arr[3])

	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}

	// ==== Check if TaskMatching already exists ====
	TaskMatchingAsBytes, err := stub.GetState(identifier)
	if err != nil {
		return shim.Error("Failed to get TaskMatching: " + err.Error())
	} else if TaskMatchingAsBytes != nil {
		fmt.Println("This TaskMatching already exists: " + identifier)
		return shim.Error("This TaskMatching already exists: " + identifier)
	}

	// ==== Create TaskMatching object and marshal to JSON ====
	TaskMatching := &TaskMatching{identifier, runtimes, row, col}
	TaskMatchingJSONasBytes, err := json.Marshal(TaskMatching)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save taskmatching to state ===
	err = stub.PutState(identifier, TaskMatchingJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== taskmathing saved. Return success ====
	fmt.Println("- end init TaskMatching")
	return shim.Success(nil)
}

// ================================================================================================================
// readTaskMatching: This method can actually read anything that is saved onto the ledger not just taskmatchings
// ================================================================================================================
func (t *SimpleChaincode) readTaskMatching(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var identifier, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the TaskMatching to query")
	}

	identifier = args[0]
	TaskMatchingAsbytes, err := stub.GetState(identifier) //get the TaskMatching from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + identifier + "\"}"
		return shim.Error(jsonResp)
	} else if TaskMatchingAsbytes == nil {
		jsonResp = "{\"Error\":\"TaskMatching does not exist: " + identifier + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(TaskMatchingAsbytes)
}

func Decimal(value float64) float64{
	value,_ = strconv.ParseFloat( fmt.Sprintf("%.2f", value), 64)
	return value
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
		result[i][0] = Decimal(rand.Float64()*(taskBound-1.0) + 1.0)
	}

	start := 1

	for i := 0; i < task; i++ {
		start = 1
		for j := 0; j < resource; j++ {
			if j == (resource - 1) {
				start = resource - 1
			}
			result[i][start] = Decimal(result[i][0] * (rand.Float64()*(resourceBound-1.0) + 1.0))
			start++
		}
	}

	for i := 0; i < task; i++ {
		result[i][0] = Decimal(result[i][0] * (rand.Float64()*(resourceBound-1.0) + 1.0))
	}

	return result
}



func assignTask(inputMatrix [][]float64, input string) float64{
	var emptyArr []indexValuePair
	var emptyArr1 []float64
	tempMatrix := inputMatrix
	timespent := helper(tempMatrix, emptyArr, emptyArr1, input) // choices contains the row and col number of the original matrix
	var timeCost float64
	timeCost = -1
	fmt.Println(timespent)
	for i := 0; i < len(timespent); i++ {
		if timespent[i] > timeCost {
			timeCost = timespent[i]
		}
	}
	return timeCost
}

func helper(inputMatrix [][]float64, result []indexValuePair, timespent []float64, input string) []float64 {
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
		return timespent
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



// PSO

func multiplyNumAndArr(factor float64, arrIn []float64) []float64 {
	result := make([]float64, len(arrIn))
	for i := 0; i < len(arrIn); i++ {
		result[i] = arrIn[i] * factor
	}
	return result
}

func trimPosition(inputVec []float64, lower int, upper int) {
	for i := 0; i < len(inputVec); i++ {
		inputVec[i] = math.Max(inputVec[i], float64(lower))
		inputVec[i] = math.Min(inputVec[i], float64(upper-1))
	}
}

func addArrs(arrs ...[]float64) []float64 {
	result := make([]float64, len(arrs[0]))
	for _, arr := range arrs {
		for i := 0; i < len(result); i++ {
			result[i] += arr[i]
		}
	}
	return result
}

func multiplyArrs(arr1 []float64, arr2 []float64) []float64 {
	result := make([]float64, len(arr1))
	for i := 0; i < len(arr1); i++ {
		result[i] = arr1[i] * arr2[i]
	}
	return result
}

func subtractArrs(arr1 []float64, arr2 []float64) []float64 {
	result := make([]float64, len(arr1))
	for i := 0; i < len(arr1); i++ {
		result[i] = arr1[i] - arr2[i]
	}
	return result
}

func generateRandomArr(lower float64, upper float64, size int) []float64 {
	result := make([]float64, size)
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		result[i] = rand.Float64()*(upper-lower) + lower
	}
	return result
}

func fetchRunTime(inputMatrix [][]float64, task int, resource int) float64 {
	return inputMatrix[task][resource]
}

func evaluate(inputMatrix [][]float64, inputSol []int) float64 {
	makespan := make([]float64, len(inputMatrix[0]))
	resources := len(inputMatrix[0])
	for i := 0; i < len(inputSol); i++ { // length of input solution = # of tasks
		temp := inputSol[i] // temp => corresponding resource assigned to each task
		if temp > resources || temp < 0 {
			temp = temp % (resources - 1)
		}
		result := fetchRunTime(inputMatrix, i, temp)
		makespan[temp] = makespan[temp] + result
	}
	maxCompletion := float64(-1)
	for j := 0; j < len(makespan); j++ {
		if makespan[j] > maxCompletion {
			maxCompletion = makespan[j]
		}
	}
	return maxCompletion
}


func pso(inputProblem Problem, inputMatrix [][]float64, maxIter int, popSize int, c1 float64, c2 float64, w float64, wdamp float64) (Position, []Particle) {
	// Initialize an empty object of type "Particle"
	var emptyParticle Particle

	// Extract problem information
	varMin := inputProblem.varmin
	varMax := inputProblem.varMax
	nVar := inputProblem.nVar

	gBest := Position{nil, math.Inf(1)}

	pop := []Particle{}

	// This loop is for initialization
	for i := 0; i < popSize; i++ {
		pop = append(pop, emptyParticle)
		pop[i].position = generateRandomArr(float64(varMin), float64(varMax), nVar)
		pop[i].velocity = generateRandomArr(float64(-varMax), float64(varMax), nVar)
		x := make([]int, len(pop[i].position))
		for j := 0; j < len(x); j++ {
			x[j] = int(pop[i].position[j])
		}
		pop[i].cost = evaluate(inputMatrix, x)
		// copy(pop[i].pBest, pop[i].position)
		pop[i].pBest = pop[i].position
		pop[i].bestCost = pop[i].cost

		if pop[i].bestCost < gBest.cost {
			// copy(gBest.position, pop[i].pBest)
			gBest.position = pop[i].pBest
			gBest.cost = pop[i].bestCost
		}
		// fmt.Println(pop[i].velocity)
	}
	//PSO loop
	for iter := 0; iter < maxIter; iter++ {
		for i := 0; i < popSize; i++ {
			pop[i].velocity = addArrs(multiplyNumAndArr(w, pop[i].velocity),
				multiplyArrs(multiplyNumAndArr(c1, generateRandomArr(0, 1, nVar)), subtractArrs(pop[i].pBest, pop[i].position)),
				multiplyArrs(multiplyNumAndArr(c2, generateRandomArr(0, 1, nVar)), subtractArrs(gBest.position, pop[i].position)))

			pop[i].position = addArrs(pop[i].position, pop[i].velocity)
			trimPosition(pop[i].position, varMin, varMax)

			x := make([]int, len(pop[i].position))
			for j := 0; j < len(x); j++ {
				x[j] = int(pop[i].position[j])
			}

			pop[i].cost = evaluate(inputMatrix, x)
			if pop[i].cost < pop[i].bestCost {
				// copy(pop[i].pBest, pop[i].position)
				pop[i].pBest = pop[i].position
				pop[i].bestCost = pop[i].cost
				if pop[i].bestCost < gBest.cost {
					// copy(gBest.position, pop[i].pBest)
					gBest.position = pop[i].pBest
					gBest.cost = pop[i].bestCost
				}
			}
			// fmt.Printf("%s%f\n", "The current position is:", pop[i].position)
		}
		w *= wdamp
		// fmt.Printf("%s%d%s%f%s%f\n", "Iteration: ", iter, " Best Cost: ", gBest.cost, " ,the position chosen is:", gBest.position)

	}
	return gBest, pop
}

