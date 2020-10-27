package v1

import (
	"JalanJiang/douban-tool-server/models"
	"JalanJiang/douban-tool-server/pkg/e"
	"JalanJiang/douban-tool-server/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	UserID uint   `valid:"Required"`
	OpenID string `valid:"Required"`
}

// UserSession 微信用户会话信息
type UserSession struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"union_id"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

// Login 用户小程序登录
func Login(c *gin.Context) {
	// 小程序授权 code
	wxCode := c.PostForm("wx_code")
	fmt.Println(wxCode)

	// 校验数据
	valid := validation.Validation{}
	valid.Required(wxCode, "wx_code").Message("code 不能为空")

	resp, err := util.WxLogin(wxCode)
	// fmt.Println(resp)
	if err != nil {
		// 登录失败
		log.Println("小程序登录接口请求失败")
		util.ReturnError(c, http.StatusBadRequest, e.LOGIN_FAIL)
		return
	}

	// 解析用户会话信息
	var session UserSession
	jsonDecodeErr := json.Unmarshal(resp, &session)
	if jsonDecodeErr != nil {
		log.Printf("JSON decode failed: %v", jsonDecodeErr)
		util.ReturnError(c, http.StatusBadRequest, e.LOGIN_FAIL)
		return
	}

	// 判断返回是否错误
	if session.Errcode != 0 {
		// 发生错误
		log.Printf("调用微信登录接口错误：errcode=%d, errmsg=%s", session.Errcode, session.Errmsg)
		util.ReturnError(c, http.StatusInternalServerError, e.LOGIN_FAIL)
		return
	}

	// TODO 请求微信小程序授权接口，获取 openid, session_key, unionid
	// TODO 授权失败返回错误
	openID := session.OpenID
	unionID := session.UnionID
	secretKey := session.SessionKey

	// 查找是否存在用户
	result, user := models.GetUserByOpenID(openID)
	if result.RowsAffected == 0 {
		// 不存在，用户入库
		_, user = models.AddUser(openID, unionID, secretKey)
	}
	// 获取用户 ID
	userID := user.ID

	// 校验授权返回的数据
	a := auth{UserID: userID, OpenID: openID}
	fmt.Println(a)
	ok, _ := valid.Valid(&a)
	if !ok {
		// 打印错误数据
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		// 校验数据失败
		util.ReturnError(c, http.StatusBadRequest, e.LOGIN_FAIL)
		return
	}

	// 生成 Token
	token, err := util.GenerateToken(userID, openID)
	if err != nil {
		// Token 生成失败
		log.Println("Token 生成失败")
		util.ReturnError(c, http.StatusBadRequest, e.LOGIN_FAIL)
		return
	}

	// 返回 token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
