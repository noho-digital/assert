package assert

import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

func (t *Panic) Errorf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

func NewPanic() *Panic {
	p := &Panic{}
	p.Assertions = assert.New(p)
	return p
}

type Panic struct {
	*assert.Assertions
}