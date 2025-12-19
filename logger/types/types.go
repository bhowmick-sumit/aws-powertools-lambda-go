package types

import (
	"io"
)

type LogConfig struct {
	Writer         io.Writer
	IncludeContext bool
	Properties     map[string]string
}
