package support

import (
	"fmt"
	"reflect"
	"testing"
)

type ImplTest struct {
	//not implemented the new method
}

func (ts *ImplTest) New(s *LedgerStub) interface{} {
	return ts
}

func TestHasImplements(t *testing.T) {
	impl := ImplTest{}
	right := isModel(&impl)
	fmt.Printf("%v \n", right)
}

func isModel(i interface{}) bool {
	modelType := reflect.TypeOf((*ledgerStub)(nil)).Elem()
	return reflect.TypeOf(i).Implements(modelType)
}

func TestConstructCompositeKey(t *testing.T) {
	key := constructCompositeKey(`namesp`, `testkey`)
	fmt.Printf("%v \n", string(key))
}

var compositeKeySep = []byte{0x01}
var lastKeyIndicator = byte(0x01)

func constructCompositeKey(ns string, key string) []byte {
	return append(append([]byte(ns), compositeKeySep...), []byte(key)...)
}
