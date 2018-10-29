package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"merculet.io/support"
	"testing"
)

var (
	acct = Account{
		Payload: &support.Payload{
			TxId:      "",
			TxTs:      0,
			Memo:      "create a test account",
			CreatedAt: 0,
			CreatedBy: "",
		},
		Name:         "kevin",
		Organization: "kevin@nasa",
	}
)

func Test1(t *testing.T) {
	//New
	cc := NewCC()
	stub := shim.NewMockStub(cc.name, cc)

	txId := "001"

	args := make([][]byte, 2)
	args[0] = []byte(`Transfer`)
	args[1] = bytes

	stub.MockTransactionStart(txId)
	response := stub.MockInvoke(``, nil)
	stub.MockTransactionEnd(txId)

	if response.Status != 200 {
		t.Errorf("Expected status  of 200, but it was %d instead.", response.Status)
	} else {
		fmt.Printf("pass")
	}

}
