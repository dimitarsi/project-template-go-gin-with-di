package v1

import (
	"net/http"
	"regin/contracts"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"
)

// @Route /api/v1
func GetV1Response(cnt *di.Container) gin.HandlerFunc {
	return func(c *gin.Context) {

		service, ok := (*cnt).Get("dummy-data").(contracts.GetDummyData)

		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"message": "cannot convert dummy-data to contracts.GetDummyData",
			})
		}

		if service != nil {
			c.JSON(http.StatusOK, gin.H{
				"dummyData": service.GetData(),
			})
		}
	}

}
