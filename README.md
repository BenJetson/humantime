# HumanTime
Golang library to create human-readable US English strings from time.Duration 
instances.

## Installation

```
go get github.com/BenJetson/humantime
```

## Import

```golang
import "github.com/BenJetson/humantime"
```

## Examples

### Example using `func humantime.Since(t time.Time) string`

```golang
import (
    "fmt"
    "time"
    
    "github.com/BenJetson/humantime"
)

func main() {
    t := time.Now()

    time.Sleep(5)

    fmt.Println(humantime.Since(t)) // just now

    time.Sleep(30)

    fmt.Println(humantime.Since(t)) // seconds ago

    time.Sleep(25)

    fmt.Println(humantime.Since(t)) // a minute ago

    time.sleep(240)

    fmt.Println(humantime.Since(t)) // 5 minutes ago
}
```

### Example using `func humantime.Duration(d time.Duration) string`

```golang
import (
    "fmt"
    "time"
    
    "github.com/BenJetson/humantime"
)

func main() {

    var d time.Duration

    d, _ = time.ParseDuration("5s")
    fmt.Println(humantime.Since(d)) // just now

    d, _ = time.ParseDuration("25s")
    fmt.Println(humantime.Since(d)) // seconds ago

    d, _ = time.ParseDuration("60s")
    fmt.Println(humantime.Since(d)) // a minute ago

    d, _ = time.ParseDuration("5m")
    fmt.Println(humantime.Since(d)) // 5 minutes ago
}
```
