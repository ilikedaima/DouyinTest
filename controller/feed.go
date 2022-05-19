package controller

import (
	"net/http"
	"simpledemo/model"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	model.Response
	VideoList []model.VideoInfo `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// var videos []Video 
	
	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
