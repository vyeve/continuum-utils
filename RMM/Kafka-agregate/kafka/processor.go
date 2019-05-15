package kafka

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/ContinuumLLC/platform-common-lib/src/messaging"
	lj "gopkg.in/natefinch/lumberjack.v2"
)

const (
	// logFileName         = "kafka.log"
	logJSONFileName     = "kafkaJson.log"
	maxLogFileSizeInMB  = 10
	maxLogFileBackups   = 3
	maxLogFileAgeInDays = 30
)

type collector struct {
	// writer     io.Writer
	jsonWriter io.Writer
}

func NewWriter() MessageProcessor {
	mp := new(collector)
	mp.jsonWriter = &lj.Logger{
		LocalTime:  true,
		Compress:   true,
		MaxSize:    maxLogFileSizeInMB,
		MaxBackups: maxLogFileBackups,
		MaxAge:     maxLogFileAgeInDays,
		Filename:   logJSONFileName,
	}

	return mp
}

type Envelope struct {
	Header  messaging.Header
	Topic   string
	Message string
}

func (e Envelope) GetBytes() []byte {
	return []byte(fmt.Sprintf("%v\n", e))
}

func (e Envelope) GetJson() []byte {
	out, err := json.Marshal(&e)
	if err != nil {
		return nil
	}
	out = append(out, '\n')
	return out
}

func (c collector) Process(message *messaging.Message) {
	en := Envelope{
		Message: message.Envelope.Message,
		Header:  message.Envelope.Header,
		Topic:   message.Envelope.Topic,
	}
	if len(en.Header) != 0 {
		var err error
		_, err = c.jsonWriter.Write(en.GetJson())

		if err != nil {
			log.Println(err)
		}
	}
}
