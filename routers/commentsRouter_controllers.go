package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bee-myblog/controllers:AccessTokenController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:AccessTokenController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:ArchiveController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:ArchiveController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:ArchivesController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:ArchivesController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:ArchivesController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:ArchivesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:BlogController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:BlogController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:BlogController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:BlogsController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:BlogsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:BlogsController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:BlogsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:CollectionController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:CollectionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:CollectionsController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:CollectionsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:CollectionsController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:CollectionsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:CommentController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:CommentsController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:CommentsController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-myblog/controllers:RefreshTokenController"] = append(beego.GlobalControllerRouter["bee-myblog/controllers:RefreshTokenController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
