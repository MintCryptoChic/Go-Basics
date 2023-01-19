package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Init initializes the chaincode
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("nft Init")

	//
	// Demonstrate the use of Attribute-Based Access Control (ABAC) by checking
	// to see if the caller has the "nft.init" attribute with a value of true;
	// if not, return an error.
	//
	err := cid.AssertAttributeValue(stub, "nft.init", "true")
	if err != nil {
		return shim.Error(err.Error())
	}

	_, args := stub.GetFunctionAndParameters()
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (p *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Rezoomex nft chaincode Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "create_nft" {
		return p.createNFT(stub, args)
	} else if function == "get_nft" {
		return p.getNFT(stub, args)
	} else if function == "getHistoryFornft" { //get history of values for a marble
		return p.getHistoryForNFT(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting  \"create_nft\" \"get_nft\" \"getHistoryFornft\"")
}

func (t *SimpleChaincode) createNFT(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("=== create nft ===")
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting key-value pair")
	}

	metamaskId := args[0]
	resumeDetails := args[1]

	resumeJSONasBytes, err := json.Marshal(resumeDetails)
	err = stub.PutState(metamaskId, resumeJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *SimpleChaincode) getNFT(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("=== get NFT ===")

	var metamaskId string
	var err error
	var jsonResp string

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key-value pair")
	}

	metamaskId = args[0]

	resumeBytes, err := stub.GetState(metamaskId)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get data for " + metamaskId + "\"}"
		return shim.Error(jsonResp)
	}

	if resumeBytes == nil {
		jsonResp = "{\"Error\":\"No data for " + metamaskId + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(resumeBytes)

}

func (t *SimpleChaincode) getHistoryForNFT(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	metamaskId := args[0]

	fmt.Printf("- start getHistoryForNFT: %s\n", metamaskId)

	resultsIterator, err := stub.GetHistoryForKey(metamaskId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForNFT returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
