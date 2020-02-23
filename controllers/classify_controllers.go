package controllers

import (
	"encoding/json"
	"fmt"

	//	"io/ioutil"
	//"net/http"

	//"strings"

	//"database/sql"
	//	"test_web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ClassifyController struct {
	beego.Controller
}

func (c *ClassifyController) Show() {
	c.TplName = "sort.tpl"
}

func (c *ClassifyController) Classify() []orm.Params {
	c.TplName = "sort.tpl"
	classify := c.GetString("book_type")
	fmt.Printf("ooo" + classify + "ooo")
	o := orm.NewOrm()
	var maps []orm.Params
	_, _ = o.Raw("select * from book_info where book_type like ?", "%"+classify+"%").Values(&maps)

	b, err := json.Marshal(maps)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)

	}

	//fmt.Println(maps)
	beego.Info(string(b))
	//b[book_name]
	c.Data["errmsg"] = (string(b))
	return maps
}
