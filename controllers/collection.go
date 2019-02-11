package controllers

import (
	"bee-myblog/models"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"gopkg.in/mgo.v2/bson"
	btime "time"
)

type CollectionsController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all collections
// @Success 200 {object} []models.Collection
// @Failure 400 params error
// @router / [get]
func (this *CollectionsController) Get() {
	conn := models.GetDB()
	defer conn.Close()
	var collections []models.Collection
	err := conn.DB("beeblog").C("collection").Find(nil).All(&collections)
	handleError(this, err)
	this.Data["json"] = collections
	this.ServeJSON()
}

// @Title CreateCollection
// @Description create one collection
// @Param 	title 	body   string   true  "收藏内容的标题"
// @Param   author  body   string 	true  "收藏的作者"
// @Param  url   body   string  true  "收藏的url"
// @Param  time  body   string  true "收藏的时间"
// @Success 201 {object} models.Collection
// @Failure 403 body is empty
// @router / [post]
func (this *CollectionsController) Post() {
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	handleError(this, err)
	title, err := js.Get("title").String()
	handleError(this, err)
	author, err := js.Get("author").String()
	handleError(this, err)
	url, err := js.Get("url").String()
	handleError(this, err)
	//time, err := js.Get("time").String()
	//handleError(this, err)

	conn := models.GetDB()
	defer conn.Close()

	collection := models.Collection{
		Id:     bson.NewObjectId(),
		Title:  title,
		Author: author,
		Url:    url,
		Time:   btime.Now(),
	}
	err = conn.DB("beeblog").C("collection").Insert(&collection)
	this.Ctx.ResponseWriter.WriteHeader(201)
	this.Data["json"] = collection
	this.ServeJSON()

}

type CollectionController struct {
	beego.Controller
}

// @Title Get one collection
// @Description get one collection by ObjectId
// @Param	id		path 	string	true		"collection id"
// @Success 200 {object} models.Collection
// @Failure 403 :uid is empty
// @router / [get]
func (this *CollectionController) Get() {
	if !bson.IsObjectIdHex(this.GetString(":id")) {
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = map[string]interface{}{"error": "id error"}
		this.ServeJSON()
		return
	}
	objectId := bson.ObjectIdHex(this.GetString(":id"))
	conn := models.GetDB()
	defer conn.Close()
	var collection []models.Collection
	err := conn.DB("beeblog").C("collection").Find(bson.M{"_id": objectId}).All(&collection)
	handleError(this, err)

	this.Data["json"] = collection[0]
	this.ServeJSON()
}
