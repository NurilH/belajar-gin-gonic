package http

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/NurilH/belajar-gin-gonic/pkg/common"
	"github.com/NurilH/belajar-gin-gonic/pkg/common/helpers"
	"github.com/gin-gonic/gin"
)

type DocumentsHTTPDelivery struct {
	common.Controller
	route *gin.RouterGroup
}

func DocumentsNewDelivery(route *gin.RouterGroup) (routeGroup *gin.RouterGroup) {
	documentsHTTPDelivery := DocumentsHTTPDelivery{
		route: route,
	}

	routeGroup = route.Group("/document")
	{
		routeGroup.POST("", documentsHTTPDelivery.UploadDocument)
		routeGroup.DELETE("/:file_name", documentsHTTPDelivery.DeleteDocument)
	}
	route.DELETE("/documents", documentsHTTPDelivery.BulkDeleteDocument)
	route.GET("/documents", documentsHTTPDelivery.ListDocuments)

	return
}

func (h *DocumentsHTTPDelivery) UploadDocument(ctx *gin.Context) {

	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "File Is Required",
			"error":   err.Error(),
		})
		return
	}

	pathUpload := helpers.GetUploadDir()
	unixName := h.UnixFileName("", filepath.Ext(file.Filename))
	filePath := fmt.Sprint(pathUpload + "/" + unixName)

	if err := os.MkdirAll(pathUpload, 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create upload directory",
			"error":   err.Error(),
			"path":    pathUpload,
		})
	}

	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed To Save File",
			"error":   err.Error(),
			"path":    pathUpload,
		})
		return
	}

	url := h.BaseURL(ctx) + filePath
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Success Upload",
		"filename":   unixName,
		"path":       filePath,
		"url":        url,
		"upload_dir": pathUpload,
	})
}

func (h *DocumentsHTTPDelivery) DeleteDocument(ctx *gin.Context) {
	fileName := ctx.Param("file_name")
	if fileName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file_name is required",
		})
		return
	}

	err := os.Remove(helpers.GetUploadDir() + "/" + fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed Remove Document",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Document",
	})
}

func (h *DocumentsHTTPDelivery) BulkDeleteDocument(ctx *gin.Context) {
	fileName := ctx.QueryArray("file_name")
	if len(fileName) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file_name is required",
		})
		return
	}

	for _, name := range fileName {
		err := os.Remove(helpers.GetUploadDir() + "/" + name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed Remove Document",
				"erorr":   err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Document",
	})
}

func (h *DocumentsHTTPDelivery) ListDocuments(ctx *gin.Context) {

	files, err := os.ReadDir(helpers.GetUploadDir())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to read upload file",
			"error":   err.Error(),
		})
	}

	var documents []model.DocumentInfo
	baseUrl := h.BaseURL(ctx)

	for _, file := range files {
		fileName := file.Name()

		doc := model.DocumentInfo{
			Filename: fileName,
			URL:      baseUrl + helpers.GetUploadDir() + "/" + fileName,
		}
		documents = append(documents, doc)
	}

	ctx.JSON(http.StatusOK, documents)
}
