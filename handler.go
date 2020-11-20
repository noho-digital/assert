package assert

import "github.com/stretchr/testify/assert"

type Handler interface {
	assert.TestingT
}
