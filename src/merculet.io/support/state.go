package support

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
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
	return strings.Join([]string{l.admin, l.bucket, k}, defaultKeySeparator)

	//return append(append([]byte(ns), compositeKeySep...), []byte(key)...)

}

func (l *LedgerStub) Put(k string, v interface{}) error {
	bytes, e := Marshal(v)
	if e != nil {
		return e
	}
	lk := l.ledgerKey(k)
	//TODO
	fmt.Println(`PUT `, lk)

	return l.stub.PutState(lk, bytes)
}

func (l *LedgerStub) Get(k string, target interface{}) (interface{}, error) {
	lk := l.ledgerKey(k)

	//TODO
	fmt.Println(`GET  `, lk)

	bytes, e := l.stub.GetState(lk)
	if e != nil {
		return nil, e
	}
	if bytes == nil {
		//no such record in ledger
		return nil, nil
	}
	return UnMarshal(bytes, target)
}

func (l *LedgerStub) Exists(k string) (bool, error) {
	lk := l.ledgerKey(k)

	//TODO
	fmt.Println(`GET `, lk)

	bytes, e := l.stub.GetState(lk)
	if e != nil || bytes == nil {
		return false, e
	}

	return true, nil
}

// GetWithMultipleKeys implements method
func (l *LedgerStub) GetWithMultipleKeys(keys []string) ([]interface{}, error) {
	vals := make([]interface{}, len(keys))
	for i, key := range keys {
		val, err := l.stub.GetState(key)
		if err != nil {
			return nil, err
		}
		vals[i] = val
	}
	return vals, nil
}

// ApplyUpdates implements method in VersionedDB interface
func (l *LedgerStub) ApplyUpdates(batch *statedb.UpdateBatch, height *version.Height) error {
	//dbBatch := leveldbhelper.NewUpdateBatch()
	//namespaces := batch.GetUpdatedNamespaces()
	//for _, ns := range namespaces {
	//	updates := batch.GetUpdates(ns)
	//	for k, vv := range updates {
	//		compositeKey := constructCompositeKey(ns, k)
	//		logger.Debugf("Channel [%s]: Applying key(string)=[%s] key(bytes)=[%#v]", vdb.dbName, string(compositeKey), compositeKey)
	//
	//		if vv.Value == nil {
	//			dbBatch.Delete(compositeKey)
	//		} else {
	//			dbBatch.Put(compositeKey, EncodeValue(vv.Value, vv.Version))
	//		}
	//	}
	//}
	//dbBatch.Put(savePointKey, height.ToBytes())
	//// Setting snyc to true as a precaution, false may be an ok optimization after further testing.
	//if err := vdb.db.WriteBatch(dbBatch, true); err != nil {
	//	return err
	//}
	return nil
}
