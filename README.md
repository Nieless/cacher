**Simple Cache**

`usage:`

import "github.com/nieless/cacher"

* cache any type of item 

value := cacher.SetCacheItem("key","hey there",nil)

it also accepts expiration time of item, it will delete item from cache if expiry time is less than 1 minute of current time.


* get cached item

value := cacher.GetCachedItem("key")
it will return value as a nil if key doesn't exist in cache.

* delete cached item

value := cacher.DeleteCachedItem("key")

it will return value as a nil if key doesn't exist in cache.
