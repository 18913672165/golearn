package trace

import "time"

type Handler func(error, ...string)
type Callback func(string, string, time.Time)

type Trace struct {
	Service string
	Method  string
	Start   time.Time
	handler []Handler

	callback []Callback
}

func New(service, method string) *Trace {
	return &Trace{
		Service: service,
		Method:  method,
		Start:   time.Now(),
	}
}

var std = &Trace{
	Service: "default",
	Method:  "default",
}

func SetService(service string) {
	std.Service = service
}
func SetMethod(method string) {
	std.Method = method
}

func RegisterHandler(handler Handler) {
	std.RegisterHandler(handler)
}

func RegisterCallback(callback Callback) {
	std.RegisterCallback(callback)
}

func (t *Trace) RegisterHandler(handler Handler) {
	t.handler = append(t.handler, handler)
}

func (t *Trace) RegisterCallback(callback Callback) {
	t.callback = append(t.callback, callback)
}

func (t *Trace) Error(err error, msg ...string) {
	for _, hdlr := range t.handler {
		hdlr(err, msg...)
	}
}

func (t *Trace) Close() {
	for _, callback := range t.callback {
		callback(t.Service, t.Method, t.Start)
	}
}
