package hash

import (
	"math/rand"
	"reflect"
)

type RandomizedSet struct {
	m map[int]struct{}
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{m: make(map[int]struct{})}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = struct{}{}
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	delete(this.m, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	keys := reflect.ValueOf(this.m).MapKeys()
	return keys[rand.Intn(len(keys))].Interface().(int)
}
