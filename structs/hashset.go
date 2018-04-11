package structs

// HashSet for int impl
type HashSet struct {
	values []LinkedList
	count  int
}

// NewHashSet constructor
func NewHashSet(size int) *HashSet {
	var set = new(HashSet)
	set.values = make([]LinkedList, size)
	return set
}

// Add value to hashset. Return false if value already exists
func (t *HashSet) Add(value int) bool {
	var h = t.hash(value)
	if t.values[h].Contains(value) {
		return false
	}
	t.values[h].Add(value)
	t.count++
	return true
}

// Contains returns true if value is already in the Set
func (t *HashSet) Contains(value int) bool {
	var h = t.hash(value)
	return t.values[h].Contains(value)
}

// Count of nodes
func (t *HashSet) Count() int {
	return t.count
}

func (t *HashSet) hash(value int) int {
	return value % len(t.values)
}
