package cache

type Cache interface {
	Users() UserCache
	Close()
}

var cli *Cache

func CacheFactory() Cache {
	return *cli
}

func SetCacheFactory(cache Cache) {
	cli = &cache
}
