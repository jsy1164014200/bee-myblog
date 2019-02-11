package controllers

import (
	"bee-myblog/models"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type CommentsController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all comments
// @Success 200 {object} []models.Comment
// @Failure 400 params error
// @router / [get]
func (this *CommentsController) Get() {

}

// @Title CreateComment
// @Description create one comment
// @Param blogId body string true  "文章ID"
// @Param 	username 	body   string   true  "评论者"
// @Param   content  body   string 	true  "评论的内容"
// @Success 201 {object} models.Comment
// @Failure 403 body is empty
// @router / [post]
func (this *CommentsController) Post() {
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	handleError(this, err)
	blogIdStr, err := js.Get("blogId").String()
	handleError(this, err)
	if !bson.IsObjectIdHex(blogIdStr) {
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = map[string]interface{}{"error": "id error"}
		this.ServeJSON()
		return
	}

	blogId := bson.ObjectIdHex(blogIdStr)
	username, err := js.Get("username").String()
	handleError(this, err)
	content, err := js.Get("content").String()
	handleError(this, err)

	conn := models.GetDB()
	defer conn.Close()
	commentId := bson.NewObjectId()
	comment := models.Comment{
		Id:        commentId,
		Username:  username,
		Content:   content,
		CreatedAt: time.Now(),
	}

	err = conn.DB("beeblog").C("comment").Insert(&comment)
	handleError(this, err)
	err = conn.DB("beeblog").C("blog").Update(bson.M{"_id": blogId}, bson.M{
		"$push": bson.M{"comments": commentId},
		"$inc":  bson.M{"commentCount": 1},
	})
	handleError(this, err)

	this.Ctx.ResponseWriter.WriteHeader(201)
	this.Data["json"] = comment
	this.ServeJSON()
}

type CommentController struct {
	beego.Controller
}

// @Title Get one comment
// @Description get one comment by ObjectId
// @Param	id		path 	string	true		"collection id"
// @Success 200 {object} models.Collection
// @Failure 403 :uid is empty
// @router / [get]
func (this *CommentController) Get() {

}
