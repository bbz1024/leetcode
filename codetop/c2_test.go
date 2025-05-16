package codetop

import "testing"

func TestLRUCache(t *testing.T) {
	lruCache := Constructor(2) // null
	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	lruCache.Put(1, 1)     // null
	lruCache.Put(2, 2)     // null
	lruCache.Get(1)        // 1
	lruCache.Put(3, 3)     // null
	t.Log(lruCache.Get(1)) // 1
	t.Log(lruCache.Get(2)) // -1
	lruCache.Put(4, 4)
	t.Log(lruCache.Get(1)) // -1
	lruCache.Get(4)        // 4

}
