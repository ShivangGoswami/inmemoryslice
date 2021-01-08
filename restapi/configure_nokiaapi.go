// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"nokia/models"
	"strings"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"nokia/restapi/ops"
	"nokia/restapi/ops/nokiaapi"
)

//go:generate swagger generate server --target ..\..\nokia --name Nokiaapi --spec ..\api.yml --api-package ops

func configureFlags(api *ops.NokiaapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *ops.NokiaapiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	var result []string
	// if api.NokiaapiPostAddHandler == nil {
	api.NokiaapiPostAddHandler = nokiaapi.PostAddHandlerFunc(func(params nokiaapi.PostAddParams) middleware.Responder {
		result = append(result, *params.Input.Input)
		fmt.Println("result", result)
		return nokiaapi.NewPostAddOK().WithPayload(&models.Status{"ok"})
		//return middleware.NotImplemented("operation nokiaapi.PostAdd has not yet been implemented")
	})
	//}
	//	if api.NokiaapiPostAddHandler == nil {
	api.NokiaapiDeleteDeleteHandler = nokiaapi.DeleteDeleteHandlerFunc(func(params nokiaapi.DeleteDeleteParams) middleware.Responder {
		if temp := *params.Input.Input; true {
			index := -1
			for i, val := range result {
				if strings.EqualFold(val, temp) {
					index = i
					break
				}
			}
			if index == -1 {
				return nokiaapi.NewDeleteDeleteOK().WithPayload(&models.Status{"Not found!"})
			} else {
				result = append(result[0:index], result[index+1:len(result)]...)
			}
		}
		fmt.Println("result", result)
		return nokiaapi.NewDeleteDeleteOK().WithPayload(&models.Status{"ok"})
		//return middleware.NotImplemented("operation nokiaapi.PostAdd has not yet been implemented")
	})
	//	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
