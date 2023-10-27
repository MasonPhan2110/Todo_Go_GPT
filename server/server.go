package server

import (
	db "MasonPhan2110/Todo_Go_GPT/db/sqlc"
	"MasonPhan2110/Todo_Go_GPT/pkg/token"
	"MasonPhan2110/Todo_Go_GPT/routes"
	"MasonPhan2110/Todo_Go_GPT/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	r := gin.Default()

	routes.AddRoutes(r)

	server.router = r
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
