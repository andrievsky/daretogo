package structs_test

import (
	"testing"

	"github.com/andrievsky/daretogo/structs"
	"github.com/stretchr/testify/assert"
)

func TestHasSetInit(t *testing.T) {
	var set = structs.NewHashSet(10)
	assert.Equal(t, false, set.Contains(1))
	assert.Equal(t, 0, set.Count())
}

func TestHasSetAdd(t *testing.T) {
	var set = structs.NewHashSet(10)
	assert.Equal(t, true, set.Add(1))

	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, 1, set.Count())

	assert.Equal(t, false, set.Add(1))
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, 1, set.Count())
}
