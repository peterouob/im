package service

import (
	"fmt"
	"gin-chat/models"
	"gin-chat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Summary 所有用戶
// @Tags 使用者
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"status": data,
	})
}

// CreateUser
// @Summary 新增用戶
// @Tags 使用者
// @Param name query string false "用戶名"
// @Param password query string false "密碼"
// @Param repassword query string false "確認密碼"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	salt := fmt.Sprintf("%06d", rand.Int31())
	//避免mysql 5.7以後的datatime 存在000-000-000 錯誤
	user.LoginTime = time.Now().UTC()
	user.HeartbeatTime = time.Now().UTC()
	user.LogOutTime = time.Now().UTC()

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "該用戶已經存在",
		})
		return
	}
	if password != repassword {
		c.JSON(-1, gin.H{"msg": "密碼不一致"})
		return
	}
	//user.Password = password
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(200, gin.H{"msg": "新增用戶成功"})
}

// DeleteUser
// @Summary 刪除用戶
// @Tags 使用者
// @Param id query string false "ID"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// UpdateUser
// @Summary 修改使用者
// @Tags 使用者
// @Param id formData string false "ID"
// @Param name formData string false "Name"
// @Param password formData string false "Password"
// @Param email formData string false "Email"
// @Param phone formData string false "Phone"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	models.UpdateUser(user)
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
			"msg":   "valid error",
		})
	} else {
		fmt.Println("updated successfully", user)
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

// FindUserByNameAndPwd
// @Summary 登入使用者
// @Tags 使用者
// @Param name formData string false "Name"
// @Param password formData string false "Password"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	name := c.PostForm("name")
	password := c.PostForm("password")

	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    0,
			"message": "該用戶不存在",
		})
		return
	}
	fmt.Println(user)
	flag := utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"message": "密碼錯誤",
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改用戶成功",
		"data":    data,
	})

}

//防止跨域的偽造請球
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMessage(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("websocket upgrade error:", err)
		return
	}

	//一班使用defer ws.Close();使用回呼能多判斷錯誤
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println("websocket close error:", err)
			return
		}
	}(ws)
	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		fmt.Println("websocket error:", err)
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	tm := fmt.Sprintf("[ws][%s]:%s", msg, t)
	err = ws.WriteMessage(1, []byte(tm))
	if err != nil {
		fmt.Println("websocket send message error:", err)
	}
}
