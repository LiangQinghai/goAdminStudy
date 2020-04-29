package middlewires

import (
	"github.com/gin-gonic/gin"
	"goAdminStudy/log"
	"time"
)

func LoggerMid() gin.HandlerFunc {
	return func(context *gin.Context) {

		start := time.Now()

		context.Next()

		// 请求处理时间
		duration := -start.Sub(time.Now())

		//
		request := context.Request
		method := request.Method
		host := request.Host
		uri := request.RequestURI
		addr := request.RemoteAddr
		proto := request.Proto
		ip := context.ClientIP()
		//
		status := context.Writer.Status()

		log.Infoln(ip, addr, host, proto, uri, method, duration, status)

	}
}
