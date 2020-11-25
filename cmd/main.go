package main

import (
	"SeminarioGo/internal/config"
	"SeminarioGo/internal/database"
	"SeminarioGo/internal/services"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	configFile := flag.String("config", "../config.yaml", "this is the service config")
	flag.Parse()
	fmt.Println("Primera parte")
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Segunda parte")

	db, err := database.NewDataBase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Tercera parte")

	err = database.CreateSchema(db)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Cuarta parte")
	service, err := services.NewService(db, cfg)
	htppService := services.NewHTTPTransport(service)
	router := gin.Default()
	htppService.Register(router)
	router.Run()
}
