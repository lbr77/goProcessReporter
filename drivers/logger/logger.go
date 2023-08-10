package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	file, err := os.OpenFile(fmt.Sprintf("%s\\log.log", filepath.Dir(os.Args[0])), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		Log.SetOutput(os.Stdout)
	}
}
