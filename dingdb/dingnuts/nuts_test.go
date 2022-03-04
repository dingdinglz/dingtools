package dingnuts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDingNuts(t *testing.T) {
	dingNuts, err := NewDingNuts("./testdb")
	assert.NoError(t, err)
	assert.NotEqual(t, dingNuts, nil)

	err = dingNuts.Close()
	assert.NoError(t, err)
}
