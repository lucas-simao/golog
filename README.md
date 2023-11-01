# goLog

goLog is a structured logger for Go (golang), this service use the logical to work with [GCP Log Severity](https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity)

### Add go module

```
go get "github.com/lucas-simao/golog"
```

```go
package main

import (
 "context"
 "os"

  "github.com/lucas-simao/golog"
)

func main() {
 // Create on .env/Dockerfile the GOLOG_SERVICE with project name to goLog use on service field
 os.Setenv("GOLOG_SERVICE", "service-name")

 //The available methods has is:
 goLog.Debug("debug")
 goLog.Info("info")
 goLog.Notice("notice")
 goLog.Warning("warning")
 goLog.Error("error")
 goLog.Alert("alert")
 goLog.Emergency("alert")

 //If you need extra fields, you can add map[string]interface{}{} with extras fields
 //This extras fields show inside of "extraFields":{"partnerId":"74dadd69-df69-4c12-a908-267a5182c3a4", "isContact":true}...
   //All methods is possible use the extra fields
 goLog.Default("with extra fields", map[string]interface{}{
  "partnerId": "74dadd69-df69-4c12-a908-267a5182c3a4",
  "isContact": true,
 })
}
```

```json
{"service":"service-name","userId":null,"file":"goLog/teste/main.go:19","func":"main.go","message":"debug","severity":"debug","timestamp":"2023-01-10T15:19:32.044-04:00"}

{"service":"service-name","userId":null,"file":"goLog/teste/main.go:20","func":"main.go","message":"info","severity":"info","timestamp":"2023-01-10T15:19:32.044-04:00"}

{"service":"service-name","userId":null,"file":"goLog/teste/main.go:21","func":"main.go","message":"notice","severity":"notice","timestamp":"2023-01-10T15:19:32.044-04:00"}

{"service":"service-name","userId":null,"file":"goLog/teste/main.go:22","func":"main.go","message":"warning","severity":"warning","timestamp":"2023-01-10T15:19:32.044-04:00"}

{"service":"service-name","userId":null,"file":"goLog/teste/main.go:23","func":"main.go","message":"error","severity":"error","timestamp":"2023-01-10T15:19:32.044-04:00"}

{"service":"service-name","userId":null,"file":"goLog/teste/main.go:24","func":"main.go","message":"alert","severity":"alert","timestamp":"2023-01-10T15:19:32.044-04:00"}

{"service":"service-name","userId":null,"file":"goLog/teste/main.go:25","func":"main.go","message":"alert","severity":"emergency","timestamp":"2023-01-10T15:19:32.044-04:00"}

{"service":"service-name","userId":null,"file":"goLog/teste/main.go:29","func":"main.go","message":"with extra fields","severity":"default","timestamp":"2023-01-10T15:19:32044-04:00","extraFields":{"isContact":true,"partnerId":"74dadd69-df69-4c12-a908-267a5182c3a4"}
```

### kill the application, its use the panic(message)

```go
goLog.Critical("kill the application")
```
