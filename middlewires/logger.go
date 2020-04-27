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
		duration := start.Sub(time.Now())

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

		log.Infof("ClientIp: %14s ### Address: %14s ### Host: %14s ### Proto: %9s ### Uri: %s  ### Method: %6s ### Cost: %d ### Status: %3d ",
			ip, addr, host, proto, uri, method, duration, status)

	}
}
