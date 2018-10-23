package support

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func HandlerTx(ext Extension, stub shim.ChaincodeStubInterface) (*LedgerStub, error) {

	funcName, args := stub.GetFunctionAndParameters()

	//is function name valid and allowed?

	if !isFuncValid(funcName, ext) {
		return nil, errFunctionNameIsEmptyORNotAllowed
	}

	//is args valid and allowed?
	if args == nil || len(args) == 0 || len(args) != 1 {
		return nil, errPayloadInvalid
	}

	fmt.Printf("invoking [%v] with args: %v \n", funcName, args[0])

	//interface convention

	return &LedgerStub{
		admin:  "admin",
		bucket: "tx",
		stub:   stub,
	}, nil

}
