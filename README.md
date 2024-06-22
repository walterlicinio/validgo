# ValidGo

ValidGo is flexible data validation library for Go, providing a fluent interface for validating complex data structures.

## Installation

To install ValidGo, use `go get`:

```bash
go get github.com/walterlicinio/validgo
```

## Usage

Here's a quick example of how to use ValidGo:

```go
package main

import (
    "fmt"
    "github.com/walterlicinio/validgo"
)

type User struct {
    Name  string
    Age   int
    Email string
}

func main() {
    user := User{Name: "John Doe", Age: 25, Email: "john@example.com"}

    v := validgo.New(user)
    v.Field("Name").IsString().MinLength(2).MaxLength(50)
    v.Field("Age").IsInt().Min(18).Max(120)
    v.Field("Email").IsEmail()

    if v.HasErrors() {
        for _, err := range v.Errors() {
            fmt.Println(err)
        }
    } else {
        fmt.Println("Validation passed")
    }
}
```

## Features

- Fluent interface for easy and readable validation rules
- Support for nested struct validation
- Built-in validators for common types (string, int, float)
- Custom validation rules
- Data transformation during validation
- Comprehensive error collection

## Available Validators

- `IsString()`, `MinLength()`, `MaxLength()`
- `IsInt()`, `Min()`, `Max()`
- `IsFloat()`, `MinFloat()`, `MaxFloat()`
- `IsEmail()`
- `IsPositive()`
- `Matches()` (regex matching)
- `IsAlpha()`, `IsAlphanumeric()`, `IsNumeric()`
- `Custom()` (for custom validation rules)
- `Transform()` (for data transformation)


## Advanced Usage Examples

### Custom Validation

You can use the `Custom()` method to implement custom validation rules. Here's an example that checks if the user's name contains a space:

```go
user := User{Name: "John Doe", Age: 25, Email: "john@example.com"}

v := validgo.New(user)
v.Field("Name").Custom(func(value interface{}) bool {
    name, ok := value.(string)
    if !ok {
        return false
    }
    return strings.Contains(name, " ")
}).WithError("Name must contain a space")

if v.HasErrors() {
    for _, err := range v.Errors() {
        fmt.Println(err)
    }
} else {
    fmt.Println("Validation passed")
}
```

## Contributing

Contributions to ValidGo are welcome! Please feel free to submit a Pull Request.