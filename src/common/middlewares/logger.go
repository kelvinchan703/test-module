package middleware

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var GlobalLogger *zap.SugaredLogger
var atomicLevel zap.AtomicLevel

func InitGlobalLogger(logLevel zapcore.Level, filename string, maxSize int, maxBackups int, maxAge int) {
	// Configure encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Create console encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Create file encoder
	rollingFileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// Create lumberjack logger for rolling file
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename + ".log",
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   true,
	}

	// Create rolling file sink
	rollingFileSink := zapcore.AddSync(lumberjackLogger)

	// Define log level
	atomicLevel = zap.NewAtomicLevel()
	atomicLevel.SetLevel(logLevel)

	// Create file core
	fileCore := zapcore.NewCore(rollingFileEncoder, rollingFileSink, atomicLevel)

	// Create console core
	consoleCore := zapcore.NewCore(consoleEncoder, os.Stdout, atomicLevel)

	// Use multi-core to write to both console and file
	core := zapcore.NewTee(consoleCore, fileCore)

	// Create logger
	_logger := zap.New(core)

	defer _logger.Sync()
	GlobalLogger = _logger.Sugar()
}

// New returns a logger middleware for chi, that implements the http.Handler interface.
func InitLogger(sugarLogger *zap.SugaredLogger) func(next http.Handler) http.Handler {
	if sugarLogger == nil {
		return func(next http.Handler) http.Handler { return next }
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			sugarLogger.Infow("Start Request", zap.String("path", r.URL.Path),
				zap.String("traceId", r.Header.Get("traceId")),
				zap.String("tenant", r.Header.Get("tenant")),
				zap.String("sourceId", r.Header.Get("sourceId")),
				zap.String("previousId", r.Header.Get("previousId")),
				zap.String("proto", r.Proto))
			// API 响应后的 response
			defer func() {

				if ww.Status() == 200 {
					sugarLogger.Infow("End Request", zap.String("path", r.URL.Path),
						zap.String("traceId", r.Header.Get("traceId")),
						zap.String("tenant", r.Header.Get("tenant")),
						zap.String("sourceId", r.Header.Get("sourceId")),
						zap.String("previousId", r.Header.Get("previousId")),
						zap.String("proto", r.Proto))
				} else {
					sugarLogger.Errorw("End Bad Request", zap.String("path", r.URL.Path),
						zap.String("traceId", r.Header.Get("traceId")),
						zap.String("tenant", r.Header.Get("tenant")),
						zap.String("sourceId", r.Header.Get("sourceId")),
						zap.String("previousId", r.Header.Get("previousId")),
						zap.String("proto", r.Proto))
				}
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}

// Function to update logger level
func UpdateLoggerLevel(logLevel zapcore.Level) {
	atomicLevel.SetLevel(logLevel)
}

// registrateLogFormat
func LogFormat(r *http.Request) []interface{} {

	return []interface{}{}
}
