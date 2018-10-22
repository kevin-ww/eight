package main

import (
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
	return shim.Success([]byte("Initial ..." + cc.name))
}

func (cc *ChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	return support.Handle(*cc.Extension, stub)
}

func main() {

	cc := &ChainCode{
		name: `TokenChainCode`,
		Extension: &support.Extension{
			FunctionAllowed: []string{`get`, `has`, `create`, `update`},
			Entity:          &Token{},
			ServiceImpl:     &TokenService{},
		},
	}

	if err := shim.Start(cc); err != nil {
		fmt.Printf("Error starting %s: %s", cc.name, err)
	}
}
