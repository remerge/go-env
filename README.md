# go-env

Package `env` provides a global `Env` variable, similar to `Rails.env`.

## Install

```bash
go get github.com/remerge/go-env
```

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
