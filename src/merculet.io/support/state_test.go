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
