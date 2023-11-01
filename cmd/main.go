package main

import (
	"errors"

	"github.com/google/uuid"
	goLog "github.com/lucas-simao/golog"
)

func main() {
	goLog.Info("teste", map[string]interface{}{
		"uuid":    uuid.New(),
		"uuidNil": uuid.Nil,
		"true":    true,
		"false":   false,
		"string":  "oi",
		"int":     1,
		"int16":   int16(1),
		"int32":   int32(1),
		"int64":   int64(1),
		"float32": float32(1),
		"float64": float64(1),
		"nil":     nil,
		"err":     errors.New("error new"),
	})
}
