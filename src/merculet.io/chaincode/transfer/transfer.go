package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"merculet.io/support"
)

type ChainCode struct {
	name string
	*support.Extension
}

func (cc *ChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte(cc.name))
}

func (cc *ChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	//return support.Handle(*cc.Extension, stub)
	ledger, err := support.HandlerTx(*cc.Extension, stub)
	if err != nil {
		return support.EncodeResponse(nil, err)
	}

	return process(stub, &TokenTransferService{
		ledgerStub: ledger,
	})
}
func process(stub shim.ChaincodeStubInterface, service *TokenTransferService) peer.Response {

	funcName, args := stub.GetFunctionAndParameters()

	payload := args[0]

	if funcName == `Transfer` {
		err := handleTransfer(payload, service)
		if err != nil {
			return support.EncodeResponse(nil, err)
		}
		return support.EncodeResponse(nil, nil)

	} else if funcName == `BatchTransfer` {
		err := handleBatchTransfer(payload, service)
		if err != nil {
			return support.EncodeResponse(nil, err)
		}
		return support.EncodeResponse(nil, nil)

	} else if funcName == `GetBalance` {
		account, err := handleGetBalance(payload, service)
		if err != nil {
			return support.EncodeResponse(nil, err)
		}
		return support.EncodeResponse(*account, nil)
	}

	return shim.Error(`function not yet implemented`)

}

func handleTransfer(payload string, service *TokenTransferService) error {
	transactions := make([]Transaction, 0)
	data, err := support.UnMarshal([]byte(payload), &transactions)
	if err != nil {
		return err
	}
	transaction := data.(Transaction)

	return service.Transfer(&transaction)
}

func handleBatchTransfer(payload string, service *TokenTransferService) error {

	ts := make([]Transaction, 0)
	err := json.Unmarshal([]byte(payload), &ts)
	if err != nil {
		return err
	}
	return service.BatchTransfer(ts)
}

func handleGetBalance(payload string, service *TokenTransferService) (*Account, error) {
	data, err := support.UnMarshal([]byte(payload), &Account{})
	if err != nil {
		return nil, err
	}
	acct := data.(*Account)
	return service.GetBalance(acct)
}

func NewCC() *ChainCode {
	return &ChainCode{
		name: `TokenTransferChainCode`,
		Extension: &support.Extension{
			FunctionAllowed: []string{`BatchTransfer`, `Transfer`, `GetBalance`},
			//Entity:          make([]Transaction, 0),
			//ServiceImpl:     &TokenTransferService{},
		},
	}

}

func main() {
	cc := NewCC()
	if err := shim.Start(cc); err != nil {
		fmt.Printf("Error starting %s: %s", cc.name, err)
	}
}
