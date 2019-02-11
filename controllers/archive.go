package controllers

import (
	"bee-myblog/models"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"gopkg.in/mgo.v2/bson"
)

type ArchivesController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all archives
// @Success 200 {object} []models.Archive
// @Failure 400 params error
// @router / [get]
func (this *ArchivesController) Get() {
	conn := models.GetDB()
	defer conn.Close()
	var archives []models.Archive
	err := conn.DB("beeblog").C("archive").Find(nil).All(&archives)
	handleError(this, err)
	result := make([]map[string]interface{}, len(archives))
	for index, value := range archives {
		result[index] = make(map[string]interface{})
		result[index]["_id"] = value.Id
		result[index]["name"] = value.Name
		var blogs []models.Blog
		err = conn.DB("beeblog").C("blog").Find(bson.M{"_id": bson.M{"$in": value.Blogs}}).All(&blogs)
		handleError(this, err)
		result[index]["blogs"] = blogs
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title CreateArchive
// @Description create one archive
// @Param 	name 	body   string   true  "文章归档"
// @Success 201 {object} models.Archive
// @Failure 403 body is empty
// @router / [post]
func (this *ArchivesController) Post() {
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	handleError(this, err)
	name, err := js.Get("name").String()
	handleError(this, err)

	conn := models.GetDB()
	defer conn.Close()

	archive := models.Archive{
		Id:    bson.NewObjectId(),
		Name:  name,
		Blogs: []bson.ObjectId{},
	}
	err = conn.DB("beeblog").C("archive").Insert(&archive)
	this.Ctx.ResponseWriter.WriteHeader(201)
	this.Data["json"] = archive
	this.ServeJSON()

}

type ArchiveController struct {
	beego.Controller
}

// @Title Get one archive
// @Description get one archive by ObjectId
// @Param	id		path 	string	true		"archive id"
// @Success 200 {object} models.Archive
// @Failure 403 :uid is empty
// @router / [get]
func (this *ArchiveController) Get() {
	if !bson.IsObjectIdHex(this.GetString(":id")) {
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = map[string]interface{}{"error": "id error"}
		this.ServeJSON()
		return
	}
	objectId := bson.ObjectIdHex(this.GetString(":id"))
	conn := models.GetDB()
	defer conn.Close()
	var archive []models.Archive
	err := conn.DB("beeblog").C("archive").Find(bson.M{"_id": objectId}).All(&archive)
	handleError(this, err)

	result := make([]map[string]interface{}, len(archive))
	for index, value := range archive {
		result[index] = make(map[string]interface{})
		result[index]["_id"] = value.Id
		result[index]["name"] = value.Name
		var blogs []models.Blog
		err = conn.DB("beeblog").C("blog").Find(bson.M{"_id": bson.M{"$in": value.Blogs}}).All(&blogs)
		handleError(this, err)
		result[index]["blogs"] = blogs
	}

	this.Data["json"] = result[0]
	this.ServeJSON()
}
