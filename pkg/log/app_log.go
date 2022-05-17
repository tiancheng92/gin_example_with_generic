package log

import (
	"github.com/tiancheng92/gf"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func Init(logLevel string) {
	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	logConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(level), // 日志级别
		Development:       false,                       // 开发模式，堆栈跟踪
		DisableStacktrace: true,                        // 关闭自动堆栈捕获
		Encoding:          "console",                   // 输出格式 console 或 json
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			CallerKey:     "caller",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(gf.StringJoin("[", level.CapitalString(), "]"))
			},
			EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(t.Format("[2006-01-02 15:04:05.000]"))
			},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller: func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(gf.StringJoin("[", caller.TrimmedPath(), "]:"))

			},
			ConsoleSeparator: " ",
		}, // 编码器配置
		InitialFields:    nil,                // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout"}, // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"}, // 错误输出到指定文件
	}

	// 构建日志
	l, err := logConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger = l
}

func DebugWithArg(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func InfoWithArg(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func WarnWithArg(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func ErrorWithArg(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func FatalWithArg(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Debug(args ...any) {
	logger.Sugar().Debug(args...)
}

func Info(args ...any) {
	logger.Sugar().Info(args...)
}

func Warn(args ...any) {
	logger.Sugar().Warn(args...)
}

func Error(args ...any) {
	logger.Sugar().Error(args...)
}

func Fatal(args ...any) {
	logger.Sugar().Fatal(args...)
}

func Debugf(template string, args ...any) {
	logger.Sugar().Debugf(template, args...)
}

func Infof(template string, args ...any) {
	logger.Sugar().Infof(template, args...)
}

func Warnf(template string, args ...any) {
	logger.Sugar().Warnf(template, args...)
}

func Errorf(template string, args ...any) {
	logger.Sugar().Errorf(template, args...)
}

func Fatalf(template string, args ...any) {
	logger.Sugar().Fatalf(template, args...)
}
