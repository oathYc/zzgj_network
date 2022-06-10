package boot

import (
	"github.com/pkg/errors"
)

func Boot() error {
	var ret error
	for ok := true; ok; ok = false {
		// 初始化配置
		if err := InitConfig(); nil != err {
			ret = errors.Wrap(err, "初始化配置出错")
			break
		}

		// 初始化http服务
		if err := InitHttpServer(); err != nil {
			ret = errors.Wrap(err, "初始化http server错误")
			break

		}
	}

	return ret
}
