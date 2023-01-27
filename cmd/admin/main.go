// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/quarkcms/quark-go/pkg/adapter/hertzadapter"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin"
	"github.com/quarkcms/quark-go/pkg/app/install"
	"github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	adminprovider "github.com/quarkcms/quark-hertz/cmd/admin/biz/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	h := server.Default(server.WithHostPorts(":3000"))

	// 注册路由
	register(h)

	// 加载html模板
	h.LoadHTMLGlob("../../web/template/*")

	// 静态文件
	h.StaticFile("/admin/", "../../website/admin/index.html")

	// 静态文件目录
	fs := &app.FS{Root: "../../website", IndexNames: []string{"index.html"}}
	h.StaticFS("/", fs)

	// 数据库配置信息
	dsn := "root:Bc5HQFJc4bLjZCcC@tcp(127.0.0.1:3306)/quarkgo?charset=utf8&parseTime=True&loc=Local"

	// 配置资源
	config := &builder.Config{
		AppKey:    "123456",
		Providers: append(admin.Providers, adminprovider.Providers...),
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
		StaticPath: "../../website",
	}

	// 创建对象
	b := builder.New(config)

	// 初始化安装
	b.Use(install.Handle)

	// 中间件
	b.Use(middleware.Handle)

	// 适配hertz
	hertzadapter.Adapter(b, h)

	// 启动服务
	h.Spin()
}
