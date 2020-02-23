package controllers

import (
	//"strings"
	"test_web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SearchController struct {
	beego.Controller
}

func (c *SearchController) ShowSearch() {
	c.TplName = "search.html"
}
func (c *SearchController) Search() {
	Id, err := c.GetInt64("bookId")
	if err != nil {
		panic(err)
	} else {
		//Name := strings.TrimSpace(c.GetString("bookName"))
		o := orm.NewOrm()
		user := models.BookInfo{}
		user.Id = Id
		err := o.Read(&user, "Id")
		if err != nil {
			beego.Info("查询失败")
			c.TplName = "sort.html"
			c.Data["errmsg"] = "查询失败"
			return
		}

		c.Data["errmsg"] = user
		//c.Ctx.WriteString("书名:" + user.Book_name + " 作者：" + user.Book_writer)
		c.TplName = "search.html"
	}
}
