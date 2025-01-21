package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/stax-refund-service/app"
	"github.com/mrcampbell/stax-refund-service/internal/sqlc"
)

type Server struct {
	routes    *gin.Engine
	resources app.ServerResources
}

func NewServer(queries *sqlc.Queries, resources app.ServerResources) *Server {
	server := &Server{
		resources: resources,
	}

	// todo: maybe find a cleaner way to do this
	// instead of setting resources, passing them to each route?
	server.routes = server.allRoutes()
	return server
}

func (s *Server) Run() {
	s.routes.Run()
}
