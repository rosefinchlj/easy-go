package log

import (
	"go.uber.org/zap"
	"os"
)

//
// todo 后面要做详细配置：日志滚动，日志文件名，日志保留时间...
//
var (
	modeEnv       = "GO_ENV"
	modeLevelProd = "release"

	// log meta:
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func init() {
	//
	mode := os.Getenv(modeEnv)

	// 日志级别:
	if mode == modeLevelProd {
		logger, _ = zap.NewProduction()
		defer logger.Sync() // flushes buffer, if any

		sugar = logger.Sugar()
		//Infof("log level: %v mode", mode)
	} else {
		logger, _ = zap.NewDevelopment()
		defer logger.Sync() // flushes buffer, if any

		sugar = logger.Sugar()
		//Info("log level: dev mode")
	}

	return
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	sugar.Warn(args...)
}

func Debug(args ...interface{}) {
	sugar.Debug(args...)
}

func Error(args ...interface{}) {
	sugar.Debug(args...)
}

func Fatal(args ...interface{}) {
	sugar.Fatal(args...)
}

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func FatalF(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}
