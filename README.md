# go-env

[![<keitaro1020>](https://circleci.com/gh/keitaro1020/go-env.svg?style=shield)]()
[![codecov](https://codecov.io/gh/keitaro1020/go-env/branch/main/graph/badge.svg)](https://codecov.io/gh/keitaro1020/go-env)

go-env is a library in Go that maps environment variables to structs.

## Usage
install
```shell
$ go get -u github.com/keitaro1020/go-env
```

environment variables
```shell
export PRODUCTION=true
export DATABASE_HOST=127.0.0.1
export DATABASE_PORT=3306
export DATABASE_USER=root
export DATABASE_PASS=password
```

writing code
```go
package main

import (
	"fmt"

	"github.com/keitaro1020/go-env"
)

type Config struct {
	Production bool
	DBConfig   DatabaseConfig `env_key:"DATABASE"`
}

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Pass string
}

func main() {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Errorf("error: %v\n", err)
	}
	fmt.Printf("cfg: %+v\n", cfg)
}
```

run
```shell
$ go run main.go     
cfg: &{Production:true Database:{Host:127.0.0.1 Port:3306 User:root Pass:password}}
```

## Supported types
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64
- float32, float64
- bool
- slices of any supported type
