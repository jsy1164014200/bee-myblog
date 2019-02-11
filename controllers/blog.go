package controllers

import (
	"bee-myblog/models"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type BlogsController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all blog
// @Success 200 {object} []models.Blog
// @Param tag query string false "获取指定标签的博文"
// @Param sort query string false "根据指定值排序，比如时间，阅读数，评论数"
// @Param offset query int false "分页值，如果有offset就一定要用limit"
// @Param limit query int false "分页"
// @Failure 400 params error
// @router / [get]
func (this *BlogsController) Get() {
	conn := models.GetDB()
	blogCollection := conn.DB("beeblog").C("blog")
	defer conn.Close()

	offsetStr := this.GetString("offset")
	limitStr := this.GetString("limimt")
	sortStr := this.GetString("sort")
	tagStr := this.GetString("tag")

	var blogs []models.Blog
	var queryPoint *mgo.Query
	if tagStr != "" {
		queryPoint = blogCollection.Find(bson.M{"tags": tagStr})
	} else {
		queryPoint = blogCollection.Find(nil)
	}

	if sortStr != "" {
		queryPoint = queryPoint.Sort(sortStr)
	}

	err := queryPoint.All(&blogs)
	handleError(this, err)
	if offsetStr != "" && limitStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		handleError(this, err)
		limit, err := strconv.Atoi(limitStr)
		handleError(this, err)
		blogs = blogs[offset : offset+limit]
	}

	this.Data["json"] = blogs
	this.ServeJSON()
}

// @Title CreateBlog
// @Description create one blog
// @Param	title		body 	string	true		"文章标题"
// @Param   summary     body   string   true  "文章概要"
// @Param 	tags 		body   []string true  "文章标签"
// @Param 	archive 	body   string   true  "文章归档"
// @Success 201 {object} models.Blog
// @Failure 403 body is empty
// @router / [post]
func (this *BlogsController) Post() {
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	handleError(this, err)
	title, err := js.Get("title").String()
	handleError(this, err)
	summary, err := js.Get("summary").String()
	handleError(this, err)
	tags, err := js.Get("tags").StringArray()
	handleError(this, err)
	archive, err := js.Get("archive").String()
	handleError(this, err)

	newId := bson.NewObjectId()
	blog := models.Blog{
		Id:           newId,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Title:        title,
		Summary:      summary,
		CommentCount: 0,
		Comments:     []bson.ObjectId{},
		ReadCount:    0,
		Tags:         tags,
	}
	conn := models.GetDB()
	defer conn.Close()
	err = conn.DB("beeblog").C("blog").Insert(&blog)
	handleError(this, err)
	err = conn.DB("beeblog").C("archive").Update(bson.M{"name": archive},
		bson.M{"$push": bson.M{"blogs": newId}})
	handleError(this, err)
	this.Ctx.ResponseWriter.WriteHeader(201)
	this.Data["json"] = blog
	this.ServeJSON()
}

// 处理one blog
type BlogController struct {
	beego.Controller
}

// @Title Get
// @Description get one blog by ObjectId
// @Param	id		path 	string	true		"blog id"
// @Success 200 {object} models.Blog
// @Failure 403 :uid is empty
// @router / [get]
func (this *BlogController) Get() {
	if !bson.IsObjectIdHex(this.GetString(":id")) {
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = map[string]interface{}{"error": "id error"}
		this.ServeJSON()
		return
	}
	objectId := bson.ObjectIdHex(this.GetString(":id"))
	conn := models.GetDB()
	defer conn.Close()
	var blog []models.Blog
	err := conn.DB("beeblog").C("blog").Find(bson.M{"_id": objectId}).All(&blog)
	handleError(this, err)
	this.Data["json"] = blog[0]
	this.ServeJSON()
}

// @Title Update
// @Description update the blog
// @Param	id		path 	string	true		"The id you want to update"
// @Param	title		body 	string	true		"标题（改了就有值，没改就为空)"
// @Param 	summary 	body 	string  true 		"总结（同上）"
// @Param   tags        body    []string true       "标签(同上)"
// @Param   archive		body    string   true       "归档信息(同上)"
// @Success 200 {object} models.User
// @Failure 403 id not right
// @router / [put]
func (this *BlogController) Put() {
	if !bson.IsObjectIdHex(this.GetString(":id")) {
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = map[string]interface{}{"error": "id error"}
		this.ServeJSON()
		return
	}
	objectId := bson.ObjectIdHex(this.GetString(":id"))
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	handleError(this, err)
	title, err := js.Get("title").String()
	handleError(this, err)
	summary, err := js.Get("summary").String()
	handleError(this, err)
	tags, err := js.Get("tags").StringArray()
	handleError(this, err)
	archive, err := js.Get("archive").String()
	handleError(this, err)

	conn := models.GetDB()
	defer conn.Close()
	err = conn.DB("beeblog").C("blog").Update(bson.M{"_id": objectId},
		bson.M{"$set": bson.M{"title": title,
			"summary": summary,
			"tags":    tags}})
	handleError(this, err)

	err = conn.DB("beeblog").C("archive").Update(bson.M{"blogs": objectId}, bson.M{
		"$pull": bson.M{"blogs": objectId},
	})
	handleError(this, err)
	err = conn.DB("beeblog").C("archive").Update(bson.M{"name": archive}, bson.M{
		"$push": bson.M{"blogs": objectId},
	})
	handleError(this, err)

	var blog []models.Blog
	err = conn.DB("beeblog").C("blog").Find(bson.M{"_id": objectId}).All(&blog)
	handleError(this, err)
	this.Data["json"] = blog[0]
	this.Ctx.ResponseWriter.WriteHeader(201)
	this.ServeJSON()
}

// @Title Delete
// @Description delete the blog
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204 {object} {}
// @Failure 403 id is empty
// @router / [delete]
func (this *BlogController) Delete() {
	if !bson.IsObjectIdHex(this.GetString(":id")) {
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = map[string]interface{}{"error": "id error"}
		this.ServeJSON()
		return
	}
	objectId := bson.ObjectIdHex(this.GetString(":id"))
	conn := models.GetDB()
	defer conn.Close()
	err := conn.DB("beeblog").C("blog").Remove(bson.M{"_id": objectId})
	handleError(this, err)
	this.Ctx.ResponseWriter.WriteHeader(204)
	this.ServeJSON()
}
