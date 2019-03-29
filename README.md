A simple package to add piping feature to Go cli tools

## Installation

`go get -u go.xitonix.io/cli`

## Usage

```go
package main

import (
	"fmt"
	"log"

	"go.xitonix.io/cli"
)

func main() {
	input, chained, err := cli.PipeInIfChained()
	if err != nil {
		log.Fatal(err)
	}
	if chained {
		fmt.Printf("Executed in pipe mode\nInput: %s", input)
	} else {
		fmt.Println("Not piped in")
	}
}

```

