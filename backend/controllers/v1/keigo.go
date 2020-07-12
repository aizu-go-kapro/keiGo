package v1

import (
	"net/http"

	"github.com/aizu-go-kapro/keiGo/backend/models"
	"github.com/gin-gonic/gin"
)

const (
	teineiKind = "teinei"
	sonkeiKind = "sonkei"
	kenjyoKind = "kenjyo"
)
// Keigo is controller of keigo
type Keigo struct {
	teinei models.Teinei
	sonkei models.Sonkei
	kenjyo models.Kenjyo
}

// NewKeigo create keigo controller
func NewKeigo() *Keigo {
	return &Keigo{}
}

// ConvertKeigo is controller
func (kc *Keigo) ConvertKeigo(c *gin.Context) {
	var (
		request  models.KeigoRequest
		response models.KeigoResponse
	)

	kind := c.Query("kind")
	if kind == "" {
		return c.JSON(
			http.StatusBadRequest, 
			map[string]string{"status": "error", "message": "kind is empty"}
			// '{"status": "error", "message": "kind is empty"}'
		)
	}
	
	if err := c.BindJSON(&request); err != nil {
		return c.Status(http.StatusBadRequest)
	}

	switch kind {
	case teineiKind:
		response.ConvertedBody = kc.teinei.Convert(request.Body)
	case sonkeiKind:
		response.ConvertedBody = kc.sonkei.Convert(request.Body)
	case kenjyoKind:
		response.ConvertedBody = kc.kenjyo.Convert(request.Body)
	}

	return c.JSON(http.StatusOK, response)
}
