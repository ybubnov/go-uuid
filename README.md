# go-uuid - A wrapper for Linux kernel UUID v4 generator.

This UUID library is an yet another attempt to reimplement the wheel, but
instead of outstanding algorithm it wraps the Linux kernel implementation
of UUID v4 generator.


## Installation

The latest version can be installed using ```go``` tool:
```sh
go get github.com/ybubnov/go-uuid
```


## Usage

### Trivial configuration
The usage is pretty straightforward. Here is the most trivial example:
```go
package main

import (
    "fmt"

    "github.com/ybubnov/go-uuid"
)

func main() {
    u1 := uuid.New()
    fmt.Printf("uuid v4: %s\n", u1)
}
```

### Advanced configuration
Another example shows an advanced configuration of generator. It defines the
```128``` buffered UUIDs, and 16 workers used to produce them.

```go
package main

import (
    "fmt"

    "github.com/ybubnov/go-uuid"
)

func main() {
    src := uuid.Kernel{MaxIdle: 128, MaxProcs: 16}
    defer src.Stop() // Terminate source when time comes.

    u1, err := src.Next()
    if err != nil {
        fmt.Printf("failed to generate uuid: %s\n", err)
    }

    fmt.Printf("uuid v4: %s\n", u1)
}
```


## Licence

The UUID library is distributed under MIT license, therefore you are free to do
with code whatever you want. See the [LICENSE](LICENSE) file for full license text.
