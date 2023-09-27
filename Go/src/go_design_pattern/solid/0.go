/*
开放封闭原则：模块可扩展，而不可修改。也就是说，模块对扩展开放，而对修改封闭
*/

package solid

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type RedisCache struct{}

func (r *RedisCache) Get(key string) (string, error) {
	// Retrieve value from Redis
	return "", nil
}

func (r *RedisCache) Set(key, value string) error {
	// Set value in Redis
	return nil
}

// 上面的 Cache 接口和 RedisCache 实现都设计为开放扩展。如果我们想要添加一个新的缓存机制，比如 Memcached，我们可以在不修改现有代码的情况下这样做：

type MemcachedCache struct{}

func (m *MemcachedCache) Get(key string) (string, error) {
	// Retrieve value from Memcached
	return "", nil
}

func (m *MemcachedCache) Set(key, value string) error {
	// Set value in Memcached
	return nil
}
