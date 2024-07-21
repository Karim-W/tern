# Tern

TernAry is a simple, lightweight, and easy-to-use function to implement a
ternary statement.

## Usage

```go
package main

import (
    "fmt"
    "github.com/karim-w/tern"
)

func main() {
    a := 1
    b := 2
    fmt.Println(tern.Ary(a > b, a, b))
}
```

## License

BSD 3-Clause License

## Author

karim-w

## Contributing

Pull requests are welcome. For major changes, please open an issue first to
discuss what you would like to change.
