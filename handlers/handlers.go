package handlers

import (
	"net/http"

	"github.com/Heads-and-Hands/handh-school-back/database"
	"github.com/Heads-and-Hands/handh-school-back/models"
	"github.com/gin-gonic/gin"
)

func GetHandler(ctx *gin.Context) {
	g := database.OrmProvider.GetRequests()
	ctx.JSON(http.StatusOK, g)
}

func PostHandler(ctx *gin.Context) {
	var body models.CreateUserBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.NewErrorByWrapping(err))
		return
	}

	database.OrmProvider.CreateRequest(body)
	ctx.JSON(http.StatusOK, body)
}
