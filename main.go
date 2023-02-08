package main

import (
	"github.com/golineshop/category/common"
	"github.com/golineshop/category/domain/repository"
	"github.com/golineshop/category/domain/service"
	"github.com/golineshop/category/handler"
	proto "github.com/golineshop/category/proto"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/prometheus/common/log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 设置配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 设置注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// 服务参数设置
	serv := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		//这里设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		//添加consul 作为注册中心
		micro.Registry(consulRegistry),
	)

	// 获取mysql配置，路径中不带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	// 连接数据库
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			println(err)
		}
	}(db)
	// 禁止复数表
	db.SingularTable(true)

	//rp := repository.NewCategoryRepository(db)
	//rp.InitTable()

	// 初始化服务
	serv.Init()
	categoryService := service.NewCategoryService(repository.NewCategoryRepository(db))
	err = proto.RegisterCategoryHandler(serv.Server(), &handler.CategoryController{CategoryService: categoryService})
	if err != nil {
		log.Error(err)
	}

	// 运行服务
	if err := serv.Run(); err != nil {
		log.Fatal(err)
	}
}
