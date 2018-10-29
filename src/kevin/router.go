package main

type Router struct {
	handlers map[string]handler
}

type service interface {
	serviceMethod(m string) error
}

func NewRouter() *Router {
	//handlers = make(map[string]handler)
	return &Router{
		handlers: make(map[string]handler),
	}
}

func (r *Router) add(path string, f func(payload string, srv interface{}) string) *Router {
	//r.handlers[path] = h.(handler)
	r.handlers[path]=f
	return r
}

func (r *Router) handle(path string, payload string) error {
	h := r.handlers[path]
	h(payload, nil)
	return nil
}

//type handler func(payload string,srv service ) string

type handler func(payload string, srv interface{}) string
