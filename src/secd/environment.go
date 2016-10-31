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

func (e *Environment) Get(index int64) interface{} {
	return e.refs[index]
}

func (e *Environment) Set(index int64, val interface{}) {
	e.refs[index] = val
}
