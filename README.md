
**Cacher - Simple go caching library**

``Overview``

Provides API for caching any type of data(race condition proof)

`usage:`


import it.
~~~~
import "github.com/nieless/cacher"
~~~~

Cache any type of item (it also accepts expiration time of item, it will delete item from cache if expiry time is less than 1 minute of current time).
~~~~
value := cacher.SetCacheItem("key","hey there",nil)
~~~~

Get cached item (it will return value as a nil if key doesn't exist in cache).
~~~~
value := cacher.GetCachedItem("key")
~~~~

Delete cached item (it will return value as a nil if key doesn't exist in cache).
~~~~
value := cacher.DeleteCachedItem("key")
~~~~
