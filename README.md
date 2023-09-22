# Lazy

The tiniest API for lazy initialization.

```go
package main

import "github.com/melt-inc/lazy"

const x = lazy.New(func() string {
    return "loaded"
})

func main() {
    fmt.Println(x())
}
```
