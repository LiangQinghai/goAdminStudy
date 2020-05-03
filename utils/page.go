package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Page struct {
	PageNo     int    `json:"pageNo"`     // 页码
	PageCount  int    `json:"pageCount"`  // 每页数
	TotalCount int    `json:"totalCount"` // 总数
	KeyWorld   string `json:"keyWorld"`   // 关键字
}

func (p *Page) GetPageNo() int {
	if p.PageNo < 1 {
		p.PageNo = 1
	}
	return p.PageNo
}

func (p *Page) GetPageCount() int {
	if p.PageCount < 1 {
		p.PageCount = 20
	}
	return p.PageCount
}

func NewDefaultPage() Page {
	return Page{
		PageCount: 20,
		PageNo:    1,
	}
}

// 分页结果
func SuccessPageResult(page *Page, data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": nil,
		"data":    data,
		"page":    page,
		"time":    time.Now().Format("2006-01-02 15:04:05:1234"),
	})
}
