// @APIVersion 1.0.0
// @Title Object API Example
// @Description an example of CRUD operations with object apis.
package routers

import (
	"github.com/Sathya1099/beego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
