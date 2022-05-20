package controller

import (
	"fmt"
	"net/http"
	"simpledemo/dao"
	"simpledemo/model"
	"simpledemo/utils"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	model.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	model.Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println("username:",username,"password:",password)
	if username=="" || password == "" {
		fmt.Println("用户或者密码不能为空")
		return
	}
	// token := username + password
	//password 用"golang.org/x/crypto/bcrypt"加密
	encodePWD,err := utils.PasswordHash(password)
	if err!=nil {
		fmt.Println(err)
	}
	if _,ok := dao.Mgr.GetUserByPassAndUsername(username,encodePWD); ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		// atomic.AddInt64(&userIdSequence, 1)
		newUser := model.User{
			Name: username,
			Password: encodePWD,
		}
		dao.Mgr.InsertUser(&newUser)

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + encodePWD,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var token string
	
	realPasswd := dao.Mgr.GetUserPasswd(username)
	b := utils.CheckPasswd(password,realPasswd)
	if b {
		user := dao.Mgr.GetUserByUserName(username)
		token = user.Name + realPasswd
		fmt.Println(token)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response:model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
