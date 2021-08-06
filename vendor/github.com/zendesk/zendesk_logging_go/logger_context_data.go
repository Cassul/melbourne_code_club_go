package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type contextKey string

const loggerCtxKey contextKey = "zdlogging_logger_ctx"

// loggerContextData wraps a map of fields to thread through via context.Context
type loggerContextData struct {
	fields Fields
}

func extractLoggerDataWithFieldExtractors(ctx context.Context) loggerContextData {
	newLoggerData := loggerContextData{
		fields: map[string]interface{}{},
	}

	existingLoggerData, ok := ctx.Value(loggerCtxKey).(loggerContextData)
	if ok {
		for key, value := range existingLoggerData.fields {
			newLoggerData.fields[key] = value
		}
	}

	for _, fieldExtractorFunc := range globalConfig.FieldExtractorFuncs {
		for key, value := range fieldExtractorFunc(ctx) {
			newLoggerData.fields[key] = value
		}
	}

	return newLoggerData
}

// takeReferenceAsLogrusEntry converts a loggerContextData to a logrus.Fields
func (data loggerContextData) takeReferenceAsLogrusEntry() *LogEntry {
	return logrus.WithFields(logrus.Fields(data.fields))
}
