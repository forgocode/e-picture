package page

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page      int       `json:"page"`
	PageSize  int       `json:"pageSize"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func GetPagination(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
	startTime := ctx.Query("startTime")
	endTime := ctx.Query("endTime")
	fmt.Println(startTime, endTime)
	loc, _ := time.LoadLocation("Local")
	timeLayout := "2006-01-02 15:04:05"
	theStartTime, _ := time.ParseInLocation(timeLayout, startTime, loc)
	theEndTime, _ := time.ParseInLocation(timeLayout, endTime, loc)
	pagination := &Pagination{Page: page, PageSize: pageSize, StartTime: theStartTime, EndTime: theEndTime}
	fmt.Printf("page:%+v\n", pagination)

}
