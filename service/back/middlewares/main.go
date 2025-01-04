package middlewares

import (
	"github.com/eavlongs/go_backend_template/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MainMiddleware struct {
	db *gorm.DB
}

func NewMainMiddleware(db *gorm.DB) *MainMiddleware {
	return &MainMiddleware{db: db}
}

func (m *MainMiddleware) TestFail(ctx *gin.Context) {
	utils.RespondWithUnauthorizedError(ctx)
	ctx.Abort()
}

func (m *MainMiddleware) TestSuccess(ctx *gin.Context) {
	ctx.Next()
}
