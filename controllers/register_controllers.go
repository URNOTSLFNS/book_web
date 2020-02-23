package controllers

import (
	"strings"
	"test_web/models"

	"crypto/md5"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

// 注册展示页面
func (this *RegisterController) ShowRegister() {
	this.TplName = "register.tpl"
}

// 注册获取数据页面
func (this *RegisterController) HandleRegister() {
	// 获取浏览器传递的值，并去除两边的空格

	Name := strings.TrimSpace(this.GetString("userName"))
	Pwd := strings.TrimSpace(this.GetString("passWord"))
	//beego.Info("账号:", Name, "密码:", Pwd)
	// 数据处理
	if Name == "" || Pwd == "" {
		beego.Info("用户名或密码不能为空")
		this.TplName = "register.tpl"
		this.Data["errmsg"] = "用户名或密码不能为空 ！"
		return
	}
	// 插入数据库（数据库表，Users）
	//获取orm对象
	o := orm.NewOrm()
	//   获取插入对象
	user := models.Users{}
	//md5

	data := []byte(Pwd)

	has := md5.Sum(data)

	md5str1 := fmt.Sprintf("%x", has)
	//   插入数值
	user.Name = Name
	user.Pwd = md5str1

	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败，用户相同或者其他错误！！！")
		this.TplName = "register.tpl"
		this.Data["errmsg"] = "插入数据失败，用户相同或者其他错误！！！！"
		return
	}
	//this.Data["errmsg"] = "插入数据失败，用户相同或者其他错误！！！！"
	// 测试返回视图
	//this.Ctx.WriteString("注册成功！！！")
	// 实际情况注册成功返回到登录页面
	this.Redirect("login", 302)
}
