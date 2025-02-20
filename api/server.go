package api

import (
	db "github.com/RG-7/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccount)

	server.router = router
	return server
}

// error response
func errorResonsponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// start rns the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
