package controllers

import (
	//"test_web/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"

	//	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego/orm"
)

type RankController struct {
	beego.Controller
}

func (c *RankController) Rank() []orm.Params {

	o := orm.NewOrm()
	var maps []orm.Params
	_, _ = o.Raw("select * from book_info").Values(&maps)
	c.TplName = ("index.tpl")

	b, err := json.Marshal(maps)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)

	}
	fmt.Println(string(b))
	beego.Info(string(b))
	//c.Data["errmsg"] = string(b)
	return maps

}
