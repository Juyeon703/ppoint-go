package logue

import (
	"io"
)

type LogueOptioner interface {
	Setup() (io.Writer, error)
}

//////////////EOF
