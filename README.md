# Simple logging in Go.

```go
package main

import "github.com/gaemma/logging"

func main() {
   logger := logging.NewNopLogger()
   logger.Info("the counter is %d.", 1)
}
```

`NewNopLogger` creates a logger does nothing when calling the logger method.

```go
package main

import "github.com/gaemma/logging"

func main() {
   logger := logging.NewSimpleLogger()
   logger.Info("the counter is %d.", 1)
}
```

`NewSimpleLogger` creates a logger using the stdlib `log`.