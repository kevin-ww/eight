package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"merculet.io/support"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	cc := NewCC()
	stub := shim.NewMockStub(cc.name, cc)

	txId := "001"

	stub.MockTransactionStart(txId)
	response := stub.MockInit(``, nil)
	stub.MockTransactionEnd(txId)

	if response.Status != 200 {
		t.Errorf("Expected status  of 200, but it was %d instead.", response.Status)
	}

}

func TestTransfer(t *testing.T) {

	transaction := &Transaction{
		Payload: &support.Payload{
			TxId:      "tx001",
			TxTs:      0,
			Memo:      "test transaction ",
			CreatedAt: time.Now(),
			CreatedBy: "admin",
		},
		From:   "alias",
		To:     "bob",
		Symbol: "mvp",
		Amount: 10,
		Ts:     time.Now().Unix(),
	}

	cc := NewCC()
	stub := shim.NewMockStub(cc.name, cc)

	bytes, e := support.Marshal(transaction)

	if e != nil {
		t.Errorf("unable to marshal the payload %v %v \n ", e, transaction)
	}

	args := make([][]byte, 2)
	args[0] = []byte(`Transfer`)
	args[1] = bytes

	txId := "001"
	stub.MockTransactionStart(txId)
	response := stub.MockInvoke(txId, args)
	stub.MockTransactionEnd(txId)

	if response.Status != 200 {
		t.Errorf("Expected status  of 200, but it was %d instead.", response.Status)
		fmt.Printf("msg : %v  \n", response.GetMessage())
		fmt.Printf("payload: %v  \n", string(response.GetPayload()))

	} else {
		fmt.Printf("payload status: %v  \n", response.GetStatus())
	}

}
