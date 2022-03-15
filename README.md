# Go Esprima

Esprima JS implementation in Go for code generation & script execution. Based on the [Esprima AST](https://esprima.org/) and
the [JQuery implementation](https://github.com/jquery/esprima/).

## Installation

```
go install github.com/MichaelCombs28/goesprima
```

## Usage

Currently goesprima only supports manual AST code generation.

```
package main

import (
	"fmt"

	esp "github.com/MichaelCombs28/goesprima"
)

func main() {
	gen := esp.NewGenerator()
	gen.AddStatements(
		&esp.ImportDeclaration{
			Source: "@aws-amplify/core",
			Specifiers: []esp.ImportDeclarationSpecifier{
				&esp.ImportDefaultSpecifier{
					Local: &esp.Identifier{
						Name: "Amplify",
					},
				},
			},
		},
	)
	fmt.Println(gen.String())
}
```

## Roadmap

- Parsing & Tokenization
- Code Execution

## License

MIT
