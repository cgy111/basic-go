//go:build wireinject

// 让wire来注入这里的代码
package wire

import (
	"basic-go/wire/repository"
	"basic-go/wire/repository/dao"
	"github.com/google/wire"
)

func InitRepository() *repository.UserRepository {
	//我只在这里声明我要用的各种东西,但是具体怎么构造,怎么编排顺序
	//wire会帮我完成
	//这个方法里面传入各个组件的初始化方法
	wire.Build(repository.NewUserRepository,
		dao.NewUserDao, InitDB)
	return new(repository.UserRepository)
}
