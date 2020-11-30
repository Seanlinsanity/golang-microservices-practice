package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "status OK here"
)

func Marco(c *gin.Context) {
	c.String(http.StatusOK, polo)
}
