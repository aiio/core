package logger

import (
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

func init() {

	// 构造EncoderConfig
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, //FullCallerEncoder ShortCallerEncoder
	}

	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	lowLevel := zap.DebugLevel
	//if strings.EqualFold(l.level, debugLevel) {
	//	lowLevel = zap.DebugLevel
	//}

	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= lowLevel
	})

	highFile := &lumberjack.Logger{ // concurrent-safed
		Filename:   getFilePath("logs", "error2", "error"), // 文件路径
		MaxSize:    128,                                    // 单个文件最大尺寸，默认单位 M
		MaxBackups: 300,                                    // 最多保留 300 个备份
		MaxAge:     30,                                     // 最大时间，默认单位 day
		LocalTime:  true,                                   // 使用本地时间
		Compress:   true,                                   // 是否压缩 disabled by default
	}
	infoFile := &lumberjack.Logger{ // concurrent-safed
		Filename:   getFilePath("logs", "info2", "info"), // 文件路径
		MaxSize:    128,                                  // 单个文件最大尺寸，默认单位 M
		MaxBackups: 300,                                  // 最多保留 300 个备份
		MaxAge:     30,                                   // 最大时间，默认单位 day
		LocalTime:  true,                                 // 使用本地时间
		Compress:   true,                                 // 是否压缩 disabled by default
	}

	highCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(highFile), highPriority)
	lowCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(infoFile), lowPriority)

	Logger = zap.New(zapcore.NewTee(highCore, lowCore))

	// 然后是SugarLogger
	SugarLogger = Logger.Sugar()
}

func Debug(args ...interface{}) {
	SugarLogger.Debug(args...)
}

func Info(args ...interface{}) {
	SugarLogger.Info(args...)
}

func Warn(args ...interface{}) {
	SugarLogger.Warn(args...)
}

func Error(args ...interface{}) {
	SugarLogger.Error(args...)
}
func DPanic(args ...interface{}) {
	SugarLogger.DPanic(args...)
}
func Panic(args ...interface{}) {
	SugarLogger.Panic(args...)
}
func Fatal(args ...interface{}) {
	SugarLogger.Fatal(args...)
}

func Infof(template string, args ...interface{}) {
	SugarLogger.Infof(template, args...)
}

func Debugf(template string, args ...interface{}) {
	SugarLogger.Debugf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	SugarLogger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	SugarLogger.Warnf(template, args...)
}
func DPanicf(template string, args ...interface{}) {
	SugarLogger.DPanicf(template, args...)
}
func Panicf(template string, args ...interface{}) {
	SugarLogger.Panicf(template, args...)
}
func Fatalf(template string, args ...interface{}) {
	SugarLogger.Fatalf(template, args...)
}

func getFilePath(dir, prefile, level string) string {
	return filepath.Join(dir, level, prefile) + ".log"
}
