# Cachew

Generic in-memory caching
- Suitable for holdin small amount of data in your go application.
- Suitable for concurrent read and write
- Can be used with all data types


# Usage
```bash 
go get github.com/mehran75/cachew
```

```go
import "github.com/mehran75/cachew/memorycache"

// key is from type string (comparable)
// data type can be anything
cachew := memorycache.NewCache[string, int]()

cachew.Write("key", 3, 10*time.seconds)

fmt.Println(cahcew.Read("key"))

cachew.Delete("key")

```


# Example
check [here](https://github.com/mehran75/cachew/blob/main/example/main/main.go)
