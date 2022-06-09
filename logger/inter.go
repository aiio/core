package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	infoLevel  = "info"
	errorLevel = "error"
	debugLevel = "debug"
)

func NewLog(dir, prefix, level string, maxSize, maxBackups, maxAge int) (*zap.Logger, func(), error) {

	if maxSize < 0 {
		maxSize = 0
	}
	if maxBackups < 0 {
		maxBackups = 0
	}
	if maxAge < 0 {
		maxAge = 0
	}

	infoLog := &lumberjack.Logger{
		Filename:   getFilePath(dir, prefix, infoLevel),
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  true,
		Compress:   false,
	}

	errLog := &lumberjack.Logger{
		Filename:   getFilePath(dir, prefix, errorLevel),
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  true,
		Compress:   false,
	}

	cronLogger := &lumberLog{
		dir:     dir,
		prefix:  prefix,
		level:   level,
		infoLog: infoLog,
		errLog:  errLog,
	}

	return cronLogger.getLoggerWithCron(), func() { cronLogger.StopLogger() }, nil
}

func (l *lumberLog) rotate() {
	if l.infoLog != nil {
		l.infoLog.Rotate()
	}
	if l.errLog != nil {
		l.errLog.Rotate()
	}
}

type lumberLog struct {
	dir     string
	prefix  string
	level   string
	infoLog *lumberjack.Logger
	errLog  *lumberjack.Logger
}

func (l *lumberLog) StopLogger() {
	l.infoLog.Close()
	l.errLog.Close()
}

func (l *lumberLog) getLoggerWithCron() *zap.Logger {
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	lowLevel := zap.InfoLevel
	if strings.EqualFold(l.level, debugLevel) {
		lowLevel = zap.DebugLevel
	}

	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= lowLevel
	})

	prodEncoder := zap.NewProductionEncoderConfig()
	prodEncoder.EncodeTime = zapcore.ISO8601TimeEncoder
	prodEncoder.CallerKey = ""

	highCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), zapcore.AddSync(l.errLog), highPriority)
	lowCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), zapcore.AddSync(l.infoLog), lowPriority)

	return zap.New(zapcore.NewTee(highCore, lowCore))
}
