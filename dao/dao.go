package dao

import (
	"fmt"
	"log"
	"runtime"
	"simpledemo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var Mgr Manager
type manager struct {
	db *gorm.DB
}

type Manager interface {
	Feed() []model.Video
	Publish(video model.Video) 
	PublishList() []model.Video
	GetUser(pid int64) model.UserInfo
	InsertUser(user *model.User)
	GetUserPasswd(username string) string
	GetUserByPassAndUsername(username string,password string) (model.UserInfo,bool)
	GetUserByUserName(username string) model.User
	GetUserByUUID(uuid string) model.User
}
var sysType = runtime.GOOS
func init() {
	
	var dsn string
	if sysType == "linux" {
		// LINUX系统
		dsn = "root:cjghaolihai666__@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	}

	if sysType == "windows" {
		// windows系统
		dsn = "root:cjghaolihai666__@tcp(110.42.180.195:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	}
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Video{})
}
func (mgr *manager) Feed() (videos []model.Video){
	fmt.Println("feed test")
	return
}

func (mgr *manager) Publish(video model.Video) {
	mgr.db.Create(video)
}

func (mgr *manager) PublishList() []model.Video{
	var videos []model.Video
	mgr.db.Find(&videos)
	return videos
}

func (mgr *manager) GetUser(pid int64) model.UserInfo{
	user := model.UserInfo{}
	mgr.db.First(&user,pid)
	return user
}

func (mgr *manager) GetUserByPassAndUsername(username string,password string) (model.UserInfo,bool){
	user := model.User{}
	mgr.db.Where("name = ? AND password = ?",username,password).Find(&user)
	userInfo := model.UserInfo{
		Id: user.Id,
		Name: user.Name,
		FollowCount: user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow: user.IsFollow,
	}
	var ok =false
	if user.Name != "" {
		ok = true
	}
	return userInfo,ok
}
func (mgr *manager) GetUserPasswd(username string) string{
	user := model.User{}
	mgr.db.Where("name = ? ",username).First(&user)

	return user.Password
}

func (mgr *manager) InsertUser(user *model.User){
	mgr.db.Create(user)
}

func (mgr *manager) GetUserByUserName(username string) model.User{
	user := model.User{}
	mgr.db.Where("name = ?",username).First(&user)
	return user
}

func(mgr *manager) GetUserByUUID(uuid string) model.User{
	user := model.User{}
	mgr.db.Where("uuid = ?",uuid).First(&user)
	return user
}