package cache

type Cache struct {
	data map[string]interface{}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c Cache) Get(key string) interface{} {
	return c.data[key]
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

func New() Cache {
	return Cache{make(map[string]interface{})}
}
