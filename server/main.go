package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prazd/ethAirdrop/server/handlers"
	"os"
	"log"
)

func main(){
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())


	r.GET("/send/:address", handlers.Airdrop)

	if err := r.Run(":3000"); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}