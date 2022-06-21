package main

import (
	"fmt"
	"github.com/expirycache/expirycache"
	"time"
)

func fetch() (int64, error) {
	// may be slow/costly etc.
	return 0, nil
}

func main() {
	c := expirycache.NewInt64(time.Second, time.Now, fetch)
	fmt.Println(c.Get())
}
