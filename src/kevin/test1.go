package main

import "fmt"

//type handler func(payload string, srv interface{}) string


//func (h handler) testFunc(payload string, srv interface{}) string{
//	return h(payload, srv)
//}

func testHandler(payload string, srv interface{}) string {
	out := payload + `-kevin`
	fmt.Printf(`%s \v`, out)
	srv.(*testService).serviceMethod(payload)
	return out
}

type testService struct {
}

func (t *testService) serviceMethod(m string) error {
	fmt.Println(`test service`)
	return nil
}

func main() {
	r := NewRouter().add(`test`, testHandler)
	r.handle(`test`, `test-payload`)
	//testHandler
}
