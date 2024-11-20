package dal

import (
	"github.com/leehaohaohao/go-mall/user/biz/dal/mysql"
	"github.com/leehaohaohao/go-mall/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
