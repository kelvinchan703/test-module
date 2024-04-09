package logger

import (
	commonConfig "ctint-conv/src/common/config"
	commonMiddlewares "ctint-conv/src/common/middlewares"
	"net/http"
)

func Infow(r *http.Request, msg string, keysAndValues ...interface{}) {
	// Append request header info to the log message
	for _, header := range commonConfig.GlobalConfig.Logger.Headers {
		value := r.Header.Get(header.Name)
		keysAndValues = append(keysAndValues, header.Name, value)
	}

	// Log the message with request header info
	commonMiddlewares.GlobalLogger.Infow(msg, keysAndValues...)
}

func Debugw(r *http.Request, msg string, keysAndValues ...interface{}) {
	// Append request header info to the log message
	for _, header := range commonConfig.GlobalConfig.Logger.Headers {
		value := r.Header.Get(header.Name)
		keysAndValues = append(keysAndValues, header.Name, value)
	}

	// Log the message with request header info
	commonMiddlewares.GlobalLogger.Debugw(msg, keysAndValues...)
}

func Errorw(r *http.Request, msg string, keysAndValues ...interface{}) {
	// Append request header info to the log message
	for _, header := range commonConfig.GlobalConfig.Logger.Headers {
		value := r.Header.Get(header.Name)
		keysAndValues = append(keysAndValues, header.Name, value)
	}

	// Log the message with request header info
	commonMiddlewares.GlobalLogger.Errorw(msg, keysAndValues...)
}

func Fatalw(r *http.Request, msg string, keysAndValues ...interface{}) {
	// Append request header info to the log message
	for _, header := range commonConfig.GlobalConfig.Logger.Headers {
		value := r.Header.Get(header.Name)
		keysAndValues = append(keysAndValues, header.Name, value)
	}

	// Log the message with request header info
	commonMiddlewares.GlobalLogger.Fatalw(msg, keysAndValues...)
}

func Warnw(r *http.Request, msg string, keysAndValues ...interface{}) {
	// Append request header info to the log message
	for _, header := range commonConfig.GlobalConfig.Logger.Headers {
		value := r.Header.Get(header.Name)
		keysAndValues = append(keysAndValues, header.Name, value)
	}

	// Log the message with request header info
	commonMiddlewares.GlobalLogger.Warnw(msg, keysAndValues...)
}

func Panicw(r *http.Request, msg string, keysAndValues ...interface{}) {
	// Append request header info to the log message
	for _, header := range commonConfig.GlobalConfig.Logger.Headers {
		value := r.Header.Get(header.Name)
		keysAndValues = append(keysAndValues, header.Name, value)
	}

	// Log the message with request header info
	commonMiddlewares.GlobalLogger.Panicw(msg, keysAndValues...)
}
