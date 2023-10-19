# Logger

## A custom logger written in Go


## Usage

```go
package main

import "github.com/sirgallo/logger"


func main() {
  // init the logger
  cLog := logger.NewCustomLog("Custom Log")

  // log for each log level
  cLog.Debug("this is a debug message.")
  cLog.Error("this is an error message.")
  cLog.Info("this is an info message.")
  cLog.Warn("this is a warning message.")
  cLog.Fatal("this is an error message that will terminate your program.")
}
```


## Tests

`logger`
```bash
go test -v ./tests
```