package structs_test

import (
	"testing"

	"github.com/andrievsky/daretogo/structs"
	"github.com/stretchr/testify/assert"
)

func TestHashSetInit(t *testing.T) {
	var set = structs.NewHashSet(10)
	assert.Equal(t, false, set.Contains(1))
	assert.Equal(t, 0, set.Count())
}

func TestHashSet(t *testing.T) {
	var set = structs.NewHashSet(10)
	assert.Equal(t, true, set.Add(1))

	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, 1, set.Count())

	assert.Equal(t, false, set.Add(1))
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, 1, set.Count())
}

func TestGenericHashSet(t *testing.T) {
	var set = structs.NewGenericHashSet()
	assert.Equal(t, true, set.Add(1))

	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, 1, set.Count())

	assert.Equal(t, false, set.Add(1))
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, 1, set.Count())
}

func TestGreedyHashSet(t *testing.T) {
	var set = structs.NewGreedyHashSet(10)
	assert.Equal(t, true, set.Add(1))

	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, 1, set.Count())

	assert.Equal(t, false, set.Add(1))
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, 1, set.Count())
}

func BenchmarkHashSet(b *testing.B) {
	var set = structs.NewHashSet(b.N)
	for n := 0; n < b.N; n++ {
		set.Add(n)
		set.Contains(n)
		set.Count()
	}
}

func BenchmarkGenericHashSet(b *testing.B) {
	var set = structs.NewGenericHashSet()
	for n := 0; n < b.N; n++ {
		set.Add(n)
		set.Contains(n)
		set.Count()
	}
}

func BenchmarkGreedyHashSet(b *testing.B) {
	var set = structs.NewGreedyHashSet(b.N)
	for n := 0; n < b.N; n++ {
		set.Add(n)
		set.Contains(n)
		set.Count()
	}
}