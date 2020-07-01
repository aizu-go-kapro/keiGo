package v1

import (
	"github.com/gin-gonic/gin"
	"keigo/models"
	"net/http"
)

type KeigoController struct{}

func (kc *KeigoController) ConvertKeigo(c *gin.Context) {
	//kind := c.Query("kind")
	//print(kind)
	var request models.KeigoRequest
	var response models.KeigoResponse
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		var kagome models.Kagome
		tokens := kagome.MorphologicalAnalysis(request.Body)
		var teinei models.Teinei
		response.ConvertedBody = teinei.Convert(tokens)
		c.JSON(http.StatusOK, response)
	}
}
