package models

import (
	//	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 用户信息
type Users struct {
	Id   int
	Name string `orm:"unique"` // 用户名唯一
	Pwd  string
}
type Admin struct {
	//如果field名称为Id，而且类型为int64，并没有定义tag，则会被xorm视为主键，并且拥有自增属性
	Id   int64  `xorm:"pk autoincr" json:"id"` //主键 自增
	Name string `xorm:"varchar(32)" json:"name"`
	Pwd  string `xorm:"varchar(255)" json:"pwd"` //管理员密码

}
type BookInfo struct {
	Id          int64
	Book_type   string
	Book_name   string
	Book_writer string
}
type BookShelf struct {
	Id          int64 `orm:"column(uid);pk"`
	Book_id     int64
	User_id     int64
	Book_name   string
	Book_writer string
}

func init() {
	//设置数据库连接信息
	orm.RegisterDataBase("default", "mysql", "chrislee:19871987@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local")
	// 映射modle数据
	orm.RegisterModel(new(Users))
	orm.RegisterModel(new(Admin))
	orm.RegisterModel(new(BookInfo))
	orm.RegisterModel(new(BookShelf))
	// 生成表，第二个false要是改成true 就会强制更新表，数据全部丢失
	orm.RunSyncdb("default", false, true)

}
