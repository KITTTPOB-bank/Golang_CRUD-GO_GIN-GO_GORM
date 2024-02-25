package main

import (
	"backenddemo/pkg/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	routes.SetupRouter(r)
	http.Handle("/", r)
	http.ListenAndServe(":8010",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			// handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Accept", "AllowedOrigins", "AllowedHeaders", "AllowedMethods", "Content-Length", "Accept-Encoding", "X-Requested-With"}),
		)(r))

}


