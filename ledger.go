package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRepositoryIndex(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "POST request handled"})
}

