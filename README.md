# Ref

[![Discord](https://img.shields.io/discord/428990244952735764.svg?style=flat&logo=discord&colorB=green)](https://discord.gg/M9nxJ3g)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE.md)
[![Release](https://img.shields.io/badge/release-v1.0.1-e36397.svg)](https://github.com/syncaide/ref/releases/tag/v1.0.1)

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
	
	"github.com/syncaide/ref"
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