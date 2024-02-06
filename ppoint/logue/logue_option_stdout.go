package logue

import (
	"io"
	"os"
)

type StdoutLogueOption struct {
}

func (s *StdoutLogueOption) Setup() (io.Writer, error) {
	w := io.Writer(os.Stdout)
	return w, nil
}

//////////EOF
