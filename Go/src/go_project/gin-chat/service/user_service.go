package service

import (
	"fmt"
	"gin-chat/models"
	"gin-chat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

// GetUserList
// @Summary 所有用户
// @Tags 首页
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}

// FindUserByNameAndPwd
// @Summary 登录
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message","data"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}

	name := c.Query("name")
	password := c.Query("password")

	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"message": "该用户不存在",
		})
		return
	}

	flag := utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "密码不正确",
			"data":    data,
		})
		return
	}

	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)

	c.JSON(200, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param re_password query string false "确认密码"
// @param phone query string false "手机号"
// @param email query string false "邮件"
// @param identify query string false "身份标识"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	var user models.UserBasic
	user.Name = c.Query("name")
	password := c.Query("password")
	rePassWord := c.Query("re_password")

	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt

	data := models.FindUserByName(user.Name)
	fmt.Println("data is ", data)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"message": "用户名已经注册",
		})
		return
	}

	if password != rePassWord {
		c.JSON(200, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	//user.Password = password
	user.Password = utils.MakePassword(password, salt)
	user.Phone = c.Query("phone")
	user.Email = c.Query("email")
	fmt.Println("create:", user)

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "格式不正确",
		})
		return
	} else {
		models.CreateUser(user)

		c.JSON(200, gin.H{
			"message": "新增用户成功",
		})
	}

}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [post]
func DeleteUser(c *gin.Context) {
	var user models.UserBasic
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)

	models.DeleteUser(user)

	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	var user models.UserBasic
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)

	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	models.UpdateUser(user)

	c.JSON(200, gin.H{
		"message": "更新用户成功",
	})
}
