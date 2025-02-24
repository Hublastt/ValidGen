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
cpf, err := cpf.CPFFunctions{}.FormatGeneratedCPF()
if err != nil {
    fmt.Println("Error generating CPF:", err)
} else {
    fmt.Println("Generated CPF:", cpf)
}
```
