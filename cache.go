package simple_cache1

type Cache struct {
	items map[string]string
}

func New() *Cache {
	return &Cache{
		make(map[string]string),
	}
}

func (c *Cache) Set(key, value interface{}) {
	c.items[key.(string)] = value.(string)
}

func (c *Cache) Get(key string) interface{} {
	return c.items[key]
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}
