# Cachew

generic memory caching in golang. 
- Suitable for holdin small amount of data in your go application.
- Suitable for concurrent read and write

# Usage
```go
// key is from type string (comparable)
// data type can be anything
cachew := memorycache.NewCache[string, int]()

cachew.Write("key", 3, 10*time.seconds)

fmt.Println(cahcew.Read("key"))

cachew.Delete("key")

```
