[![Build Status](https://travis-ci.org/hasansino/gcmap.svg?branch=main)](https://travis-ci.org/hasansino/gcmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/hasansino/environment)](https://goreportcard.com/report/github.com/hasansino/environment)

# gcmap
Tiny package that provides environment awareness

# Installation

```bash
~ $ go get -u github.com/hasansino/environment
```

# Example usage
```go
import "github.com/hasansino/environment"

func main() {
    env := environment.GetEnvironment()
    if env.IsDevelopment() {
    	// do some stuff
    }
    
    // override value
    environment.OverrideEnvironment(environment.Production)
}
```