package structs

// HashSet for int impl
type GreedyHashSet struct {
	keys []bool
	count  int
}

// NewHashSet constructor
func NewGreedyHashSet(size int) *GreedyHashSet {
	var set = new(GreedyHashSet)
	set.keys = make([]bool, size)
	return set
}

// Add value to hashset. Return false if value already exists
func (t *GreedyHashSet) Add(key int) bool {
	if !t.keys[key] {
		t.keys[key] = true
		t.count++
		return true
	}

	return false
}

// Contains returns true if value is already in the Set
func (t *GreedyHashSet) Contains(key int) bool {
	return t.keys[key]
}

// Count of nodes
func (t *GreedyHashSet) Count() int {
	return t.count
}
