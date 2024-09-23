package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yjwj-3/content-release/cmd/api/app/config"
	"github.com/yjwj-3/content-release/internal/api/types"
	"net/http"
)

type Service struct {
	Gin      *gin.Engine
	Handlers map[string]types.Handler
}

func NewApiServer(c *config.ApiConfig) *Service {
	s := new(Service)
	// 注册Handler
	s.Gin.POST(fmt.Sprintf("/%s", c.App), func(c *gin.Context) {
		action := c.Param("action")
		if _, ok := s.Handlers[action]; !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "action not found"})
			return
		}
		s.Handlers[action](c)
	})
	return s
}

func (s *Service) Serve() error {
	err := s.Gin.Run(":8080")
	return err
}
