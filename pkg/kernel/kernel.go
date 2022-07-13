package kernel

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hmdrzaa11/example-go-api/pkg/config"
	"github.com/hmdrzaa11/example-go-api/pkg/database"
	"go.uber.org/zap"
)

// Application is our main app struct that holds our configs.
type Application struct {
	Server *http.Server
	Router *mux.Router
	Logger *zap.Logger
	Config *config.Config
	DB     *sql.DB
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

// Respond is used to send the outgoing response in success cases
func (a *Application) Respond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", a.Config.Http.Content)
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			a.Logger.Fatal(err.Error())
			panic(err)
		}
		return
	}

}

// Boot is going to start our app and set all the configs and return it
func Boot() *Application {
	configs := config.NewConfig()       //get our configs
	router := mux.NewRouter()           //get our router
	logger, err := zap.NewDevelopment() //get our logger
	if err != nil {
		panic(err)
	}
	db, err := database.NewDatabase(configs.App.DatabaseURI) //get a db client

	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	server := &http.Server{
		Addr:         ":" + configs.App.Port,
		Handler:      router,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	//set up CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	router.Use(corsHandler)

	return &Application{
		Router: router,
		Config: configs,
		Logger: logger,
		Server: server,
		DB:     db.Client,
	}
}
