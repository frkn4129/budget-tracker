package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	log *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	raw, err := zap.NewDevelopment()
	if err != nil {
		panic("Zap logger başlatılamadı: " + err.Error())
	}
	return &ZapLogger{
		log: raw.Sugar(),
	}
}

func (z *ZapLogger) Info(msg string, fields ...interface{}) {
	z.log.Infow(msg, fields...)
}

func (z *ZapLogger) Error(msg string, fields ...interface{}) {
	z.log.Errorw(msg, fields...)
}

func (z *ZapLogger) Fatal(msg string, fields ...interface{}) {
	z.log.Fatalw(msg, fields...)
}
