package server

import (
	"kalendar/defined"

	"github.com/gin-gonic/gin"
)

func registerEndpoint[I any, O any](path string, useCase func(*defined.Context[I, O]) *defined.Context[I, O]) {
	Server.POST(path, func(c *gin.Context) {
		ctx := &defined.Context[I, O]{}
		if err := c.ShouldBindJSON(&ctx.Input); err != nil {
			ctx.WithError("ERROR_PARSE_JSON", err.Error())
			return
		}

		ctx = useCase(ctx)
		if ctx.Code == "OK" {
			c.JSON(ctx.StatusCode, gin.H{
				"code": ctx.Code,
				"data": ctx.Output,
			})
		} else {
			c.JSON(ctx.StatusCode, gin.H{
				"code":  ctx.Code,
				"error": ctx.Message,
			})
		}
	})
}
