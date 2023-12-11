# Golang in-memory cache

Project for Golang basic training: structs, methods, pointers, errors and packages.

## Installation
```bash
go get "github.com/psevdocoder/InMemorycache"
```

## Usage example

```go
package main

import (
	"fmt"
	cache "github.com/psevdocoder/InMemorycache"
	"log"
)


func main() {
	c := cache.New()

	setErr := c.Set("userId", 42)
	if setErr != nil {
		log.Panic(setErr)
	}

	userId, getErr := c.Get("userId")
	if getErr != nil {
		log.Panic(getErr)
	}
	fmt.Println(userId)

	fmt.Println(*c)

	if delErr := c.Delete("userId"); delErr != nil {
		log.Panic(delErr)
	}

	fmt.Println(userId)
	fmt.Println(*c)
}
```