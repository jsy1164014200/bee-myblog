package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Options() {

}

func handleError(this interface{}, err error) {
	if err != nil {
		switch this.(type) {
		case *BlogsController:
			newThis := this.(*BlogsController)
			beego.Error(err)
			fmt.Println(err)
			newThis.Data["json"] = map[string]string{"error": "params error"}
			newThis.Ctx.ResponseWriter.WriteHeader(400)
			newThis.ServeJSON()
			newThis.StopRun()
		case *BlogController:
			newThis := this.(*BlogController)
			beego.Error(err)
			fmt.Println(err)
			newThis.Data["json"] = map[string]string{"error": "params error"}
			newThis.Ctx.ResponseWriter.WriteHeader(400)
			newThis.ServeJSON()
			newThis.StopRun()
		case *ArchivesController:
			newThis := this.(*ArchivesController)
			beego.Error(err)
			fmt.Println(err)
			newThis.Data["json"] = map[string]string{"error": "params error"}
			newThis.Ctx.ResponseWriter.WriteHeader(400)
			newThis.ServeJSON()
			newThis.StopRun()
		case *ArchiveController:
			newThis := this.(*ArchiveController)
			beego.Error(err)
			fmt.Println(err)
			newThis.Data["json"] = map[string]string{"error": "params error"}
			newThis.Ctx.ResponseWriter.WriteHeader(400)
			newThis.ServeJSON()
			newThis.StopRun()
		case *CollectionsController:
			newThis := this.(*CollectionsController)
			beego.Error(err)
			fmt.Println(err)
			newThis.Data["json"] = map[string]string{"error": "params error"}
			newThis.Ctx.ResponseWriter.WriteHeader(400)
			newThis.ServeJSON()
			newThis.StopRun()
		case *CollectionController:
			newThis := this.(*CollectionController)
			beego.Error(err)
			fmt.Println(err)
			newThis.Data["json"] = map[string]string{"error": "params error"}
			newThis.Ctx.ResponseWriter.WriteHeader(400)
			newThis.ServeJSON()
			newThis.StopRun()
		case *CommentsController:
			newThis := this.(*CommentsController)
			beego.Error(err)
			fmt.Println(err)
			newThis.Data["json"] = map[string]string{"error": "params error"}
			newThis.Ctx.ResponseWriter.WriteHeader(400)
			newThis.ServeJSON()
			newThis.StopRun()
		case *CommentController:
			newThis := this.(*CommentController)
			beego.Error(err)
			fmt.Println(err)
			newThis.Data["json"] = map[string]string{"error": "params error"}
			newThis.Ctx.ResponseWriter.WriteHeader(400)
			newThis.ServeJSON()
			newThis.StopRun()
		case *AccessTokenController:
			newThis := this.(*AccessTokenController)
			beego.Error(err)
			fmt.Println(err)
			newThis.Data["json"] = map[string]string{"error": "params error"}
			newThis.Ctx.ResponseWriter.WriteHeader(400)
			newThis.ServeJSON()
			newThis.StopRun()
		case *RefreshTokenController:
			newThis := this.(*RefreshTokenController)
			beego.Error(err)
			fmt.Println(err)
			newThis.Data["json"] = map[string]string{"error": "params error"}
			newThis.Ctx.ResponseWriter.WriteHeader(400)
			newThis.ServeJSON()
			newThis.StopRun()
		default:
			beego.Error("type transfer error")
			panic("type transfer error")
			return
		}

	}
}
