# pinyin - 🏳️An simple pinyin messages manage implementation for Go.

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/pinyin)](https://pkg.go.dev/github.com/go-zoox/pinyin)
[![Build Status](https://github.com/go-zoox/pinyin/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/pinyin/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/pinyin)](https://goreportcard.com/report/github.com/go-zoox/pinyin)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/pinyin/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/pinyin?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/pinyin.svg)](https://github.com/go-zoox/pinyin/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/pinyin.svg?label=Release)](https://github.com/go-zoox/pinyin/releases)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/pinyin
```

## Getting Started

```go
package main

import (
	"context"

	"github.com/go-zoox/pinyin"
)

func main() {
	fmt.Println(pinyin.String("你好世界"))
	// Output: nihaoshijie

	fmt.Println(pinyin.Slice("你好世界"))
	// Output: [ni hao shi jie]

	fmt.Println(pinyin.Abbr("你好世界"))
	// Output: nhsj
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
