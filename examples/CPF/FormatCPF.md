---
sidebar_position: 2
---

# FormatCPF

Formats an unformatted CPF to the standard `XXX.XXX.XXX-XX` pattern.

Parameters:
- cpfString (string): The unformatted CPF.

Return:
- string: The formatted CPF.
- error: An error if the CPF is invalid.

Usage example:
```go
package main

import (
    "fmt"
    validgen "github.com/Hublastt/ValidGen"
)

func main() {

    formattedCPF, err := validgen.FormatCPF("12345678909")

    if err != nil {
        fmt.Println("Error formatting CPF:", err)
    } else {
        fmt.Println("Formatted CPF:", formattedCPF)
    }
}
```

Output:
```
Formatted CPF: 123.456.789-09
```
