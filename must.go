package assert

import (
	"fmt"
	"sync"
)

func MustNot(cnd bool, msg string, args ...interface{}) {
	if cnd {
		return
	}
	panic(fmt.Errorf(msg, args...))
}

var mustAssertions *Assertions
var mustOnce sync.Once
func Must() *Assertions {
	mustOnce.Do(func() {
		mustAssertions = newPanicHandler().Assertions
	})
	return mustAssertions
}

