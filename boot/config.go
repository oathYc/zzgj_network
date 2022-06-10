package boot

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"zzgj_network/library/global"
)

func InitConfig() error {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev"
	}
	_, err := toml.DecodeFile(fmt.Sprintf("config/%s.toml", appEnv), &global.Config)
	if err != nil {
		return err
	}

	return nil
}
