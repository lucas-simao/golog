package log

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const RFC3339MilliFormat = "2006-01-02T15:04:05.999Z07:00"

const (
	SeverityDefault   string = "default"
	SeverityDebug     string = "debug"
	SeverityInfo      string = "info"
	SeverityNotice    string = "notice"
	SeverityWarning   string = "warning"
	SeverityError     string = "error"
	SeverityCritical  string = "critical"
	SeverityAlert     string = "alert"
	SeverityEmergency string = "emergency"
)

type goLogService struct {
	Service     string                 `json:"service"`
	File        string                 `json:"file"`
	Func        string                 `json:"func"`
	Message     interface{}            `json:"message"`
	Severity    string                 `json:"severity"`
	Timestamp   string                 `json:"timestamp"`
	ExtraFields map[string]interface{} `json:"extraFields,omitempty"`
}

func Default(message interface{}, extraFields ...map[string]interface{}) []byte {
	return printLog(SeverityDefault, message, extraFields...)
}
func Debug(message interface{}, extraFields ...map[string]interface{}) []byte {
	return printLog(SeverityDebug, message, extraFields...)
}
func Info(message interface{}, extraFields ...map[string]interface{}) []byte {
	return printLog(SeverityInfo, message, extraFields...)
}
func Notice(message interface{}, extraFields ...map[string]interface{}) []byte {
	return printLog(SeverityNotice, message, extraFields...)
}
func Warning(message interface{}, extraFields ...map[string]interface{}) []byte {
	return printLog(SeverityWarning, message, extraFields...)
}
func Error(message interface{}, extraFields ...map[string]interface{}) []byte {
	return printLog(SeverityError, message, extraFields...)
}
func Critical(message interface{}, extraFields ...map[string]interface{}) []byte {
	defer panic(printLog(SeverityCritical, message, extraFields...))
	return printLog(SeverityCritical, message, extraFields...)
}
func Alert(message interface{}, extraFields ...map[string]interface{}) []byte {
	return printLog(SeverityAlert, message, extraFields...)
}
func Emergency(message interface{}, extraFields ...map[string]interface{}) []byte {
	return printLog(SeverityEmergency, message, extraFields...)
}

func getFuncLine() (funcName, fileName string) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return
	}

	fSprint := strings.Split(file, "/")

	return fSprint[len(fSprint)-1], fmt.Sprintf("%s:%d", file, line)
}

func printLog(severity string, message interface{}, extraFields ...map[string]interface{}) []byte {
	funcName, line := getFuncLine()

	var data = goLogService{
		File:      line,
		Func:      funcName,
		Timestamp: time.Now().Local().Format(RFC3339MilliFormat),
		Severity:  severity,
		Message:   fmt.Sprintf("%v", message),
		Service:   os.Getenv("GOLOG_SERVICE"),
	}

	if len(extraFields) > 0 {
		data.ExtraFields = extraFields[0]
	}

	b, _ := json.Marshal(data)

	fmt.Fprintln(os.Stdout, string(b))

	return b
}
