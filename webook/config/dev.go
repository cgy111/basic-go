//go:build !k8s

// 没有k8s编译标签
package config

var Config = config{
	DB: DBConfig{
		//本地连接
		DSN: "root:root@tcp(47.123.5.217:13316)/webook",
	},
	Redis: RedisConfig{
		Addr: "47.123.5.217:6379",
	},
}
