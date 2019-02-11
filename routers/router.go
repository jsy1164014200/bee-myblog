// @APIVersion 1.0.0
// @Title bee-myblog API
// @Description myblog Api document
// @Contact gng@bingyan.net
// @TermsOfServiceUrl https://blog.jiangshiyi.top
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"bee-myblog/controllers"
	"bee-myblog/models"
	"bee-myblog/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/gomodule/redigo/redis"
)

func init() {
	authNS := beego.NewNamespace("/auth",
		beego.NSNamespace("/access_token",
			beego.NSInclude(&controllers.AccessTokenController{}),
		),
		beego.NSNamespace("/refresh_token",
			beego.NSInclude(&controllers.RefreshTokenController{}),
		),
	)
	ns := beego.NewNamespace("/v1",
		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Request.Method == "GET" {
				return true
			}
			access_token := ctx.Request.Header.Get("Authorization")
			claims, err := utils.DecodeJWT(access_token)
			if err != nil || claims["IssueType"].(string) != "AccessToken" {
				return false
			}
			key := claims["UserId"].(string) + "AccessToken"
			c := models.GetRedis()
			defer c.Close()
			token, err := redis.String(c.Do("get", key))
			if err != nil || token != access_token {
				return false
			}

			return true
		}),
		beego.NSNamespace("/blogs/:id",
			beego.NSInclude(&controllers.BlogController{}),
		),
		beego.NSNamespace("/blogs",
			beego.NSInclude(&controllers.BlogsController{}),
		),
		beego.NSNamespace("/archives/:id",
			beego.NSInclude(&controllers.ArchiveController{}),
		),
		beego.NSNamespace("/archives",
			beego.NSInclude(&controllers.ArchivesController{}),
		),
		beego.NSNamespace("/collections/:id",
			beego.NSInclude(&controllers.CollectionController{}),
		),
		beego.NSNamespace("/collections",
			beego.NSInclude(&controllers.CollectionsController{}),
		),
		beego.NSNamespace("/comments/:id",
			beego.NSInclude(&controllers.CommentController{}),
		),
		beego.NSNamespace("/comments",
			beego.NSInclude(&controllers.CommentsController{}),
		),
	)
	beego.AddNamespace(authNS)
	beego.AddNamespace(ns)
}
