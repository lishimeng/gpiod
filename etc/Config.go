package etc

import (
	"fmt"
	"github.com/lishimeng/etc"
)

var Config Configuration

var ConfigFile = "gpiod.toml"

var env = []string{ ".", "/etc/gpiod", }

func Load() {
	if err := etc.LoadEnvs(ConfigFile, env, &Config); err != nil {
		fmt.Println(err)
	}
}