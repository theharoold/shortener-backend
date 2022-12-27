package main

import (
	"github.com/gin-gonic/gin"
	"github.com/theharoold/shortener-backend/api"
)

func main() {
	r := gin.Default()
	shortener := r.Group("/shortener-api")
	v1 := shortener.Group("/v1")

	var a api.API
	a.SetupRoutes(v1)
}
