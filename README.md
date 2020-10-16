[![Build Status](https://travis-ci.org/hasansino/environment.svg?branch=master)](https://travis-ci.org/hasansino/environment)
[![Go Report Card](https://goreportcard.com/badge/github.com/hasansino/environment)](https://goreportcard.com/report/github.com/hasansino/environment)

# environment
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