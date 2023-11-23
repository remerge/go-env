# Global Env for Go

Package `env` provides a global `Env` variable, similar to `Rails.env`.

## Usage

```go
package main

import (
  "fmt"
  "github.com/remerge/go-env"
)

func main() {
  fmt.Println(env.Env)
}
```
