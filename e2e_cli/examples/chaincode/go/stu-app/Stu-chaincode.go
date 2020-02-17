
package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Stu struct {
	Name   string `json:"name"`
	PID  string `json:"pid"`
	Type string `json:"type"`
	Time  string `json:"time"`
	Score  string `json:"score"`
}

//The Init method is called when the Smart Contract  is instantiated by the blockchain network

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

//The Invoke method is called as a result of an application request to run the Smart Contract

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryStu" {
		return s.queryStu(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordStu" {
		return s.recordStu(APIstub, args)
	} else if function == "queryAllStu" {
		return s.queryAllStu(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

//The queryStu method

func (s *SmartContract) queryStu(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	stuAsBytes, _ := APIstub.GetState(args[0])
	if stuAsBytes == nil {
		return shim.Error("Could not locate Student")
	}
	return shim.Success(stuAsBytes)
}

//The initLedger method

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	stu := []Stu{
		Stu{Name: "lilingxiao", PID: "41411033", Type: "CET4", Time: "2016.12", Score: "550"},
		Stu{Name: "Peter", PID: "12341234", Type: "CET4", Time: "2016.12", Score: "480"},
		Stu{Name: "Bob", PID: "12345678", Type: "CET6", Time: "2016.12", Score: "530"},
		Stu{Name: "Alice", PID: "87654321", Type: "CET6", Time: "2016.12", Score: "600"},
	}

	i := 0
	for i < len(stu) {
		fmt.Println("i is ", i)
		stuAsBytes, _ := json.Marshal(stu[i])
		APIstub.PutState("Stu"+strconv.Itoa(i), stuAsBytes)
		fmt.Println("Added", stu[i])
		i = i + 1
	}

	return shim.Success(nil)
}

//The recordStu method

func (s *SmartContract) recordStu(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	var stu = Stu{Name: args[1], PID: args[2], Type: args[3], Time: args[4], Score: args[5]}

	stuAsBytes, _ := json.Marshal(stu)
	err := APIstub.PutState(args[0], stuAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record Student: %s", args[0]))
	}
	return shim.Success(nil)
}

//The queryAllStu method

func (s *SmartContract) queryAllStu(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "Stu0"
	endKey := "Stu999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllStu:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
