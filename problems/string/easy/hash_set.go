package string

type MyHashSet struct {
	values map[any]struct{}
}

func Constructor() MyHashSet {
	return MyHashSet{
		values: map[any]struct{}{},
	}
}

func (this *MyHashSet) Add(key int) {
	this.values[key] = struct{}{}

}

func (this *MyHashSet) Remove(key int) {
	delete(this.values, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, exists := this.values[key]
	return exists
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
