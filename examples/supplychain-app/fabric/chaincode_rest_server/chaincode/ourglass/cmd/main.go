package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"

	ourglass "github.com/chaincode/ourglass"
)

// main function starts up the chaincode in the container during instantiate
func main() {
	err := shim.Start(new(ourglass.SmartContract))

	if err != nil {
		fmt.Printf("Error starting ourglass chaincode: %s", err)
	}
}
