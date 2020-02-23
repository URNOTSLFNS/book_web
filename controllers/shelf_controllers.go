package controllers

import (
	"encoding/json"
	//"database/sql"
	"fmt"
	//"strings"
	"test_web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ShelfController struct {
	beego.Controller
}

//显示我的书架
func (c *ShelfController) MyShelf() []orm.Params {

	userId := "0"
	o := orm.NewOrm()
	var maps []orm.Params
	_, _ = o.Raw("select * from book_shelf where user_id like ?", "%"+userId+"%").Values(&maps)

	b, err := json.Marshal(maps)

	if err != nil {
		fmt.Println("json.Marshal failed:", err)

	}
	fmt.Println(maps)
	c.TplName = "shelf.tpl"
	beego.Info(maps)
	c.Data["errmsg"] = string(b)
	return maps
}

func (c *ShelfController) ShowDelete() {

	c.TplName = "user_delete.tpl"
}

//从我的书架中删除书本
func (c *ShelfController) Delete() {
	bookId, err := c.GetInt64("bookId")

	if err != nil {
		panic(err)
		return
	} else {
		//		db, err := gorm.Open("mysql", "chrislee:19871987@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local")
		if err != nil {
			panic("连接数据库失败")
		}
		//		defer db.Close()

		// 自动迁移模式
		//db.AutoMigrate(&User{}, &Car{})

		o := orm.NewOrm()
		book := models.BookShelf{}
		book.Book_id = bookId

		//		db.Where("user_id = ? and book_id = ?", userId, bookId).Find(&book)
		_, err := o.Delete(&book)

		if err != nil {
			beego.Info("删除失败")
			c.Data["errmsg"] = "删除失败"
			return
		}
		beego.Info("删除成功", book)
		c.Data["errmsg"] = "删除成功"

	}
	c.TplName = "user_delete.tpl"
}

func (c *ShelfController) ShowInsert() {

	c.TplName = "user_insert.tpl"
}

//把书本加入我的书架
func (c *ShelfController) Insert() {
	Id, err := c.GetInt64("bookId")
	if err != nil {
		panic(err)
		return
	} else {

		o := orm.NewOrm()
		book := models.BookInfo{}
		user := models.BookShelf{}
		book.Id = Id
		err = o.Read(&book, "id")
		if err != nil {
			fmt.Println("o.Read err:", err)
		} else {
			user.Book_id = book.Id
			user.Book_name = book.Book_name
			user.Book_writer = book.Book_writer
			//插入操作
			_, err := o.Insert(&user) //id表示插入的数据的主键
			if err != nil {
				fmt.Println("o.Insert err:", err)
				return
			}
			c.Data["errmsg"] = book

			c.TplName = "user_insert.tpl"
		}
	}
}
