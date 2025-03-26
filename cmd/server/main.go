package main

import (
	"API_Books/internal/database"
	"API_Books/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao banco de dados
	err := database.Connect()
	if err != nil {
		log.Fatalln("Falha ao conectar com banco de dados:", err)
	}
	
	// Criar o roteador
	router := gin.Default()

	// Configurar IPs confiáveis
	err = router.SetTrustedProxies([]string{"192.168.1.1", "192.168.0.46"})
	if err != nil {
		log.Fatalln("Erro ao definir proxies confiáveis:", err)
	}

	// Registrar rotas
	routes.RegisterRoutes(router)

	// Iniciar o servidor
	log.Println("http://localhost:8080/books")
	router.Run(":8080")
}
