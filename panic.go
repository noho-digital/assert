package assert

import (
	"fmt"
)

func (t *panicHandler) Errorf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

func newPanicHandler() *panicHandler {
	p := &panicHandler{}
	p.Assertions = New(p)
	return p
}

type panicHandler struct {
	*Assertions
}

var panicAssertions = newPanicHandler().Assertions

func PanicUnless() *Assertions {
	return panicAssertions
}



