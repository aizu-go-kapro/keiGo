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
	request := models.KeigoRequest{}
	response := models.KeigoResponse{}
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		switch kind {
		case Teinei:
			teinei := models.Teinei{}
			response.ConvertedBody = teinei.Convert(request.Body)
		case Sonkei:
			sonkei := models.Sonkei{}
			response.ConvertedBody = sonkei.Convert(request.Body)
		case Kenjyo:
			kenjyo := models.Kenjyo{}
			response.ConvertedBody = kenjyo.Convert(request.Body)
		}
		c.JSON(http.StatusOK, response)
	}
}
