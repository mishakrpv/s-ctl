package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() http.Handler {
	engine := gin.New()

	var handler http.Handler = engine

	return handler
}
