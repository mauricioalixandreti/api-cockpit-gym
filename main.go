package main

import (
	repository "API-COCKPIT-GYM/Repository"
	routes "API-COCKPIT-GYM/Routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.InitFirebase()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.File("./View/EndPoints/index.html")
	})

	router.StaticFS("/static", http.Dir("./View/EndPoints"))

	routes.SetupRoutes(router)

	log.Println("Servidor web rodando em http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
