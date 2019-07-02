package utils

import (
	// "encoding/json"
	// "net/http"

	// "math"
	"meg-bot-backend/app/models"

	"github.com/gin-gonic/gin"
)

// func StructToMap(data)  map[string]interface{}{
// 	var inInterface map[string]interface{}
// 	inrec, _ := json.Marshal(data)
// 	json.Unmarshal(inrec, &inInterface)
// 	return inInterface
// }

// MessageOne Set Message For Respond
func MessageOne(status bool, message string, data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

// MessageMany Set Message For Respond
func MessageMany(status bool, message string, pageans models.PageNoAndsize, data []interface{}) map[string]interface{} {
	// var c int
	// if c = b; a > b {
	// 	c = a
	// }
	totalPages := pageans.CountList / pageans.Size
	var previousPage int
	if previousPage = 1; pageans.PageNo-1 != 0 {
		previousPage = pageans.PageNo - 1
	}
	var nextPage int
	if nextPage = totalPages; pageans.PageNo+1 <= totalPages {
		nextPage = pageans.PageNo + 1
	}
	return map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
		// ปัดเลขขึ้น math.Ceil(page),
		"page": map[string]interface{}{
			"previous_page": previousPage,
			"current_page":  pageans.PageNo,
			"next_page":     nextPage,
			"total_pages":   totalPages, // หารแล้วปัดเศษขึ้น
			"total_item":    pageans.CountList,
		},
	}
}

// Respond Set Respond
func Respond(c *gin.Context, number int, data map[string]interface{}) {
	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(data)
	c.JSON(number, data)
}
