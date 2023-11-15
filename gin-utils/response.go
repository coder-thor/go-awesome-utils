package ginUtils

import "github.com/gin-gonic/gin"

type MessageOptions struct {
	Error  string `json:"error"`
	Data   any    `json:"data"`
	Status int    `json:"code"`
}

func DistributeErrorMessage(err string, status int) gin.H {
	return gin.H{
		"status": status,
		"error":  err,
		"data":   "",
	}
}

func DistributeOKMessage(data any) gin.H {
	return gin.H{
		"status": 200,
		"error":  "",
		"data":   data,
	}
}
