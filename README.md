# Expiry Cache
A simple expiry-based single-value caching.

## Example

### Import
```go
import "github.com/gavraz/expirycache/expirycache"
```

### Usage
```go
c := expirycache.NewInt64(time.Second, time.Now, fetchFunc)

value, err := c.Get()
if err != nil {
    return
}

fmt.Println("current value: ", value)
```