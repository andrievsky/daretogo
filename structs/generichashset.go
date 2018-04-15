package structs

type GenericHashSet struct {
	keys map[int] struct{}
}

func NewGenericHashSet() *GenericHashSet{
	var set = new(GenericHashSet)
	set.keys = make(map[int] struct{})
	return set
}

func (t *GenericHashSet) Add(key int) bool{
	_, exist := t.keys[key]
	t.keys[key] = struct{}{}
	return !exist
}

func (t *GenericHashSet) QuickAdd(key int){
	t.keys[key] = struct{}{}
}

func (t *GenericHashSet) Contains(key int) bool{
	_, exist := t.keys[key]
	return exist
}

func (t *GenericHashSet) Count() int{
	return len(t.keys)
}