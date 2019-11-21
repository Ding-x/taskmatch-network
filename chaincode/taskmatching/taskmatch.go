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
	value int
}

type TaskMatching struct {
	identifier string `json:"id"`       //docType is used to distinguish the various types of objects in state database
	Runtimes   string `json:"runtimes"` //the fieldtags are needed to keep case from bouncing around
}

type Peer struct {
	identifier string `json:"id"`
	Status     string `json:"status"`
	Solution   []indexValuePair  `json:"sol"`
	Runtime    int    `json:"runtime"`
	Name       string `json:"name"`
}

type TaskMatchingSol struct {
	identifier string `json:"id"`
	Runtime    int    `json:"runtime"`
	Solution   []indexValuePair  `json:"sol"`
	Owner      string `json:"owner"`
	Algorithm  string `json:"alg"`
	Runtimes   string `json:"runtimes"`
}

type Count struct {
	identifier string `json:"id"`
	Counter    int    `json:"count"`
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

	p1 := &Peer{"p1", "waiting", make([]indexValuePair, 0), -1, "Peer 1"}
	p1JSONasBytes, _ := json.Marshal(p1)

	err = stub.PutState("p1", p1JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p2 := &Peer{"p2", "waiting", make([]indexValuePair, 0), -1, "Peer 2"}
	p2JSONasBytes, _ := json.Marshal(p2)

	err = stub.PutState("p2", p2JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p3 := &Peer{"p3", "waiting", make([]indexValuePair, 0), -1, "Peer 3"}
	p3JSONasBytes, _ := json.Marshal(p3)

	err = stub.PutState("p3", p3JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p4 := &Peer{"p4", "waiting", make([]indexValuePair, 0), -1, "Peer 4"}
	p4JSONasBytes, _ := json.Marshal(p4)

	err = stub.PutState("p4", p4JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p5 := &Peer{"p5", "waiting", make([]indexValuePair, 0), -1, "Peer 5"}
	p5JSONasBytes, _ := json.Marshal(p5)

	err = stub.PutState("p5", p5JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	p6 := &Peer{"p6", "waiting", make([]indexValuePair, 0), -1, "Peer 6"}
	p6JSONasBytes, _ := json.Marshal(p6)

	err = stub.PutState("p6", p6JSONasBytes) //write the peer
	if err != nil {
		return shim.Error(err.Error())
	}

	

	count := &Count{"count", 0}
	countAsBytes, _ := json.Marshal(count)

	stub.PutState("count", countAsBytes)

	return shim.Success(nil)
}

func (t *SimpleChaincode) calculateTaskMatching(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//Get the Task Math matrix
	TaskMatchAsBytes, _ := stub.GetState("work")
	tmpTM := TaskMatching{}

	json.Unmarshal(TaskMatchAsBytes, &tmpTM)

	//Convert matrix string to float matrix
	var matrix [][]int = strToMatrix(tmpTM.Runtimes)

	//pass matrix to solution calculator
	var sol []indexValuePair
	var runtime int

	sol, runtime = Assign(matrix, args[0])

	//change Peer info
	PeerasBytes, _ := stub.GetState(args[0])
	tmpPeer := Peer{}

	json.Unmarshal(PeerasBytes, &tmpPeer)

	tmpPeer.Status = "done"
	tmpPeer.Solution = sol
	tmpPeer.Runtime = runtime

	PeerAsJSONbytes, _ := json.Marshal(tmpPeer)

	stub.PutState(args[0], PeerAsJSONbytes)

	return shim.Success(nil)
	//
}

func strToMatrix(input string) [][]int {
	var parsed [][]int
	json.Unmarshal([]byte(input), &parsed)
	return parsed
}

func Assign(matrix [][]int, peer string) ([]indexValuePair, int) {
	var sol []indexValuePair
	rand.Seed(time.Now().UnixNano())
	timeCost := -1
	if peer == "p1" {
		sol, timeCost = assignTask(matrix,"MIN-MIN-TASK")
	} else if peer == "p2" {
		sol, timeCost = assignTask(matrix,"MIN-MAX-TASK")
	} else if peer == "p3" {
		sol, timeCost = assignTask(matrix,"MAX-MIN-TASK")
	} else if peer == "p4" {
		sol, timeCost = assignTask(matrix,"MIN-MIN-RESOURCE")
	} else if peer == "p5" {
		sol, timeCost = assignTask(matrix,"MIN-MAX-RESOURCE")
	} else if peer == "p6" {
	 	sol, timeCost = assignTask(matrix,"MAX-MIN-RESOURCE")
	}

	return sol, timeCost
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
	var min int = math.MaxInt8

	//find which peer found the best solution and save their information
	for i := 0; i < len(peerArray); i++ {
		PeerasBytes, _ := stub.GetState(peerArray[i])
		json.Unmarshal(PeerasBytes, &tmpPeer)

		if tmpPeer.Runtime < min {
			min = tmpPeer.Runtime
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
		algName = "MIN-MIN-RESOURCE"
	} else if solPeer.Name == "Peer 5" {
		algName = "MIN-MAX-RESOURCE"
	} else if solPeer.Name == "Peer 6" {
		algName = "MAX-MIN-RESOURCE"
	} 

	TMSol := TaskMatchingSol{solNum, solPeer.Runtime, solPeer.Solution, solPeer.Name, algName, tmpTM.Runtimes}

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
	runtimes := strings.ToLower(args[1])

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
	TaskMatching := &TaskMatching{identifier, runtimes}
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
