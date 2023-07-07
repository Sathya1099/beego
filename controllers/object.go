package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sathya1099/beego/models"
	"github.com/Sathya1099/beego/utils"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} success message
// @Failure 403 {string} error message
// @router / [post]
func (o *ObjectController) Post() {
	logs.Info("Creating New Object")
	o.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	var ob models.Object
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logs.Error("Error unmarshalling object")
		o.Data["json"] = map[string]string{"errMessage": "error while unmarshaling"}
		o.Ctx.Output.Status = http.StatusForbidden
	} else {
		if validateErrors := utils.Validate(&ob); validateErrors != nil {
			logs.Error(validateErrors)
			o.Data["json"] = map[string][]string{"errMessages": validateErrors}
			o.Ctx.Output.Status = http.StatusForbidden
		} else if err := models.Insert(&ob); err != nil {
			logs.Error(err)
			o.Data["json"] = map[string]string{"errMessage": err.Error()}
			o.Ctx.Output.Status = http.StatusForbidden
		} else {
			logs.Info("Object Created Successfully")
			o.Data["json"] = map[string]string{"message": "object created successfully."}
		}
	}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *ObjectController) Get() {
	logs.Info("Getting User-Specified Object")
	o.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId == "" {
		logs.Error("objectId is empty")
		o.Data["json"] = map[string]string{"errMessage": "objectId is empty"}
		o.Ctx.Output.Status = http.StatusForbidden
	} else {
		id, _ := strconv.Atoi(objectId)
		ob, err := models.Read(id)
		if err != nil {
			logs.Error(err)
			o.Data["json"] = map[string]string{"errMessage": err.Error()}
			o.Ctx.Output.Status = http.StatusForbidden
		} else {
			logs.Info("Specific Object Accessed Successfully")
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 {string} error message
// @router / [get]
func (o *ObjectController) GetAll() {
	logs.Info("Getting All Objects")
	o.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	obs, err := models.ReadAll()
	if err != nil {
		logs.Error(err)
		o.Data["json"] = map[string]string{"errMessage": err.Error()}
		o.Ctx.Output.Status = http.StatusForbidden
	} else {
		logs.Info("All Objects Accessed Successfully")
		o.Data["json"] = obs
	}
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 error messages
// @router / [put]
func (o *ObjectController) Put() {
	logs.Info("Updating User-Specified Object")
	o.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	var ob models.Object
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logs.Error("Error unmarshalling object")
		o.Data["json"] = map[string]string{"errMessage": "error while unmarshaling"}
		o.Ctx.Output.Status = http.StatusForbidden
	} else {
		if validateErrors := utils.ValidateForUpdate(&ob); validateErrors != nil {
			logs.Error(validateErrors)
			o.Data["json"] = map[string][]string{"errMessages": validateErrors}
			o.Ctx.Output.Status = http.StatusForbidden
		} else if err = models.Update(&ob); err != nil {
			logs.Error(err)
			o.Data["json"] = map[string]string{"errMessage": err.Error()}
			o.Ctx.Output.Status = http.StatusForbidden
		} else {
			logs.Info("object updated successfully.")
			o.Data["json"] = map[string]string{"message": "object updated successfully."}
		}
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *ObjectController) Delete() {
	logs.Info("Deleting User-Specified Object")
	o.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId == "" {
		logs.Error("objectId is empty")
		o.Data["json"] = map[string]string{"errMessage": "objectId is empty"}
		o.Ctx.Output.Status = http.StatusForbidden
	} else {
		id, _ := strconv.Atoi(objectId)
		err := models.Delete(id)
		if err != nil {
			logs.Error(err)
			o.Data["json"] = map[string]string{"errMessage": err.Error()}
			o.Ctx.Output.Status = http.StatusForbidden
		} else {
			logs.Info("object deleted successfully.")
			o.Data["json"] = map[string]string{"message": "object deleted successfully."}
		}
	}
	o.ServeJSON()
}
