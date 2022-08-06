package main

import (
	"context"
	"os"
	"os/signal"
	"simplems/app"

	"syscall"
	"time"

	"go.uber.org/zap"
)

// @title          Swagger Products API
// @version        1.0
// @description    This is a sample Products server.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:3000
// @BasePath /
func main() {
	logger, _ := zap.NewDevelopment()
	app := app.NewApplication(logger)
	defer logger.Sync()

	// anonymouse go routine which runs concurrently in the backgroun!
	go func() {
		app.Start()
	}()

	notifyChannel := make(chan os.Signal, 1)
	signal.Notify(notifyChannel, os.Interrupt, syscall.SIGTERM)

	// block untill signal is recieved
	signal := <-notifyChannel
	logger.Info("Gracefully Shutting down...", zap.String("reason", signal.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	app.Shutdown(ctx)
}
