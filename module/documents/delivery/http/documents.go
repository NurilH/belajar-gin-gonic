package http

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/NurilH/belajar-gin-gonic/pkg/common"
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

	// fileType := file.Header.Get("content-type")
	unixName := h.UnixFileName("", filepath.Ext(file.Filename))
	fileDir := fmt.Sprintf("/static/file/%s", unixName)
	err = ctx.SaveUploadedFile(file, "."+fileDir)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed To Save File",
			"error":   err.Error(),
		})
		return
	}

	url := h.BaseURL(ctx) + fileDir
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Upload",
		"path":    url,
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

	err := os.Remove("./static/file/" + fileName)
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
		err := os.Remove("./static/file/" + name)
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
