package support

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type Extension struct {
	FunctionAllowed []string
	Entity          interface{}
	ServiceImpl     interface{}
}

var (
	errNoSuchChainCodeFunction         = errors.New(`no such chaincode function provided`)
	errPayloadInvalid                  = errors.New(`require at least one arg as payload for chaincode invocation `)
	errFunctionNameShouldNotEmpty      = errors.New(`function name should not be empty`)
	errFunctionNameIsEmptyORNotAllowed = errors.New(`function name is empty or not allowed`)
)

func isFuncValid(funcName string, ext Extension) bool {

	funcName = strings.TrimSpace(funcName)

	if len(funcName) == 0 {
		return false
	}

	for _, allowed := range ext.FunctionAllowed {
		if funcName == allowed {
			return true
		}
	}

	return false
}

func postProcess(ext Extension, stub shim.ChaincodeStubInterface) error {

	return nil
}

func Handle(ext Extension, stub shim.ChaincodeStubInterface) peer.Response {

	funcName, args := stub.GetFunctionAndParameters()

	//is function name valid and allowed?

	if !isFuncValid(funcName, ext) {
		return EncodeResponse(nil, errFunctionNameIsEmptyORNotAllowed)
	}

	//is args valid and allowed?
	if args == nil || len(args) == 0 || len(args) != 1 {
		return EncodeResponse(nil, errPayloadInvalid)
	}

	fmt.Printf("invoking [%v] with args: %v \n", funcName, args[0])

	serviceImpl := ext.ServiceImpl

	if !hasImplemented(serviceImpl) {
		return EncodeResponse(nil, errors.New(fmt.Sprintf(`%v has not implemented the implied interface \n`, reflect.TypeOf(serviceImpl))))
	}

	//interface convention
	legerStubImpl := ext.ServiceImpl.(ledgerStub).New(&LedgerStub{
		admin:  "admin",
		bucket: "",
		stub:   stub,
	})

	//
	ccMethod := reflect.ValueOf(legerStubImpl).MethodByName(funcName)

	if ccMethod.Kind() == reflect.Invalid {
		return EncodeResponse(nil, errNoSuchChainCodeFunction)
	}

	//TODO

	return handlerInternal(ccMethod, []byte(args[0]), ext.Entity)
}

func handlerInternal(ccMethod reflect.Value, payload []byte, target interface{}) peer.Response {

	data, e := UnMarshal(payload, &target)

	if e != nil {
		return EncodeResponse(nil, errors.Wrap(e, `unable to unmarshal the payload as expected`))
	}

	//call
	var res []reflect.Value

	res = ccMethod.Call([]reflect.Value{reflect.ValueOf(data)})

	//TODO
	if err := res[len(res)-1]; !err.IsNil() {
		return EncodeResponse(nil, err.Interface().(error))
	}

	return EncodeResponse(res[0], nil)
}

func EncodeResponse(data interface{}, err error) peer.Response {
	if err != nil {
		return shim.Error(err.Error())
	}

	bytes, _ := json.Marshal(data)
	g

	return shim.Success(bytes)
}

func hasImplemented(i interface{}) bool {
	iType := reflect.TypeOf((*ledgerStub)(nil)).Elem()
	return reflect.TypeOf(i).Implements(iType)
}
