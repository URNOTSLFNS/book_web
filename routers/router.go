package routers

import (
	"test_web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admindelete", &controllers.AdminController{}, "get:ShowDelete;post:Delete")
	beego.Router("/update", &controllers.AdminController{}, "get:ShowUpdate;post:Update")
	beego.Router("/search", &controllers.SearchController{}, "get:ShowSearch;post:Search")
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegister;post:HandleRegister")
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/delete", &controllers.ShelfController{}, "get:ShowDelete;post:Delete")
	beego.Router("/insert", &controllers.ShelfController{}, "get:ShowInsert;post:Insert")
	beego.Router("/shelf", &controllers.ShelfController{}, "get:MyShelf")
	beego.Router("/rank", &controllers.RankController{}, "get:Rank")
	beego.Router("/classify", &controllers.ClassifyController{}, "get:Show;post:Classify")
}
