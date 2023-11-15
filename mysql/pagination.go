package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParsePaginationFromQuery(context *gin.Context) (int, int, string) {
	start := context.DefaultQuery("start", "0")
	limit := context.DefaultQuery("limit", "10")
	words := context.DefaultQuery("words", "")

	startInt, _ := strconv.Atoi(start)
	limitInt, _ := strconv.Atoi(limit)

	return startInt, limitInt, words
}

func GetRecordsTotalCount(db *gorm.DB, table string, owner string) (int64, error) {
	var count int64
	result := db.Table(table).Where("owner = ?", owner).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
