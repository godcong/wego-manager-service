//go:generate statik -f -src=./dist
package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/godcong/wego-spread-service/service"

	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/signal"
	"syscall"
)

var configPath = flag.String("config", "config.toml", "load config from path")
var logPath = flag.String("log", "manager.log", "set log name")
var sync = flag.Bool("sync", false, "open to sync the model")

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @schemes http https

// @license.name MIT
// @license.url https://github.com/godcong/wego-auth-manager/blob/master/LICENSE

// @host localhost:8080
// @BasePath /v0
func main() {
	flag.Parse()
	file, err := os.OpenFile(*logPath, os.O_SYNC|os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFormatter(&log.JSONFormatter{})

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//start
	service.Start(cfg)

	go func() {
		sig := <-sigs
		fmt.Println(sig, "exiting")
		service.Stop()
		done <- true
	}()
	<-done

}
