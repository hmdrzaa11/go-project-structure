package kernel

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hmdrzaa11/example-go-api/pkg/config"
	"go.uber.org/zap"
)

// Application is our main app struct that holds our configs.
type Application struct {
	Server *http.Server
	Router *mux.Router
	Logger *zap.Logger
	Config *config.Config
}

// GracefullShutdown is going to wait for signals to shutdown gracefully
func (a *Application) GracefullShutdown() {
	//create a channel to listen for os signals
	intruptChan := make(chan os.Signal, 1)

	//now we want to pass message to this channel when this things happend
	signal.Notify(intruptChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	//listen to the channel and block the main thread
	<-intruptChan
	//this code will run after we get any msg from channel
	a.Logger.Warn("Received a shutdown signal, gracefully shutting down...")
	//now we create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	//we tell the server to shutdown
	a.Server.Shutdown(ctx)
	os.Exit(0)
}

// Run starts the server
func (app *Application) Run() {
	app.Logger.Info(fmt.Sprintf("listening on port %s", app.Config.App.Port))
	if err := app.Server.ListenAndServe(); err != nil {
		app.Logger.Fatal(err.Error())
	}
}

// Boot is going to start our app and set all the configs and return it
func Boot() *Application {
	configs := config.NewConfig()       //get our configs
	router := mux.NewRouter()           //get our router
	logger, err := zap.NewDevelopment() //get our logger
	server := &http.Server{
		Addr:         ":" + configs.App.Port,
		Handler:      router,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	if err != nil {
		panic(err)
	}
	//set up CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	router.Use(corsHandler)

	return &Application{
		Router: router,
		Config: configs,
		Logger: logger,
		Server: server,
	}
}
