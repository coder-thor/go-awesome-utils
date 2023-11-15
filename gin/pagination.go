package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParsePaginationFromQuery(context *gin.Context) (int, int, string) {
	start := context.DefaultQuery("start", "0")
	limit := context.DefaultQuery("limit", "10")
	words := context.DefaultQuery("words", "")

	startInt, _ := strconv.Atoi(start)
	limitInt, _ := strconv.Atoi(limit)

	return startInt, limitInt, words
}
