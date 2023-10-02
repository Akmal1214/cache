package cache

type Cache struct {
	items map[string]Item
}

type Item struct {
	Value interface{}
}

func New() *Cache {
	items := make(map[string]Item)
	cache := Cache{
		items: items,
	}
	return &cache
}

func (c *Cache) Set(key int, value interface{}) {
	c.items[key] = Item{
		Value: value,
	}
}

func (c *Cache) Get(key int) interface{} {
	item := c.items[key]
	return item
}

func (c *Cache) Delete(key int) interface{} {
	item := c.items[key]
	delete(c.items, key)
	return item
}
