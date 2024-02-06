package logue

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Rolling struct {
	BasePath string
	MaxSize  int64
	Compress bool

	file *os.File
	size int64
	name string
}

var (
	currentTime = time.Now
)

const backupTimeFormat = "2006-01-02T15-04-05.000"

// file이 존재하면 append 할 수 있도록 파일을 open한다.
// 없으면 새로 open한다
func (r *Rolling) OpenFile(name string) (*os.File, error) {
	r.name = name
	info, err := os.Stat(name)
	if os.IsExist(err) {
		if info.Size() >= r.MaxSize {
			f, err := r.rotate()
			r.size = 0
			r.file = f
			return f, err
		} else {
			descriptor, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
			r.size = info.Size()
			r.file = descriptor
			return descriptor, err
		}
	}

	descriptor, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	r.size = 0
	r.file = descriptor
	return descriptor, err

}

func (r *Rolling) rotate() (*os.File, error) {
	//backup name
	dir := filepath.Dir(r.name)
	filename := filepath.Base(r.name)
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]

	t := currentTime()

	timestamp := t.Format(backupTimeFormat)
	backupName := filepath.Join(dir, fmt.Sprintf("%s-%s%s", prefix, timestamp, ext))

	if err := os.Rename(r.name, backupName); err != nil {
		return nil, fmt.Errorf("can't rename log file: %s", err)
	}

	//compress
	if r.Compress == true {
		bdescr, _ := os.Open(backupName)
		compressName := filepath.Join(dir, fmt.Sprintf("%s-%s%s", prefix, timestamp, ".gz"))
		gzf, _ := os.OpenFile(compressName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
		gz := gzip.NewWriter(gzf)

		io.Copy(gz, bdescr)
		gz.Close()
		gzf.Close()
		bdescr.Close()

		os.Remove(backupName)

	}
	descriptor, err := os.OpenFile(r.name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)

	r.size = 0
	return descriptor, err

}

func (r *Rolling) Write(p []byte) (n int, err error) {
	//var err error

	ln := int64(len(p))
	if ln > r.MaxSize {
		return 0, fmt.Errorf("exceeds maximum size %d byte", r.MaxSize)
	}

	if r.size+ln > r.MaxSize {
		if r.file, err = r.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = r.file.Write(p)
	r.size += int64(n)

	return n, err
}

type FileRollingLogueOption struct {
	ModuleName string
	BasePath   string

	// byte
	MaxSize int64

	// 백업 로그 파일 압축 여부
	Compress bool
}

func (f *FileRollingLogueOption) Setup() (io.Writer, error) {
	//var w io.Writer
	//var descriptor *os.File
	var err error
	var logBase string

	rolling := &Rolling{BasePath: f.BasePath, MaxSize: f.MaxSize, Compress: f.Compress}

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
	logPath := fmt.Sprintf("%s/%s.log", logBase, f.ModuleName)

	_, err = rolling.OpenFile(logPath)

	return rolling, err

}

///////////EOS
