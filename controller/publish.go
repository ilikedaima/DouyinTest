package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"simpledemo/dao"
	"simpledemo/model"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.VideoInfo `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	user := dao.Mgr.GetUserByUUID(token)
	if user.Name=="" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	// user := usersLoginInfo[token]
	// user := dao.Mgr.GetUserByUUID(token)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)

	
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//把信息存到数据库
	video := model.Video{
		Author: user.Id,
		PlayUrl: model.UrlBase+finalName,
		
	}
	dao.Mgr.Publish(&video)
	//-------------------


	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
 
	vidvideoInfos := dao.Mgr.PublishList()
	
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: vidvideoInfos,
	})
}
