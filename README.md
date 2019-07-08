# Ref

[![CircleCI branch](https://img.shields.io/circleci/project/github/trivigy/ref/master.svg?label=master&logo=circleci)](https://circleci.com/gh/trivigy/workflows/ref)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE.md)
[![](https://godoc.org/github.com/trivigy/ref?status.svg&style=flat)](http://godoc.org/github.com/trivigy/ref)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/trivigy/ref.svg?style=flat&color=e36397&label=release)](https://github.com/trivigy/ref/releases/latest)

## Introduction
In golang v1, unfortunately, there is no nice builtin way to represent optional 
`struct` fields for primitives whose zero value is individualized for each 
primitive. A lot of debating has been done and people hold different opinions 
as to what is the best solution. This library is my personal preferred way
of dealing with the situation.

## Example
```go
package main

import (
	"fmt"
	
	"github.com/trivigy/ref"
)

type Sample struct {
	Field *string `json:"field,omitempty"`
}

func main() {
	result := Sample{
		Field: ref.String("example"),
	}
	
	if result.Field != nil {
		fmt.Print("'Field' is set")
	}
}
```

## Alternatives
* https://www.reddit.com/r/golang/comments/1tst88/optional_field_in_a_struct_pointer_vs_nullintlike/

If you hold another opinion feel free to share it either via discord or on 
golang github channels. Or better yet share a PR with a new link added above.
