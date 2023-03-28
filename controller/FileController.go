package controller

import (
	"io"
	"myblog/common"
	"myblog/model"
	"myblog/response"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Upload 上传图像
func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "格式错误",
		})
		return
	}
	filename := header.Filename
	ext := path.Ext(filename)
	// 用上传时间作为文件名
	name := "image_" + time.Now().Format("20060102150405")
	newFilename := name + ext
	out, err := os.Create("static/images/" + newFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建错误",
		})
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "复制错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"filePath": "/images/" + newFilename},
		"msg":  "上传成功",
	})
}

// 查询收藏
func Collects(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的id
	id := c.Params.ByName("id")
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	// 判断是否已收藏
	for i := 0; i < len(curUser.Collects); i++ {
		if curUser.Collects[i] == id {
			response.Success(c, gin.H{"collected": true, "index": i}, "查询成功")
			return
		}
	}
	response.Success(c, gin.H{"collected": false}, "查询成功")
}

// NewCollect 新增收藏
func NewCollect(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的id
	id := c.Params.ByName("id")
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	var newCollects []string
	newCollects = append(curUser.Collects, id)
	// 更新收藏夹
	if err := db.Model(&curUser).Update("collects", newCollects).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, nil, "更新成功")
}

// UnCollect 取消收藏
func UnCollect(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的index
	index, _ := strconv.Atoi(c.Params.ByName("index"))
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	var newCollects []string
	newCollects = append(curUser.Collects[:index], curUser.Collects[index+1:]...)
	// 更新收藏夹
	if err := db.Model(&curUser).Update("collects", newCollects).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, nil, "更新成功")
}

// Following 查询关注
func Following(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的id
	id := c.Params.ByName("id")
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	// 判断是否已关注
	for i := 0; i < len(curUser.Following); i++ {
		if curUser.Following[i] == id {
			response.Success(c, gin.H{"followed": true, "index": i}, "查询成功")
			return
		}
	}
	response.Success(c, gin.H{"followed": false}, "查询成功")
}

// NewFollow 新增关注
func NewFollow(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的id
	id := c.Params.ByName("id")
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	//var newFollowing []string
	newFollowing := append(curUser.Following, id)
	// 更新关注列表
	if err := db.Model(&curUser).Update("following", newFollowing).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	// 更新粉丝数
	var followUser model.User
	db.Where("id = ?", id).First(&followUser)
	if err := db.Model(&followUser).Update("fans", followUser.Fans+1).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, nil, "更新成功")
}

// UnFollow 取消关注
func UnFollow(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的index
	index, _ := strconv.Atoi(c.Params.ByName("index"))
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	//var newFollowing []string
	newFollowing := append(curUser.Following[:index], curUser.Following[index+1:]...)
	followId := curUser.Following[index]
	// 更新关注列表
	if err := db.Model(&curUser).Update("following", newFollowing).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	// 更新粉丝数
	var followUser model.User
	db.Where("id = ?", followId).First(&followUser)
	if err := db.Model(&followUser).Update("fans", followUser.Fans-1).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, nil, "更新成功")
}
