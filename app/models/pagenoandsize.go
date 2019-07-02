package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// PageNoAndsize ข้อความ
type PageNoAndsize struct {
	PageNo    int `json:"pageNo"`
	Size      int `json:"size"`
	CountList int `json:"countList"`
}

// GetDefaultPageNoAndsize get Default
func GetDefaultPageNoAndsize(c *gin.Context) PageNoAndsize {
	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "1"))
	countList, _ := strconv.Atoi("0")
	return PageNoAndsize{
		pageNo,
		size,
		countList,
	}
}
