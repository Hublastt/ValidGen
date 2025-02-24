---
sidebar_position: 4
---

# UnformatCPF
Removes formatting from a CPF.

Parameters:
- cpfString (string): The CPF formatted in the `XXX.XXX.XXX-XX` pattern.

Return:
- string: The unformatted CPF.
- error: An error if the CPF is invalid.

## Usage example

```go
package main

import (
    "fmt"
    "github.com/yourusername/validgen"
)

func main() {
    unformattedCPF, err := validgen.UnformatCPF("123.456.789-09")

    if err != nil {
        fmt.Println("Error removing formatting:", err)
    } else {
        fmt.Println("Unformatted CPF:", unformattedCPF)
    }
}
```
