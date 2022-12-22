package api

import (
	db "github.com/francivaldo-alves/gofinance-bankend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//rotas URI da API
	router.POST("/user", server.createUser)
	router.GET("/user:username", server.getUser)
	router.GET("/user:id", server.getUserById)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Api has error": err.Error()}
}
