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
import (
    "fmt"
    "github.com/MichaelCombs28/goesprima"
)

func main() {
    g := goesprima.NewGenerator()
    g.AddStatements(
        &ImportDeclaration{
            Source: "@aws-amplify/core",
            Specifiers: []ImportDeclarationSpecifier{
                &ImportDefaultSpecifier{
                    Local: &Identifier{
                        Name: "Amplify",
                    },
                },
            },
        },
    )
    fmt.Println(g.String())
}
```

## Roadmap

- Parsing & Tokenization
- Code Execution

## License

MIT
