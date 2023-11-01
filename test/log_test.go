package logtest

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	goLog "github.com/lucas-simao/golog"
	"github.com/stretchr/testify/suite"
)

type goLogData struct {
	Service     string                 `json:"service"`
	File        string                 `json:"file"`
	Func        string                 `json:"func"`
	Message     interface{}            `json:"message"`
	Severity    string                 `json:"severity"`
	Timestamp   string                 `json:"timestamp"`
	ExtraFields map[string]interface{} `json:"extraFields,omitempty"`
}

type LogTestSuite struct {
	suite.Suite
	GOLOG_SERVICE string
}

func TestLogTestSuite(t *testing.T) {
	suite.Run(t, new(LogTestSuite))
}

func (suite *LogTestSuite) SetupSuite() {
	suite.GOLOG_SERVICE = "TestLogTestSuite"
	os.Setenv("GOLOG_SERVICE", suite.GOLOG_SERVICE)
}

func (suite *LogTestSuite) TestDefault() {
	tests := map[string]struct {
		message     string
		extraFields map[string]interface{}
	}{
		"test Default": {
			message: "test Default",
		},
		"test Default with extra fields": {
			message: "test Default with extra fields",
			extraFields: map[string]interface{}{
				"method": "Default",
			},
		},
	}

	for name, t := range tests {
		suite.Run(name, func() {
			output := goLog.Default(t.message, t.extraFields)

			data, err := unmarshalData(output)
			suite.NoError(err)

			suite.Equal(t.message, fmt.Sprintf("%s", data.Message))
			suite.Equal(t.extraFields, data.ExtraFields)
			suite.Equal(goLog.SeverityDefault, data.Severity)
		})
	}
}

func (suite *LogTestSuite) TestDebug() {
	tests := map[string]struct {
		message     string
		extraFields map[string]interface{}
	}{
		"test Debug": {
			message: "test Debug",
		},
		"test Debug with extra fields": {
			message: "test Debug with extra fields",
			extraFields: map[string]interface{}{
				"method": "Debug",
			},
		},
	}

	for name, t := range tests {
		suite.Run(name, func() {
			output := goLog.Debug(t.message, t.extraFields)

			data, err := unmarshalData(output)
			suite.NoError(err)

			suite.Equal(t.message, fmt.Sprintf("%s", data.Message))
			suite.Equal(t.extraFields, data.ExtraFields)
			suite.Equal(goLog.SeverityDebug, data.Severity)
		})
	}
}

func (suite *LogTestSuite) TestInfo() {
	tests := map[string]struct {
		message     string
		extraFields map[string]interface{}
	}{
		"test Info": {
			message: "test Info",
		},
		"test Info with extra fields": {
			message: "test Info with extra fields",
			extraFields: map[string]interface{}{
				"method": "Info",
			},
		},
	}

	for name, t := range tests {
		suite.Run(name, func() {
			output := goLog.Info(t.message, t.extraFields)

			data, err := unmarshalData(output)
			suite.NoError(err)

			suite.Equal(t.message, fmt.Sprintf("%s", data.Message))
			suite.Equal(t.extraFields, data.ExtraFields)
			suite.Equal(goLog.SeverityInfo, data.Severity)
		})
	}
}

func (suite *LogTestSuite) TestNotice() {
	tests := map[string]struct {
		message     string
		extraFields map[string]interface{}
	}{
		"test Notice": {
			message: "test Notice",
		},
		"test Notice with extra fields": {
			message: "test Notice with extra fields",
			extraFields: map[string]interface{}{
				"method": "Notice",
			},
		},
	}

	for name, t := range tests {
		suite.Run(name, func() {
			output := goLog.Notice(t.message, t.extraFields)

			data, err := unmarshalData(output)
			suite.NoError(err)

			suite.Equal(t.message, fmt.Sprintf("%s", data.Message))
			suite.Equal(t.extraFields, data.ExtraFields)
			suite.Equal(goLog.SeverityNotice, data.Severity)
		})
	}
}

func (suite *LogTestSuite) TestWarning() {
	tests := map[string]struct {
		message     string
		extraFields map[string]interface{}
	}{
		"test Warning": {
			message: "test Warning",
		},
		"test Warning with extra fields": {
			message: "test Warning with extra fields",
			extraFields: map[string]interface{}{
				"method": "Warning",
			},
		},
	}

	for name, t := range tests {
		suite.Run(name, func() {
			output := goLog.Warning(t.message, t.extraFields)

			data, err := unmarshalData(output)
			suite.NoError(err)

			suite.Equal(t.message, fmt.Sprintf("%s", data.Message))
			suite.Equal(t.extraFields, data.ExtraFields)
			suite.Equal(goLog.SeverityWarning, data.Severity)
		})
	}
}

func (suite *LogTestSuite) TestError() {
	tests := map[string]struct {
		message     string
		extraFields map[string]interface{}
	}{
		"test Error": {
			message: "test Error",
		},
		"test Error with extra fields": {
			message: "test Error with extra fields",
			extraFields: map[string]interface{}{
				"method": "Error",
			},
		},
	}

	for name, t := range tests {
		suite.Run(name, func() {
			output := goLog.Error(t.message, t.extraFields)

			data, err := unmarshalData(output)
			suite.NoError(err)

			suite.Equal(t.message, fmt.Sprintf("%s", data.Message))
			suite.Equal(t.extraFields, data.ExtraFields)
			suite.Equal(goLog.SeverityError, data.Severity)
		})
	}
}

func (suite *LogTestSuite) TestAlert() {
	tests := map[string]struct {
		message     string
		extraFields map[string]interface{}
	}{
		"test Alert": {
			message: "test Alert",
		},
		"test Alert with extra fields": {
			message: "test Alert with extra fields",
			extraFields: map[string]interface{}{
				"method": "Alert",
			},
		},
	}

	for name, t := range tests {
		suite.Run(name, func() {
			output := goLog.Alert(t.message, t.extraFields)

			data, err := unmarshalData(output)
			suite.NoError(err)

			suite.Equal(t.message, fmt.Sprintf("%s", data.Message))
			suite.Equal(t.extraFields, data.ExtraFields)
			suite.Equal(goLog.SeverityAlert, data.Severity)
		})
	}
}

func (suite *LogTestSuite) TestEmergency() {
	tests := map[string]struct {
		message     string
		extraFields map[string]interface{}
	}{
		"test Emergency": {
			message: "test Emergency",
		},
		"test Emergency with extra fields": {
			message: "test Emergency with extra fields",
			extraFields: map[string]interface{}{
				"method": "Emergency",
			},
		},
	}

	for name, t := range tests {
		suite.Run(name, func() {
			output := goLog.Emergency(t.message, t.extraFields)

			data, err := unmarshalData(output)
			suite.NoError(err)

			suite.Equal(t.message, fmt.Sprintf("%s", data.Message))
			suite.Equal(t.extraFields, data.ExtraFields)
			suite.Equal(goLog.SeverityEmergency, data.Severity)
		})
	}
}

func (suite *LogTestSuite) TestCritical() {
	tests := map[string]struct {
		message     string
		extraFields map[string]interface{}
	}{
		"test Critical": {
			message: "test Critical",
		},
		"test Critical with extra fields": {
			message: "test Critical with extra fields",
			extraFields: map[string]interface{}{
				"method": "Critical",
			},
		},
	}

	for name, t := range tests {
		suite.Run(name, func() {
			var output []byte

			defer func() {
				if err := recover(); err != nil {
					output = err.([]byte)
				}

				data, err := unmarshalData(output)
				suite.NoError(err)

				suite.Equal(t.message, fmt.Sprintf("%s", data.Message))
				suite.Equal(t.extraFields, data.ExtraFields)
				suite.Equal(goLog.SeverityCritical, data.Severity)

			}()

			goLog.Critical(t.message, t.extraFields)
		})
	}
}

func unmarshalData(b []byte) (goLogData, error) {
	var resp goLogData

	err := json.Unmarshal(b, &resp)

	return resp, err
}
