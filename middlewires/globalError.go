package middlewires

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdminStudy/log"
	"goAdminStudy/utils"
	"net/http"
	"runtime/debug"
)

// 全局异常处理
func GlobalError(c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {

			c.Status(http.StatusInternalServerError)

			errStr := fmt.Sprintf("%v", err)

			c.JSON(http.StatusInternalServerError, utils.ErrorResult(errStr))

			log.Errorln(errStr + ":")
			log.Errorf("%s", fmt.Sprintf("%v", string(debug.Stack())))

		}
	}()

	c.Next()

}
