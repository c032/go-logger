package logger

import (
	"os"
)

var Default Logger = NewJSON(os.Stdout)
