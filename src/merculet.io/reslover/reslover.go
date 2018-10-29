package reslover

type handler func(payload interface{}) (interface{}, error)

type resolver struct {
	handlers map[string]handler
}

func NewResolver() *resolver {
	return &resolver{make(map[string]handler)}
}

func (r *resolver) Add(path string, handler handler) {
	r.handlers[path] = handler
}

type createEntityHandler struct {
}

func (c *createEntityHandler) create(entity interface{}) error {
	return nil
}
