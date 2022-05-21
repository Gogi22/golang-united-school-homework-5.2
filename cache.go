package cache

import "time"

type Value struct {
	val      string
	deadline time.Time
}

type Cache struct {
	dict map[string]Value
}

func NewCache() Cache {
	c := Cache{dict: make(map[string]Value)}
	return c
}

func (c Cache) Get(key string) (string, bool) {
	v, ok := c.dict[key]

	if !ok {
		return "", false
	}

	if !c.isValid(key) {
		delete(c.dict, key)
		return "", false
	}

	return v.val, true
}

func (c Cache) Put(key, value string) {
	c.dict[key] = Value{val: value}
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.dict))
	for key := range c.dict {
		if !c.isValid(key) {
			delete(c.dict, key)
		} else {
			keys = append(keys, key)
		}
	}

	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.dict[key] = Value{val: value, deadline: deadline}
}

func (c Cache) isValid(key string) bool {
	return c.dict[key].deadline.Equal(time.Time{}) || c.dict[key].deadline.After(time.Now())
}
