package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prazd/ethAirdrop/eth"
	"log"
	"net/http"
)

func Airdrop(c *gin.Context) {
	address := c.Param("address")

	txHash, err := eth.SendEth(address)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": 500})
		return
	}

	c.JSON(http.StatusOK, gin.H{"txHash": txHash})
}
