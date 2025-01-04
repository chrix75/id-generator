package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"id-generator/generator"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	machineID, err := getMachineID()
	if err != nil {
		log.Fatal(err)
	}

	port, err := getServerPort()
	if err != nil {
		log.Fatal(err)
	}

	router := setupRouter(machineID)

	log.Fatal(router.Run(fmt.Sprintf(":%d", port)))
}

func getServerPort() (int, error) {
	v := os.Getenv("ID_GENERATOR_SERVER_PORT")
	if v == "" {
		v = "8080"
	}

	port, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("invalid server port %s", v)
	}

	return port, nil
}

func getMachineID() (uint8, error) {
	v := os.Getenv("ID_GENERATOR_MACHINE_ID")
	if v == "" {
		return 0, fmt.Errorf("the environment variable ID_GENERATOR_MACHINE_ID is not set")
	}

	machineID, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("invalid ID_GENERATOR_MACHINE_ID, %w", err)
	}

	if machineID < 1 || machineID > 255 {
		return 0, fmt.Errorf("ID_GENERATOR_MACHINE_ID out of range, must be between 1 and 255")
	}

	return uint8(machineID), nil
}

type IDValue struct {
	ID uint64 `json:"id"`
}

func setupRouter(machineID uint8) *gin.Engine {
	gen := generator.NewGenerator(machineID)
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
