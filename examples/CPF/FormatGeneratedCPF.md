---
sidebar_position: 5
---

# FormatGeneratedCPF

Generates a valid random CPF and returns it formatted.

Return:
- string: The formatted CPF.
- error: An error if something goes wrong.

Usage example:
```go
package main

import (
    "fmt"
    validgen "github.com/Hublastt/ValidGen"
)

func main() {
    cpf, err := validgen.FormatGeneratedCPF()
    
    if err != nil {
        fmt.Println("Error generating CPF:", err)
    } else {
        fmt.Println("Generated CPF:", cpf)
    }
}
```

Output:
```
Generated CPF: 123.456.789-09
```
