package simple_cache1

type Cache struct {
	items map[string]interface{}
}

func New() *Cache {
	return &Cache{
		make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.items[key]
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}
