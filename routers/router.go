// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"github.com/Sathya1099/beego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/object",
		beego.NSRouter("/", &controllers.ObjectController{}, "get:GetAll"),
		beego.NSRouter("/:objectId", &controllers.ObjectController{}, "get:Get"),
		beego.NSRouter("/", &controllers.ObjectController{}, "put:Put"),
		beego.NSRouter("/:objectId", &controllers.ObjectController{}, "delete:Delete"),
	)

	beego.AddNamespace(ns)
}
