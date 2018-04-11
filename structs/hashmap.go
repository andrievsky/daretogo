package structs

type HashMap struct {
	keys   []string
	values []*interface{}
}

func (t *HashMap) New() {

}

func (t *HashMap) Add(key string, value *interface{}) {

}

func (t *HashMap) HasKey(key string) bool {
	for _, k := range t.keys {
		if key == k {
			return true
		}
	}
	return false
}

/*
func hash(key string) int {

}
*/
