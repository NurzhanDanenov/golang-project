package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUsers(ctx *gin.Context) {
	users, err := h.services.Authorization.Users(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "database problems")
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *Handler) CreateUser(ctx *gin.Context) {

}

func (h *Handler) GetUserByEmail(ctx *gin.Context) {

}
