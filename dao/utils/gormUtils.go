package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerFuncWithDB func(c *gin.Context, db *gorm.DB)

func WithDB(db *gorm.DB, handler HandlerFuncWithDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, db)
	}
}