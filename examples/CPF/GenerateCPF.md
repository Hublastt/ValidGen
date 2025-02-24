---
sidebar_position: 3
---

# GenerateCPF

Generates a valid random CPF.

Return:
- string: The generated CPF.
- error: An error if something goes wrong.

Usage example:
```go
package main

import (
    "fmt"
    validgen "github.com/Hublastt/ValidGen"
)

func main() {
    cpf, err := validgen.GenerateCPF()
    if err != nil {
        fmt.Println("Error generating CPF:", err)
    } else {
        fmt.Println("Generated CPF:", cpf)
    }
}
```

Output:
```
Generated CPF: 12345678909
```
