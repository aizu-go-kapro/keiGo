package v1

import (
	"github.com/gin-gonic/gin"
	"keigo/models"
	"net/http"
)

type KeigoController struct{}

func (kc *KeigoController) ConvertKeigo(c *gin.Context) {
	kind := c.Query("kind")
	print(kind)
	var request models.KeigoRequest
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		response := models.KeigoResponse{ConvertedBody: "converted!"}
		c.JSON(http.StatusOK, response)
	}
}
