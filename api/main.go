package main

import (
	"github.com/gin-gonic/gin"
	"id-generator/generator"
	"log"
	"net/http"
	"time"
)

func main() {
	router := setupRouter()

	log.Fatal(router.Run(":8080"))
}

type IDValue struct {
	ID uint64 `json:"id"`
}

func setupRouter() *gin.Engine {
	gen := generator.NewGenerator(1)
	engine := generator.NewEngine(gen)

	r := gin.Default()
	r.GET("/api/id", func(c *gin.Context) {
		id := IDValue{
			ID: engine.MustGetID(time.Now()),
		}

		c.JSON(http.StatusOK, id)
	})
	return r
}
