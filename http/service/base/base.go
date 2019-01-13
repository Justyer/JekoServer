package base

import (
	"github.com/gin-gonic/gin"
)

type BaseService struct {
	Ctx *gin.Context
}
