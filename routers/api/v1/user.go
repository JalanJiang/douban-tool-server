package v1

import (
	"JalanJiang/douban-tool-server/pkg/e"
	"JalanJiang/douban-tool-server/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	userID uint `valid:"Required"`
	openID string `valid:"Required"`
}

// Login 用户小程序登录
func Login(c *gin.Context) {
	// 小程序授权 code
	wxCode := c.PostForm("wx_code")

	// 校验数据
	valid := validation.Validation{}
	valid.Required(wxCode, "wx_code").Message("code 不能为空")

	// TODO 请求微信小程序授权接口，获取 openid, session_key, unionid
	// TODO 授权失败返回错误
	openID := ""
	// unionID := ""
	// secretKey := ""
	// TODO 通过 openID 获取 userID
	userID := 1

	// 校验授权返回的数据
	a := auth{userID: uint(userID), openID: openID}
	ok, _ := valid.Valid(&a)
	if !ok {
		// 打印错误数据
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		// 校验数据失败
		util.ReturnError(c, http.StatusBadRequest, e.LOGIN_FAIL)
	}

	// 生成 Token
	token, err := util.GenerateToken(uint(userID), openID)
	if err != nil {
		// Token 生成失败
		log.Println("Token 生成失败")
		util.ReturnError(c, http.StatusBadRequest, e.LOGIN_FAIL)
	}

	// 返回 token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
