// log
package yklog

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type yklogger struct {
	sysfile *os.File
	syslog  *log.Logger
}

func create(save2file int) (*yklogger, error) {
	now := time.Now()
	filename := fmt.Sprintf("%d%02d%02d_%02d_%02d_%02d.log",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second())

	logger := new(yklogger)
	logger.sysfile = os.Stdout
	if save2file != 0 {
		file, err := os.Create(path.Join("", filename))
		if err != nil {
			return nil, err
		}
		logger.sysfile = file
	}
	logger.syslog = log.New(logger.sysfile, "", log.LstdFlags)
	return logger, nil
}

var locallog, _ = create(0) //是否存入文件的开关

func (yklog *yklogger) release() {
	if yklog.sysfile != nil {
		yklog.sysfile.Close()
	}

	yklog.syslog = nil
	yklog.sysfile = nil
}

func Logout(format string, a ...interface{}) {
	locallog.syslog.Output(3, fmt.Sprintf(format, a...))

}
