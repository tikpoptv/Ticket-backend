package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
)

func customEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    customLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func customLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var color string
	switch l {
	case zapcore.DebugLevel:
		color = colorBlue
	case zapcore.InfoLevel:
		color = colorGreen
	case zapcore.WarnLevel:
		color = colorYellow
	case zapcore.ErrorLevel:
		color = colorRed
	case zapcore.FatalLevel:
		color = colorPurple
	default:
		color = colorReset
	}
	enc.AppendString(color + l.CapitalString() + colorReset)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(colorCyan + t.Format("2006-01-02 15:04:05.000") + colorReset)
}

// InitLogger เริ่มต้น logger
func InitLogger() {
	encoderConfig := customEncoderConfig()

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	defer Log.Sync()
}

// Info log ระดับ info
func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

// Error log ระดับ error
func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

// Debug log ระดับ debug
func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

// Warn log ระดับ warn
func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

// Fatal log ระดับ fatal
func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}

// String สร้าง zap.Field สำหรับ string
func String(key, value string) zap.Field {
	return zap.String(key, value)
}

// Int สร้าง zap.Field สำหรับ int
func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

// Error สร้าง zap.Field สำหรับ error
func ErrorField(err error) zap.Field {
	return zap.Error(err)
}
