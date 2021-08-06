package logger

import (
	"context"
	"io"
	"os"
)

// Option type, for use in the Configure func
type Option func(*Config)

// FieldExtractorFunc is used to configure field extractors when setting up logger
type FieldExtractorFunc func(context.Context) map[string]interface{}

// Config contains configuration properties for the logger such as the log level, output format and io output
type Config struct {
	Level               string `env:"LOG_LEVEL" default:"info"`
	Format              string `env:"LOG_FORMAT" default:"json"`
	PrettyPrintJSON     bool
	Output              io.Writer
	DefaultFields       Fields
	FieldExtractorFuncs []FieldExtractorFunc
}

var defaultConfig = Config{
	Level:               LevelInfo,
	Format:              FormatJSON,
	PrettyPrintJSON:     false,
	Output:              os.Stdout,
	DefaultFields:       Fields{},
	FieldExtractorFuncs: []FieldExtractorFunc{},
}

// Level func, for use in the Setup functions to set the log level
// Defaults to LevelInfo
func Level(level string) Option {
	return func(config *Config) {
		if level == "" {
			config.Level = LevelInfo
		} else {
			config.Level = level
		}
	}
}

// Format func, for use in the Configure func to set the log output format
// Defaults to FormatJSON
func Format(format string) Option {
	return func(config *Config) {
		if format == "" {
			config.Format = FormatJSON
		} else {
			config.Format = format
		}
	}
}

// PrettyPrintJSON func, for use in the Configure func to toggle pretty json print
// Defaults to false
func PrettyPrintJSON(pretty bool) Option {
	return func(config *Config) {
		config.PrettyPrintJSON = pretty
	}
}

// Output func, for use in the Configure func to set the log output stream
// Defaults to os.Stdout
func Output(out io.Writer) Option {
	return func(config *Config) {
		if out == nil {
			config.Output = os.Stdout
		} else {
			config.Output = out
		}
	}
}

// DefaultFields func, for use in the Configure func to set default fields for every log
// Defaults to Fields{}
func DefaultFields(fields Fields) Option {
	return func(config *Config) {
		if fields.empty() {
			config.DefaultFields = defaultConfig.DefaultFields
		} else {
			config.DefaultFields = fields
		}
	}
}

// FieldExtractorFuncs func, for use in the Configure func to record the functions we want
// to call each time we prepare the fields for a log line
func FieldExtractorFuncs(funcs ...FieldExtractorFunc) Option {
	return func(config *Config) {
		config.FieldExtractorFuncs = funcs
	}
}

func optionsFromConfig(config Config) []Option {
	return []Option{
		Level(config.Level),
		Format(config.Format),
		PrettyPrintJSON(config.PrettyPrintJSON),
		Output(config.Output),
		DefaultFields(config.DefaultFields),
		FieldExtractorFuncs(config.FieldExtractorFuncs...),
	}
}
