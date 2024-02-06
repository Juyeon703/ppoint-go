package logue

import (
	"fmt"
	"io"
	"os"
)

type FileLogueOption struct {
	ModuleName string

	// append True: 기존 로그에 추가, False: 로그 overwrite
	Append bool

	// 로그 파일이 생성되는 기본 경로(../log)가 아닌 곳에
	// 로그 파일을 생성하고 싶을때 설정한다
	BasePath string

	Descriptor *os.File
}

func (f *FileLogueOption) Setup() (io.Writer, error) {
	var w io.Writer
	var descriptor *os.File
	var err error
	var logBase string

	if len(f.ModuleName) <= 0 {
		return nil, fmt.Errorf("ModuleName is empty")
	}

	if len(f.BasePath) == 0 {
		logBase = "../log"
	} else {
		logBase = f.BasePath
	}

	if _, e := os.Stat(logBase); os.IsNotExist(e) {
		if er_md := os.Mkdir(logBase, 0755); er_md != nil {
			fmt.Printf("Logger: Mkdir %s", er_md)
			return nil, er_md
		}
	}
	logPath := fmt.Sprintf("%s/%s_%s.log", logBase, f.ModuleName, logCreateDate())

	if f.Append {
		descriptor, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	} else {
		descriptor, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	}

	f.Descriptor = descriptor

	if err != nil {
		fmt.Printf("Logger: Open File Error(%s),\n", err.Error())
		w = io.Writer(os.Stdout)
	}
	w = io.MultiWriter(descriptor, os.Stdout)

	return w, err
}

/////////////////EOF
