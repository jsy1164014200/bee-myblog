package controllers

import (
	"bee-myblog/models"
	"bee-myblog/utils"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"github.com/gomodule/redigo/redis"
)

type AccessTokenController struct {
	beego.Controller
}

// @Title get access_token and refresh_token
// @Description get access_token and refresh_token
// @Param 	username 	body   string   true  "账号"
// @Param   password    body   string   true  "密码"
// @Success 201 {object} models.Token
// @Failure 403 body is empty
// @router / [post]
func (this *AccessTokenController) Post() {
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	handleError(this, err)
	username, err := js.Get("username").String()
	handleError(this, err)
	password, err := js.Get("password").String()
	handleError(this, err)
	if username != "jsy1164014200" && password != "sdj@##@kbl*&sdg#sdg@sdg$^sdssddf#" {
		this.Ctx.ResponseWriter.WriteHeader(401)
		this.Data["json"] = map[string]string{
			"error": "username or password error",
		}
		this.ServeJSON()
		return
	}
	// jwt  加入redis缓存
	token := utils.EncodeJWT()
	conn := models.GetRedis()
	defer conn.Close()
	_, err = conn.Do("set", "adminAccessToken", token.AccessToken)
	_, err = conn.Do("expire", "adminAccessToken", 60)
	handleError(this, err)
	_, err = conn.Do("set", "adminRefreshToken", token.RefreshToken)
	_, err = conn.Do("expire", "adminRefreshToken", 30*24*60*60)
	handleError(this, err)

	this.Ctx.ResponseWriter.WriteHeader(201)
	this.Data["json"] = token
	this.ServeJSON()
}

type RefreshTokenController struct {
	beego.Controller
}

// @Title get access_token and refresh_token
// @Description get access_token and refresh_token
// @Param 	refreshToken 	body   string   true  "refreshToken"
// @Success 201 {object} models.Token
// @Failure 403 body is empty
// @router / [post]
func (this *RefreshTokenController) Post() {
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	handleError(this, err)
	refreshToken, err := js.Get("refreshToken").String()
	handleError(this, err)

	claims, err := utils.DecodeJWT(refreshToken)
	handleError(this, err)
	userId := claims["UserId"].(string)
	c := models.GetRedis()
	defer c.Close()
	redisRefreshToken, err := redis.String(c.Do("get", userId+"RefreshToken"))
	if err != nil || redisRefreshToken != refreshToken {
		this.Ctx.ResponseWriter.WriteHeader(401)
		this.ServeJSON()
		return
	}
	token := utils.EncodeJWT()
	this.Data["json"] = token
	_, err = c.Do("set", "adminAccessToken", token.AccessToken)
	_, err = c.Do("expire", "adminAccessToken", 60)
	handleError(this, err)
	_, err = c.Do("set", "adminRefreshToken", token.RefreshToken)
	_, err = c.Do("expire", "adminRefreshToken", 30*24*60*60)
	handleError(this, err)

	this.Ctx.ResponseWriter.WriteHeader(201)
	this.ServeJSON()
}
