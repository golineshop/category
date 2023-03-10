package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (conf config.Config, err error) {
	consulSource := consul.NewSource(
		// 设置配置中心地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// 设置前缀，不设置的默认前缀是/micro/config
		consul.WithPrefix(prefix),
		// 是否移除前缀，这里设置为true 表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)
	// 配置初始化
	conf, err = config.NewConfig()
	if err != nil {
		return conf, err
	}
	// 加载配置
	err = conf.Load(consulSource)
	return conf, err
}
