# Cachew

generic memory caching in golang


# Usage
```go
// key is from type string (comparable)
// data type can be anything
cachew := memorycache.NewCache[string, int]()

cachew.Write("key", 3, 10*time.seconds)

fmt.Println(cahcew.Read("key"))

cachew.Delete("key")

```
