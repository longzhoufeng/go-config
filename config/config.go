package config

import (
	"fmt"
	"github.com/longzhoufeng/go-config/pkg"
	"github.com/longzhoufeng/go-sdk/pkg"
	"log"

	"github.com/longzhoufeng/go-config/pkg/source"
)

var (
	ExtendConfig interface{}
	_cfg         *Settings
)

// Settings 兼容原先的配置结构
type Settings struct {
	Settings  Config `yaml:"settings"`
	callbacks []func()
}

func (e *Settings) runCallback() {
	for i := range e.callbacks {
		e.callbacks[i]()
	}
}

func (e *Settings) OnChange() {
	e.init()
	log.Println("!!! config change and reload")
}

func (e *Settings) Init() {
	e.init()
	log.Println(pkg.Green("!!! config init"))
}

func (e *Settings) init() {
	e.runCallback()
}

// Config 配置集合
type Config struct {
	Application *Application          `yaml:"application"`
}


// Setup 载入配置文件
func Setup(s source.Source,
	fs ...func()) {
	_cfg = &Settings{
		Settings: Config{
			Application: ApplicationConfig,
		},
		callbacks: fs,
	}
	var err error
	config.DefaultConfig, err = config.NewConfig(
		config.WithSource(s),
		config.WithEntity(_cfg),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("New config object fail: %s", err.Error()))
	}
	_cfg.Init()
}
