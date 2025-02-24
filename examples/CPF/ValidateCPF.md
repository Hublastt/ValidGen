---
sidebar_position: 1
---
# ValidateCPF

Validates whether a given CPF is valid.

Parameters:
- cpfString (string): The CPF to be validated (formatted or unformatted).

Return:
- error: Returns an error if the CPF is invalid, or `nil` if it is valid.

Usage example:
```go
package main

import (
	"fmt"
	validgen "github.com/Hublastt/ValidGen"
)

func main() {
	cpf := "123.456.789-09" // The CPF does not need to be formatted to be validated

	if err := validgen.ValidateCPF(cpf); err != nil {
		fmt.Println("Invalid CPF:", err)
	} else {
		fmt.Println("Valid CPF")
	}
}
```

Output:
```
Invalid CPF: <error message>
```
