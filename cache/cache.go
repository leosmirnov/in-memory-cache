package cache

type Cache struct {
	cache map[string]item
}

type item struct {
	value interface{}
	ttl   int
}

func (c *Cache) Get() {

}

func (c *Cache) Set() {

}

func (c *Cache) Remove() {

}

func (c *Cache) Keys() {

}
