package handlers

import (
	commonMiddleware "ctint-conv/src/common/middlewares"
	"io"
	"net/http"

	"go.uber.org/zap/zapcore"
)

func SetLoggerLevel(w http.ResponseWriter, r *http.Request) {
	level := r.URL.Query().Get("level")

	switch level {
	case "debug":
		commonMiddleware.GlobalLogger.Info("Setting log level to DEBUG")
		commonMiddleware.UpdateLoggerLevel(zapcore.DebugLevel)
	case "info":
		commonMiddleware.GlobalLogger.Info("Setting log level to INFO")
		commonMiddleware.UpdateLoggerLevel(zapcore.InfoLevel)
	case "warn":
		commonMiddleware.GlobalLogger.Info("Setting log level to WARN")
		commonMiddleware.UpdateLoggerLevel(zapcore.WarnLevel)
	case "error":
		commonMiddleware.GlobalLogger.Info("Setting log level to WARN")
		commonMiddleware.UpdateLoggerLevel(zapcore.ErrorLevel)
	case "panic":
		commonMiddleware.GlobalLogger.Info("Setting log level to WARN")
		commonMiddleware.UpdateLoggerLevel(zapcore.PanicLevel)
	default:
		commonMiddleware.GlobalLogger.Warn("Invalid log level specified")
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Update loggger level to "+level)
}
