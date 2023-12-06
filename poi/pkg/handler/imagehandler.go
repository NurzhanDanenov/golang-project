package handler

import (
	"mime/multipart"
	"net/http"
	"restapi/internal/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createAva(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}
	defer file.Close()

	// pass the file and its name to the controller
	ctx.Set("filePath", header.Filename)
	ctx.Set("file", file)

	// another func
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var image entity.Image

	filename, ok := ctx.Get("filePath")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filename not found"})
	}

	fileGet, ok := ctx.Get("file")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "couldnt find file in request"})
		return
	}

	// upload file
	imageId, err := h.services.UploadImage.Upload(userId, image, fileGet.(multipart.File), filename.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": imageId,
	})
}

type getAllImagesResponse struct {
	Data []entity.Image `json:"data"`
}

func (h *Handler) getAllImages(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	images, err := h.services.UploadImage.GetAll(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, getAllImagesResponse{
		Data: images,
	})
}

func (h *Handler) getImageById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	image, err := h.services.UploadImage.GetById(userId, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, image)
}

func (h *Handler) deleteImage(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.UploadImage.Delete(userId, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
