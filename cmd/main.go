package main

import (
	"context"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/yin-zt/go-tamapl-server/pkg/config"
	"github.com/yin-zt/go-tamapl-server/pkg/routes"
	"github.com/yin-zt/go-tamapl-server/pkg/server"
	"github.com/yin-zt/go-tamapl-server/pkg/utils/logger"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var banner = `
                           __                               .__                                                    
   ____   ____           _/  |______    _____ _____  ______ |  |             ______ ______________  __ ___________ 
  / ___\ /  _ \   ______ \   __\__  \  /     \\__  \ \____ \|  |    ______  /  ___// __ \_  __ \  \/ // __ \_  __ \
 / /_/  >  <_> ) /_____/  |  |  / __ \|  Y Y  \/ __ \|  |_> >  |__ /_____/  \___ \\  ___/|  | \/\   /\  ___/|  | \/
 \___  / \____/           |__| (____  /__|_|  (____  /   __/|____/         /____  >\___  >__|    \_/  \___  >__|   
/_____/                             \/      \/     \/|__|                       \/     \/                 \/       

`

func main() {
	defer log.Flush()

	fmt.Println(banner)
	logger.Setup()
	log.ReplaceLogger(logger.ServerLogger)
	server.Cli.Init("do it")
	fmt.Println("see it?")

	r := routes.InitRoutes()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.GinHost, config.GinPort),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s\n", err)
		}
	}()

	log.Info(fmt.Sprintf("Server is running at  %s:%d/%s", config.GinHost, config.GinPort, config.GinUrlPrefix))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown: ", err)
	}
	log.Info("Server exiting!")

}
