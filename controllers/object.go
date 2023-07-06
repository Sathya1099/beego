package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Sathya1099/beego/models"
	"github.com/Sathya1099/beego/utils"

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
	var ob models.Object
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Data["json"] = map[string]string{"errMessage": "error while unmarshaling"}
		o.Ctx.Output.Status = http.StatusForbidden
	} else {
		if validateErrors := utils.Validate(&ob); validateErrors != nil {
			o.Data["json"] = map[string][]string{"errMessages": validateErrors}
			o.Ctx.Output.Status = http.StatusForbidden
		} else if err := models.Insert(&ob); err != nil {
			o.Data["json"] = map[string]string{"errMessage": err.Error()}
			o.Ctx.Output.Status = http.StatusForbidden
		} else {
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
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId == "" {
		o.Data["json"] = map[string]string{"errMessage": "objectId is empty"}
		o.Ctx.Output.Status = http.StatusForbidden
	} else {
		ob, err := models.Read(objectId)
		if err != nil {
			o.Data["json"] = map[string]string{"errMessage": err.Error()}
			o.Ctx.Output.Status = http.StatusForbidden
		} else {
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
	obs, err := models.ReadAll()
	if err != nil {
		o.Data["json"] = map[string]string{"errMessage": err.Error()}
		o.Ctx.Output.Status = http.StatusForbidden
	} else {
		o.Data["json"] = obs
	}
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 error messages
// @router / [put]
func (o *ObjectController) Put() {
	var ob models.Object
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Data["json"] = map[string]string{"errMessage": "error while unmarshaling"}
		o.Ctx.Output.Status = http.StatusForbidden
	} else {
		if validateErrors := utils.Validate(&ob); validateErrors != nil {
			o.Data["json"] = map[string][]string{"errMessages": validateErrors}
			o.Ctx.Output.Status = http.StatusForbidden
		} else if err = models.Update(&ob); err != nil {
			o.Data["json"] = map[string]string{"errMessage": err.Error()}
			o.Ctx.Output.Status = http.StatusForbidden
		} else {
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
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId == "" {
		o.Data["json"] = map[string]string{"errMessage": "objectId is empty"}
		o.Ctx.Output.Status = http.StatusForbidden
	} else {
		err := models.Delete(objectId)
		if err != nil {
			o.Data["json"] = map[string]string{"errMessage": err.Error()}
			o.Ctx.Output.Status = http.StatusForbidden
		} else {
			o.Data["json"] = map[string]string{"message": "object deleted successfully."}
		}
	}
	o.ServeJSON()
}