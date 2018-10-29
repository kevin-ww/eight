package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"math/big"
	"merculet.io/support"
	"testing"
	"time"
)

func TestNewCC(t *testing.T) {
	cc := NewCC()
	fmt.Printf("%v \n", cc)
}

func TestInit(t *testing.T) {
	cc := NewCC()
	stub := shim.NewMockStub(cc.name, cc)

	txId := "001"

	stub.MockTransactionStart(txId)
	response := stub.MockInit(``, nil)
	stub.MockTransactionEnd(txId)

	if response.Status != 200 {
		t.Errorf("Expected status  of 200, but it was %d instead.", response.Status)
	} else {
		fmt.Printf("pass")
	}

}

func TestTransfer(t *testing.T) {

	transaction := &Transaction{
		Payload: support.Payload{
			TxId: "tx001",
			TxTs: 0,
			Memo: "test transaction ",
			//CreatedAt: time.Now(),
			//CreatedBy: "admin",
		},
		From:   AsAccount("alias"),
		To:     AsAccount("bob"),
		Symbol: "mvp",
		Amount: big.NewInt(100),
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

func TestBatchTransfer(t *testing.T) {
	t1 := Transaction{
		Payload: support.Payload{
			TxId: "tx001",
			TxTs: 0,
			Memo: "test transaction ",
			//createdAt: time.Now(),
			//createdBy: "admin",
		},
		From:   AsAccount("alias"),
		To:     AsAccount("bob"),
		Symbol: "mvp",
		Amount: big.NewInt(100),
		Ts:     time.Now().Unix(),
	}

	t2 := Transaction{
		Payload: support.Payload{
			TxId: "tx002",
			TxTs: 0,
			Memo: "test transaction ",
			//CreatedAt: time.Now(),
			//CreatedBy: "admin",
		},
		From:   AsAccount("bob"),
		To:     AsAccount("chris"),
		Symbol: "mvp",
		Amount: big.NewInt(33),
		Ts:     time.Now().Unix(),
	}

	var batch [2]Transaction
	batch[0] = t1
	batch[1] = t2

	//bytes, _ := json.Marshal(batch)
	//fmt.Printf("%v \n",string(bytes))

	//x: = []*Transaction{t1,t2}

	cc := NewCC()
	stub := shim.NewMockStub(cc.name, cc)

	bytes, e := support.Marshal(batch)

	if e != nil {
		t.Errorf("unable to marshal the payload %v %v \n ", e, string(bytes))
	}

	args := make([][]byte, 2)
	args[0] = []byte(`BatchTransfer`)
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

func TestBatchTransferAndGetBalance(t *testing.T) {
	t1 := Transaction{
		Payload: support.Payload{
			TxId: "tx001",
			TxTs: 0,
			Memo: "test transaction ",
			//createdAt: time.Now(),
			//createdBy: "admin",
		},
		From:   AsAccount("alias"),
		To:     AsAccount("bob"),
		Symbol: "mvp",
		Amount: big.NewInt(100),
		Ts:     time.Now().Unix(),
	}

	t2 := Transaction{
		Payload: support.Payload{
			TxId: "tx002",
			TxTs: 0,
			Memo: "test transaction ",
			//CreatedAt: time.Now(),
			//CreatedBy: "admin",
		},
		From:   AsAccount("bob"),
		To:     AsAccount("chris"),
		Symbol: "mvp",
		Amount: big.NewInt(33),
		Ts:     time.Now().Unix(),
	}

	var batch [2]Transaction
	batch[0] = t1
	batch[1] = t2

	//bytes, _ := json.Marshal(batch)
	//fmt.Printf("%v \n",string(bytes))

	//x: = []*Transaction{t1,t2}

	cc := NewCC()
	stub := shim.NewMockStub(cc.name, cc)

	bytes, e := support.Marshal(batch)

	if e != nil {
		t.Errorf("unable to marshal the payload %v %v \n ", e, string(bytes))
	}

	args := make([][]byte, 2)
	args[0] = []byte(`BatchTransfer`)
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

	//TODO
	account := &Account{
		Name:   "bob",
		Symbol: "mvp",
	}
	bytes, err := support.Marshal(account)
	if err != nil {
		t.Errorf("unable to marshal the payload %v %v \n ", e, string(bytes))
	}
	args[0] = []byte(`GetBalance`)
	args[1] = bytes

	txId = "002"
	stub.MockTransactionStart(txId)
	response = stub.MockInvoke(txId, args)
	stub.MockTransactionEnd(txId)

	if response.Status != 200 {
		t.Errorf("Expected status  of 200, but it was %d instead.", response.Status)
		fmt.Printf("msg : %v  \n", response.GetMessage())
		fmt.Printf("payload: %v  \n", string(response.GetPayload()))

	} else {
		fmt.Printf("payload status: %v  \n", response.GetStatus())
		fmt.Printf("msg : %v  \n", response.GetMessage())
		fmt.Printf("payload: %v  \n", string(response.GetPayload()))
	}

}

func TestUnMarshal(t *testing.T) {

	t1 := Transaction{
		Payload: support.Payload{
			TxId: "tx001",
			TxTs: 0,
			Memo: "test transaction ",
			//createdAt: time.Now(),
			//createdBy: "admin",
		},
		From:   AsAccount("alias"),
		To:     AsAccount("bob"),
		Symbol: "mvp",
		Amount: big.NewInt(100),
		Ts:     time.Now().Unix(),
	}

	t2 := Transaction{
		Payload: support.Payload{
			TxId: "tx002",
			TxTs: 0,
			Memo: "test transaction ",
			//CreatedAt: time.Now(),
			//CreatedBy: "admin",
		},
		From:   AsAccount("bob"),
		To:     AsAccount("chris"),
		Symbol: "mvp",
		Amount: big.NewInt(33),
		Ts:     time.Now().Unix(),
	}

	var batch [2]Transaction
	batch[0] = t1
	batch[1] = t2

	bytes, _ := support.Marshal(batch)

	transactions := make([]Transaction, 0)
	x, e := support.UnMarshal(bytes, &transactions)
	if e != nil {
		fmt.Printf("%v \n", e)
	}
	fmt.Printf("%#v \n", transactions)
	fmt.Printf("%#v \n", x)
}
