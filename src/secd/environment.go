package secd

type Environment struct {
	refs []interface{}
}

func NewEnvironment() *Environment {
	return &Environment {make([]interface{}, 0)}
}

func (e *Environment) Add(val interface{}) {
	e.refs = append(e.refs, val)
}

func (e *Environment) Get(index int) interface{} {
	return e.refs[index]
}

func (e *Environment) Set(index int, val interface{}) {
	e.refs[index] = val
}
