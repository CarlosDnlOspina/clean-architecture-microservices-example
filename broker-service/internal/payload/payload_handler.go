package payload

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) Broker(c *gin.Context) {
	res, err := h.Service.Broker()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Authenticate(c *gin.Context) {
	var requestPayload RequestPayload
	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch requestPayload.Action {
	case "auth":
		res, err := h.Service.Authenticate(c.Request.Context(), &requestPayload.Auth)

		if err != nil {
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
