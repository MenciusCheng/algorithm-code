package log

import (
	"go.uber.org/zap"
)

var DefaultLogger = InitDefaultLogger()

func InitDefaultLogger() *zap.Logger {
	logger := zap.NewExample()
	return logger
}

// InitAppLog根据options的设置,初始化日志系统。
// 注意默认是测试环境模式,需要设置线上模式的需要设置TestEnv(false)
//func InitAppLog(options ...AppZapOption) error {
//	var (
//		err   error
//		Level zap.AtomicLevel
//	)
//	config := defaultLogOptions
//	for _, option := range options {
//		option.Apply(&config)
//	}
//
//	log.Println("InitAppLog - show config ", config)
//
//	if Level, appInnerLog, err = zapLogInit(&config); err != nil {
//		fmt.Printf("ZapLogInit err:%v", err)
//		return err
//	}
//
//	appInnerLog = appInnerLog.WithOptions(zap.AddCallerSkip(1))
//	logLevelHttpServer(&config, Level)
//	return nil
//}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	DefaultLogger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	DefaultLogger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	DefaultLogger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	DefaultLogger.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(msg string, fields ...zap.Field) {
	DefaultLogger.DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...zap.Field) {
	DefaultLogger.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is disabled.
func Fatal(msg string, fields ...zap.Field) {
	DefaultLogger.Fatal(msg, fields...)
}

func Sync() error {
	return DefaultLogger.Sync()
}

func Logger() *zap.Logger {
	return DefaultLogger
}

//func SetLogLevel(level string) error {
//	switch strings.ToLower(level) {
//	case "debug", "info", "warn", "error", "fatal":
//		level = strings.ToLower(level)
//	case "all":
//		level = "debug"
//	case "off", "none":
//		level = "fatal"
//	default:
//		return errors.New("not support level")
//	}
//	client := http.Client{}
//
//	type payload struct {
//		Level string `json:"level"`
//	}
//	mypayload := payload{
//		Level: level,
//	}
//	bin, err := json.Marshal(mypayload)
//	if err != nil {
//		return err
//	}
//	req, err := http.NewRequest("PUT", setlevelpath, bytes.NewReader(bin))
//	if err != nil {
//		return err
//	}
//	resp, err := client.Do(req)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//	return nil
//}
