package controllers

import (
	"strings"
	"test_web/models"

	"crypto/md5"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

// 登录页面 get
func (this *LoginController) ShowLogin() {
	this.TplName = "index.tpl"
}

// 登录页面 post
func (this *LoginController) HandleLogin() {
	// 拿到浏览器数据，并去除两边的空格
	Select := strings.TrimSpace(this.GetString("select"))
	Name := strings.TrimSpace(this.GetString("userName"))
	Pwd := strings.TrimSpace(this.GetString("passWord"))
	this.Data["errmsg"] = Name
	//this.Ctx.WriteString(Type)

	//数据处理
	data := []byte(Pwd)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	beego.Info("账号:", Name, "密码:", md5str1)

	if Name == "" || Pwd == "" {
		beego.Info("登录失败")
		this.TplName = "index.tpl"
		this.Data["errmsg"] = "登录失败"
		return
	}
	o := orm.NewOrm()

	if Select == "1" {
		//普通用户登录
		var user models.Users

		// 查询
		user.Name = Name
		err := o.Read(&user, "Name")
		if err != nil {
			beego.Info("用户名登录失败")
			this.TplName = "index.tpl"
			this.Data["errmsg"] = "用户名登录失败"
			//beego.Info("0")
			return
		}
		// 判断密码是否一致
		//this.Ctx.WriteString(user.Pwd + "|")
		//this.Ctx.WriteString(md5str1)
		if user.Pwd != md5str1 {
			beego.Info("密码登录失败")
			this.TplName = "index.tpl"
			this.Data["errmsg"] = "密码登录失败"
			//beego.Info("0")
			return
		}
		this.Redirect("", 302)
		this.TplName = "index.tpl"
		this.Data["errmsg"] = "1"
		beego.Info("1")
		this.Data["errmsg"] = user
		//		UserId := user.Id
		// 测试返回视图

		this.Ctx.WriteString("普通用户登录成功")
		// 实际情况注册成功返回到主页面

	} else if Select == "2" {
		//管理员登录
		var user models.Admin
		// 查询
		user.Name = Name
		err := o.Read(&user, "Name")
		if err != nil {
			beego.Info("用户名登录失败")
			this.TplName = "index.tpl"
			this.Data["errmsg"] = "用户名登录失败"
			//beego.Info("0")
			return
		}
		// 判断密码是否一致
		//this.Ctx.WriteString(user.Pwd + "|")
		//this.Ctx.WriteString(md5str1)
		if user.Pwd != md5str1 {
			beego.Info("密码登录失败！")
			this.TplName = "index.tpl"
			this.Data["errmsg"] = "密码登录失败！"
			//beego.Info("0")
			return
		}
		// 测试返回视图
		this.Redirect("", 302)
		this.Data["errmsg"] = "1"
		beego.Info("1")
		this.Ctx.WriteString("管理员登录成功")

		// 实际情况注册成功返回到登录页面

	}

}
