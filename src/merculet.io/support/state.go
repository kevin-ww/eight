package support

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
	"strings"
)

type LedgerStub struct {
	admin  string
	bucket string
	stub   shim.ChaincodeStubInterface
}

type ledgerStub interface {
	New(s *LedgerStub) interface{}
}

const defaultKeySeparator = `|`

var errorNoSuchRecord = errors.New(`no such record in ledger`)

func (l *LedgerStub) ledgerKey(k string) string {
	return strings.Join([]string{k, l.bucket, l.admin}, defaultKeySeparator)
}

func (l *LedgerStub) Put(k string, v interface{}) error {
	bytes, e := Marshal(v)
	if e != nil {
		return e
	}
	return l.stub.PutState(l.ledgerKey(k), bytes)
}

func (l *LedgerStub) Get(k string, target interface{}) (interface{}, error) {
	bytes, e := l.stub.GetState(l.ledgerKey(k))
	if e != nil {
		return nil, e
	}
	if bytes == nil {
		return nil, errorNoSuchRecord
	}
	return UnMarshal(bytes, target)
}

func (l *LedgerStub) Exists(k string) (bool, error) {
	bytes, e := l.stub.GetState(l.ledgerKey(k))
	if e != nil || bytes == nil {
		return false, e
	}

	return true, nil
}
