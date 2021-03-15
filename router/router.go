package router

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/FideTech/yaus/config"
	"github.com/FideTech/yaus/controllers"
	"github.com/FideTech/yaus/router/middleware"
)

var (
	server *http.Server
)

//Start initates the router and starts listening on the configured port
func Start() error {
	mux := http.NewServeMux()

	mux.Handle("/i/", middleware.NewParsedLink(controllers.InfoHandler, "i"))
	mux.Handle("/e/", middleware.NewParsedLink(controllers.ErrorLinkHandler, "e"))
	mux.Handle("/d/", middleware.NewParsedLink(controllers.DynamicLinkHandler, "d"))
	mux.Handle("/api/", middleware.NewAPIAuth(controllers.ApiHandler))

	server = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Config.System.Router.Port),
		Handler: middleware.NewLogger(middleware.NewCors(mux)),
	}

	log.Printf("router started at %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("failed to listen and serve the router: %s\n", err)
		return err
	}
	return nil
}

//Shutdown tells the http server to gracefully shutdown and close existing connections
func Shutdown(ctx context.Context) error {
	log.Println("shutting down the router 🛑")

	return server.Shutdown(ctx)
}
