package v1

import (
	"log"
	"net/http"

	"github.com/aizu-go-kapro/keiGo/backend/models"
	"github.com/gin-gonic/gin"
)

type KeigoController struct{}

const (
	Teinei = "teinei"
	Sonkei = "sonkei"
	Kenjyo = "kenjyo"
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
			cb, err := teinei.Convert(request.Body)
			if err != nil {
				log.Fatal("Convert Error!", err)
			}
			response.ConvertedBody = cb
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
