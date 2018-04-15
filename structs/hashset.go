package structs

// HashSet for int impl
type HashSet struct {
	keys []LinkedList
	count  int
}

// NewHashSet constructor
func NewHashSet(size int) *HashSet {
	var set = new(HashSet)
	set.keys = make([]LinkedList, size)
	return set
}

// Add value to hashset. Return false if value already exists
func (t *HashSet) Add(key int) bool {
	var h = t.hash(key)
	if t.keys[h].Contains(key) {
		return false
	}
	t.keys[h].Add(key)
	t.count++
	return true
}

func (t *HashSet) QuickAdd(key int){
	var h = t.hash(key)
	if t.keys[h].Contains(key) {
		return
	}
	t.keys[h].Add(key)
	t.count++
}

// Contains returns true if value is already in the Set
func (t *HashSet) Contains(key int) bool {
	return t.keys[t.hash(key)].Contains(key)
}

// Count of nodes
func (t *HashSet) Count() int {
	return t.count
}

func (t *HashSet) hash(key int) int {
	return key % len(t.keys)
}
