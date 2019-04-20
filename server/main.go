package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prazd/ethAirdrop/server/handlers"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	airdrop := r.Group("/airdrop")

	airdrop.GET("/send/:address", handlers.Airdrop)

	if err := r.Run(":3000"); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
