package writer

import (
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/models"
)

const (
	logFileName         = "endpoints.txt"
	maxLogFileSizeInMB  = 100
	maxLogFileBackups   = 3
	maxLogFileAgeInDays = 30
)

// Writer ...
type Writer interface {
	Write(models.AssetCollection) error
	WriteAll([]models.AssetCollection) error
}

// Client ...
var Client Writer

type writer struct {
	w io.Writer
}

// Load ...
func Load() (err error) {
	w := newWriter()
	Client = w
	return nil
}

func (w writer) Write(asset models.AssetCollection) (err error) {
	p, err := json.Marshal(&asset)
	if err != nil {
		return err
	}
	p = append(p, '\n')
	_, err = w.w.Write(p)
	if err != nil {
		return err
	}
	return nil
}

func (w writer) WriteAll(assets []models.AssetCollection) (err error) {
	for i, asset := range assets {
		errOne := w.Write(asset)
		if errOne != nil {
			if err == nil {
				err = errOne
			} else {
				err = fmt.Errorf("%s; %s", err, errOne)
			}
		}
		if i != len(assets)-1 {
			w.w.Write([]byte{'\n'})
		}
	}
	return
}

func newWriter() Writer {
	wr := lumberjack.Logger{
		LocalTime:  true,
		Compress:   true,
		MaxSize:    maxLogFileSizeInMB,
		MaxBackups: maxLogFileBackups,
		MaxAge:     maxLogFileAgeInDays,
		Filename:   logFileName,
	}
	return writer{w: &wr}
}
