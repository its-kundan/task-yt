package api

import (
	"net/http"
	"strconv"
	"youtube-fetcher/db"

	"github.com/gin-gonic/gin"
)

func GetVideos(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var videos []db.Video
	db.DB.Order("published_at desc").Limit(limit).Offset(offset).Find(&videos)

	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"videos": videos,
	})
}
