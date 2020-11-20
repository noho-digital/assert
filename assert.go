package assert

import (
	"github.com/stretchr/testify/assert"
)

type Assertions struct {
	*assert.Assertions
}

func New(h Handler) *Assertions {
	return &Assertions{ assert.New(h) }
}

