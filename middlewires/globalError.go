package middlewires

import (
	"github.com/gin-gonic/gin"
	"goAdminStudy/utils"
	"net/http"
)

// 全局异常处理
func GlobalError(c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {

			switch err.(type) {
			case string:

				c.Status(http.StatusInternalServerError)

				c.JSON(http.StatusInternalServerError, utils.ErrorResult(err.(string)))

			default:
				panic(err)
			}

		}
	}()

	c.Next()

}
