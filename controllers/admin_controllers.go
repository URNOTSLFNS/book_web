package controllers

import (
	"strings"
	"test_web/models"

	_ "github.com/go-sql-driver/mysql"

	//"database/sql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) ShowDelete() {
	this.TplName = "admin_delete.tpl"
}
func (c *AdminController) Delete() {
	Id, err := c.GetInt64("bookId")
	if err != nil {
		panic(err)
		return
	} else {
		o := orm.NewOrm()
		user := models.BookInfo{}
		user.Id = Id
		_, err := o.Delete(&user)
		if err != nil {
			beego.Info("删除失败")
			c.Data["errmsg"] = "删除失败"
			return
		}
		beego.Info("删除成功", user)
		c.Data["errmsg"] = "删除成功"

	}
}

func (this *AdminController) ShowUpdate() {
	this.TplName = "admin_update.tpl"
}
func (c *AdminController) Update() {
	Id, err := c.GetInt64("Id")
	c.Data["errmsg"] = Id
	if err != nil {
		panic(err)
		return
	} else {
		update_name := strings.TrimSpace(c.GetString("update_name"))
		update_category := strings.TrimSpace(c.GetString("update_category"))
		update_writer := strings.TrimSpace(c.GetString("update_writer"))
		o := orm.NewOrm()
		user := models.BookInfo{}

		user.Id = Id

		err := o.Read(&user)
		if err == nil {
			user.Book_name = update_name
			user.Book_writer = update_writer
			user.Book_type = update_category
			_, err := o.Update(&user)
			if err != nil {
				beego.Info("更新失败")
			}
		}
		beego.Info(user)
		c.Data["errmsg"] = user
	}
}
