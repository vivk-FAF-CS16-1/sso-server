package application

import (
	"context"
	"cs-lab-6/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type IApp interface {
	Start()
	Shutdown()
}

type clientApp struct {
	IApp
	server *http.Server
}

func NewApp(ctx context.Context) IApp {
	router := gin.New()

	control := controller.New()
	control.RegisterRoutes(router)

	return &clientApp{
		server: &http.Server{
			Addr:    ":" + viper.GetString("port"),
			Handler: router,
		},
	}
}

func (app *clientApp) Start() {
	crt := "./configs/localhost.crt"
	key := "./configs/localhost.key"
	if err := app.server.ListenAndServeTLS(crt, key); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error while running server: %v\n", err)
	}
}

func (app *clientApp) Shutdown() {
	if err := app.server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Unable to shutdown server: %v\n", err)
	}
}
