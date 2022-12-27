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

	//rotas URI da API USER
	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/id/:id", server.getUserById)

	//rotas URI da API Category
	router.POST("/category", server.createCategory)
	router.GET("/category/id/:id", server.getCategory)
	router.GET("/category", server.getCategories)
	router.PUT("/category/:id", server.updateCategory)
	router.DELETE("/category/:id", server.deleteCategory)

	router.POST("/account", server.createAccount)
	router.GET("/account/id/:id", server.getAccount)
	router.GET("/account/graph/:user_id/:type", server.getAccountGraph)
	router.GET("/account/reports/:user_id/:type", server.getAccountRports)
	router.DELETE("/account/:id", server.deleteAccount)
	router.PUT("/account/:id", server.updateAccount)
	router.GET("/account", server.getAccounts)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Api has error": err.Error()}
}

// teste
