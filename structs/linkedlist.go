package structs

type Node struct {
	next  *Node
	value int
}

// SimpleCollection generic interface
type SimpleCollection interface {
	Add(value int)
	Contains(value int) bool
	Count() int
}

// LinkedList simple impl. with int values for HashSet and HashMap
type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

// Add a value
func (t *LinkedList) Add(value int) {
	var node = &Node{value: value}
	if t.head == nil {
		t.head = node
		t.tail = t.head
		t.count++
		return
	}

	t.tail.next = node
	t.tail = node
	t.count++
}

// Contains return true if value is in the list
func (t *LinkedList) Contains(value int) bool {
	var node = t.head
	for node != nil {
		if node.value == value {
			return true
		}
		node = node.next
	}
	return false
}

// Count of nodes
func (t *LinkedList) Count() int {
	return t.count
}
