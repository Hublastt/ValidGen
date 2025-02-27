# ValidGen

<img align="right" width="250px" src="static/assets/go_logo.png">



**Valigen** is a  Go library for generating and validating various types of structured data, such as Brazilian CPF, CNPJ, and other formatted identifiers. It provides an efficient and easy-to-use API to handle data validation, formatting, and generation.

If you need a reliable solution for working with structured data, **Valigen** is the perfect choice.

## Key Features

- **Data Generation** – Generate valid CPF, CNPJ, and other structured identifiers.
- **Validation** – Easily check if a given value is correctly formatted and valid.
- **Formatting & Unformatting** – Convert data between formatted and raw versions.
- **Flexible & Extensible** – Supports multiple data types and validation rules.
- **User-Friendly API** – Simple and intuitive integration with any Go project.
- **High Performance** – Optimized for speed and efficiency.

For more details and usage examples, check out the **Valigen** repository.

##  **Installation**

To install the ValidGen package, use the following command:

```sh
go get github.com/HublastX/ValidGen
```

## Usage

Below are examples of how to use the various functions provided by the ValidGen package.

<details>
<summary>Validate CPF</summary>

To validate a CPF number:

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
</details>

<details>
<summary>Generate CPF</summary>

To generate a valid CPF number:

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

</details>

<details>
<summary>Format Generated CPF</summary>

To format a generated CPF number:

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

</details>

<details>
<summary>Format CPF</summary>

To format a CPF number:

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

</details>

<details>
<summary>Unformat CPF</summary>

To unformat a CPF number:

```go
package main

import (
    "fmt"
    validgen "github.com/Hublastt/ValidGen"
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

Output:

```
Unformatted CPF: 12345678909
```

</details>



## Contributing

Contributions are welcome! If you would like to contribute, feel free to open a pull request.

## License

This project is licensed under the MIT License.
