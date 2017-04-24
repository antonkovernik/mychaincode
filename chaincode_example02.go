package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("CLDChaincode")

type  SimpleChaincode struct {
}


//==============================================================================================================================
//	Init Function - Called when the user deploys the chaincode
//==============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	return nil, nil
}

//==============================================================================================================================
//	 Router Functions
//==============================================================================================================================
//	Invoke - Called on chaincode invoke. Takes a function name passed and calls that function. Converts some
//		  initial arguments passed to other things for use in the called function e.g. name -> ecert
//==============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "ping" {
      return t.ping(stub)
  } else { 																				// If the function is not a create then there must be a car so we need to retrieve the car.
			return nil, errors.New("Function of the name "+ function +" doesn't exist.")
	}
}

func (t *SimpleChaincode) get_username(stub shim.ChaincodeStubInterface) (string, error) {

  username, err := stub.ReadCertAttribute("username");
	if err != nil { return "", errors.New("Couldn't get attribute 'username'. Error: " + err.Error()) }
	return string(username), nil
}

func (t *SimpleChaincode) ping(stub shim.ChaincodeStubInterface) ([]byte, error) {
	username, error := t.get_username(stub)
	fmt.Println("Hello " + username)
	if (error) {
		fmt.Println("Error: " + error)
	}
	return []byte("Hello, world!"), error
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

  logger.Debug("function: ", function)

	return nil, errors.New("Received unknown function invocation " + function)
}


func main() {

	err := shim.Start(new(SimpleChaincode))

															if err != nil { fmt.Printf("Error starting Chaincode: %s", err) }
}
