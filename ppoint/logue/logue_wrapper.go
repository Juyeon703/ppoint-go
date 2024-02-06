package logue

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var depth = 2

type Logbook struct {
	logue       *Logue
	logOption   *FileLogueOption
	logLevel    Level
	chgLogLevel Level
	strLogDate  string
	mutex       sync.Mutex
}

func logCreateDate() string {
	return time.Now().Format("20060102")
}

func Setup(moduleLogPath, moduleName string, moduleLogLevel int) (*Logbook, error) {
	var logbook = &Logbook{}
	var log *Logue = nil
	var err error = nil
	logbook.strLogDate = logCreateDate()

	if len(moduleName) <= 0 {
		return nil, fmt.Errorf("MODULE NAME IS EMPTY")
	}

	if _, err = os.Stat(moduleLogPath); os.IsNotExist(err) {
		if err = os.Mkdir(moduleLogPath, 0755); err != nil {
			fmt.Printf("Logger: Mkdir %s", err)
			return nil, err
		}
	}

	option := &FileLogueOption{ModuleName: moduleName, Append: true, BasePath: moduleLogPath}
	logbook.logOption = option
	logbook.logLevel = Level(moduleLogLevel)
	logbook.chgLogLevel = Level(moduleLogLevel)

	if log, err = New(logbook.logLevel, 2, logbook.logOption); err != nil {
		fmt.Printf("[FATAL] %s\n", err.Error())
		return nil, fmt.Errorf("[FATAL] %s\n", err.Error())
	}
	logbook.logue = log

	return logbook, nil
}

func getLogue(logbook *Logbook) *Logue {
	var log *Logue = nil
	var err error = nil
	var nowDate = logCreateDate()

	if logbook.logue == nil {
		if log, err = New(logbook.logLevel, 2, logbook.logOption); err != nil {
			fmt.Printf("[FATAL] %s\n", err.Error())
			return logbook.logue
		}
		logbook.strLogDate = logCreateDate()
		logbook.logue = log
		return logbook.logue
	}

	//loglevel change
	if logbook.logLevel != logbook.chgLogLevel {
		logbook.logLevel = logbook.chgLogLevel
		if log, err = New(logbook.logLevel, 2, logbook.logOption); err != nil {
			fmt.Printf("[FATAL] %s\n", err.Error())
			return logbook.logue
		}
		logbook.strLogDate = logCreateDate()
		logbook.logue = log
		return logbook.logue
	}

	if logbook.strLogDate == nowDate {
		return logbook.logue
	}

	if logbook.logue != nil {
		logbook.logOption.Descriptor.Close()
		logbook.logue = nil
	}

	if log, err = New(logbook.logLevel, 2, logbook.logOption); err != nil {
		fmt.Printf("[FATAL] %s\n", err.Error())
		return logbook.logue
	}
	logbook.strLogDate = logCreateDate()
	logbook.logue = log

	return logbook.logue
}

func (logbook *Logbook) ChangeLogLevel(strLogLevel string) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if strings.EqualFold(strLogLevel, "DEBUG") {
		logbook.chgLogLevel = 5
	} else if strings.EqualFold(strLogLevel, "INFO") {
		logbook.chgLogLevel = 4
	} else if strings.EqualFold(strLogLevel, "WARN") {
		logbook.chgLogLevel = 3
	} else if strings.EqualFold(strLogLevel, "ERROR") {
		logbook.chgLogLevel = 2
	} else if strings.EqualFold(strLogLevel, "FATAL") {
		logbook.chgLogLevel = 1
	} else {
		logbook.Warn("FAIL TO CHANGE LOG LEVEL")
	}

	logbook.logue = getLogue(logbook)
}

func (logbook *Logbook) CloseLogFile() error {
	var err error = nil
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logue != nil {
		if err = logbook.logOption.Descriptor.Close(); err != nil {
			return err
		}
		logbook.logue = nil
	}

	return nil
}

func (logbook *Logbook) Debug(v ...interface{}) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logLevel < DEBUG {
		return
	}

	logbook.logue = getLogue(logbook)
	v = append([]interface{}{trace()}, v)
	logbook.logue.debug.Output(depth, fmt.Sprintln(v...))
}

func (logbook *Logbook) Info(v ...interface{}) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logLevel < INFO {
		return
	}

	logbook.logue = getLogue(logbook)
	v = append([]interface{}{trace()}, v)
	logbook.logue.info.Output(depth, fmt.Sprintln(v...))
}

func (logbook *Logbook) Warn(v ...interface{}) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logLevel < WARNING {
		return
	}

	logbook.logue = getLogue(logbook)
	v = append([]interface{}{trace()}, v)
	logbook.logue.warning.Output(depth, fmt.Sprintln(v...))
}

func (logbook *Logbook) Error(v ...interface{}) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logLevel < WARNING {
		return
	}

	logbook.logue = getLogue(logbook)
	v = append([]interface{}{trace()}, v)
	logbook.logue.error.Output(depth, fmt.Sprintln(v...))
}

func (logbook *Logbook) Fatal(v ...interface{}) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logLevel < FATAL {
		return
	}

	logbook.logue = getLogue(logbook)
	v = append([]interface{}{trace()}, v)
	logbook.logue.fatal.Output(depth, fmt.Sprintln(v...))
}

func (logbook *Logbook) Debugf(format string, v ...interface{}) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logLevel < DEBUG {
		return
	}

	logbook.logue = getLogue(logbook)
	strTrace := trace() + " "
	logbook.logue.debug.Output(depth, strTrace+fmt.Sprintf(format, v...))
}

func (logbook *Logbook) Infof(format string, v ...interface{}) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logLevel < INFO {
		return
	}

	logbook.logue = getLogue(logbook)
	strTrace := trace() + " "
	logbook.logue.info.Output(depth, strTrace+fmt.Sprintf(format, v...))
}

func (logbook *Logbook) Warnf(format string, v ...interface{}) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logLevel < WARNING {
		return
	}

	logbook.logue = getLogue(logbook)
	strTrace := trace() + " "
	logbook.logue.warning.Output(depth, strTrace+fmt.Sprintf(format, v...))
}

func (logbook *Logbook) Errorf(format string, v ...interface{}) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logLevel < ERROR {
		return
	}

	logbook.logue = getLogue(logbook)
	strTrace := trace() + " "
	logbook.logue.error.Output(depth, strTrace+fmt.Sprintf(format, v...))
}

func (logbook *Logbook) Fatalf(format string, v ...interface{}) {
	logbook.mutex.Lock()
	defer logbook.mutex.Unlock()

	if logbook.logLevel < FATAL {
		return
	}

	logbook.logue = getLogue(logbook)
	strTrace := trace() + " "
	logbook.logue.fatal.Output(depth, strTrace+fmt.Sprintf(format, v...))
}

func trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(depth+1, pc)
	f := runtime.FuncForPC(pc[0])
	//file, line := f.FileLine(pc[0])
	//return fmt.Sprintf("%s:%d %s\n", file, line, f.Name())

	s := strings.Split(fmt.Sprintf("%s", f.Name()), string('/')) //string(os.PathSeparator)
	return s[len(s)-1]
}
