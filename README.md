# Cache

Go Cache поможет записать значения *value* в кеш по ключу *key*

# Example

```golang
package main

import (
	"cache/cache"
	"fmt"
)

func main() {
	cache := cache.New()
  
	cache.Set("userId", 42)
	fmt.Println(cache.Get("userId")) // 42
  
	cache.Delete("userId")
	fmt.Println(cache.Get("userId")) // nil
}
```
