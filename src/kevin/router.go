package main

type Router struct {
	handlers map[string]handler
	impl service
}

type service interface {
	serviceMethod(m string) error
}

func NewRouter(impl service) *Router {
	//handlers = make(map[string]handler)
	return &Router{
		handlers: make(map[string]handler),
		impl:impl,
	}
}

func (r *Router) add(path string, f func(payload string, srv interface{}) string) *Router {
	//r.handlers[path] = h.(handler)
	r.handlers[path]=f
	return r
}

func (r *Router) handle(path string, payload string) error {
	h := r.handlers[path]
	h(payload, r.impl)
	return nil
}

//type handler func(payload string,srv service ) string

type handler func(payload string, srv interface{}) string
