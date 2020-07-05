package v1

import (
	"github.com/gin-gonic/gin"
	"keigo/models"
	"net/http"
)

type KeigoController struct{}

const (
	Teinei string = "teinei"
	Sonkei string = "sonkei"
	Kenjyo string = "kenjyo"
)

func (kc *KeigoController) ConvertKeigo(c *gin.Context) {
	kind := c.Query("kind")
	print(kind)
	var request models.KeigoRequest
	var response models.KeigoResponse
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		var kagome models.Kagome
		tokens := kagome.MorphologicalAnalysis(request.Body)
		switch kind {
		case Teinei:
			var teinei models.Teinei
			response.ConvertedBody = teinei.Convert(tokens)
		case Sonkei:
			var sonkei models.Sonkei
			response.ConvertedBody = sonkei.Convert(tokens)
		case Kenjyo:
			var kenjyo models.Kenjyo
			response.ConvertedBody = kenjyo.Convert(tokens)
		}
		c.JSON(http.StatusOK, response)
	}
}
