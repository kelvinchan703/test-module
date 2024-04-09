package middleware

import (
	commonHandler "ctint-conv/src/common/handlers"
	commonMiddleware "ctint-conv/src/common/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitGlobalSetting(router *chi.Mux, basePath string, version string) {

	// Global Setting
	router.Use(commonMiddleware.ValidateHeaders)
	router.Use(commonMiddleware.ValidateCdssToken)
	router.Use(commonMiddleware.InitLogger(commonMiddleware.GlobalLogger))
	router.Use(commonMiddleware.CorsMiddleware)
	router.Use(commonMiddleware.SetGlobalHeader)
	router.Use(commonMiddleware.GlobalErrorHandler)

	router.Get(basePath+"/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		commonHandler.Healthcheck(w, r, version)
	})

	router.Get(basePath+"/loggerlevel", commonHandler.SetLoggerLevel)
}
