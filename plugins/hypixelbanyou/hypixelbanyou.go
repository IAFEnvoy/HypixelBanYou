package hypixelbanyou

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/robinbraemer/event"
	"github.com/spf13/viper"
	component "go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"strings"
)

type ReasonConfig struct {
	mode     string
	message  string
	time     string
	hasBanId string
	extra    string
}

type BanConfig struct {
	routes  map[string]string
	reasons map[string]ReasonConfig
}

var banConfig BanConfig

// Plugin is a ping plugin that handles ping events.
var Plugin = proxy.Plugin{
	Name: "HypixelBanYou",
	Init: func(ctx context.Context, p *proxy.Proxy) error {
		log := logr.FromContextOrDiscard(ctx)
		log.Info("Hypixel Ban You plugin loaded")
		//Load Reasons
		viper.SetConfigType("yaml")
		viper.SetConfigFile("./reason.yml")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		} else {
			resultMap := make(map[string]map[string]string)
			err := viper.Unmarshal(&resultMap)
			if err != nil {
				panic(err)
			} else {
				banConfig.reasons = make(map[string]ReasonConfig)
				for key, value := range resultMap {
					banConfig.reasons[key] = ReasonConfig{
						mode:     value["mode"],
						message:  value["message"],
						time:     value["time"],
						hasBanId: value["has_ban_id"],
						extra:    value["extra"],
					}
				}
				log.Info("reason.yml loaded")
			}
		}
		//Load Routes
		viper.SetConfigType("yaml")
		viper.SetConfigFile("./route.yml")
		err1 := viper.ReadInConfig()
		if err1 != nil {
			panic(err1)
		} else {
			err := viper.Unmarshal(&banConfig.routes)
			if err != nil {
				panic(err)
			} else {
				log.Info("route.yml loaded")
			}
		}
		event.Subscribe(p.Event(), 0, onConnect())
		return nil
	},
}

func onConnect() func(e *proxy.LoginEvent) {
	return func(e *proxy.LoginEvent) {
		domain := e.Player().VirtualHost().String()
		t := strings.Split(domain, ".")[0]
		data, exists := banConfig.routes[t]
		if exists {
			t = data
		}
		content, err := GetBanMessage(e.Player().Username(), t)
		if err != nil {
			content = "§c你的访问已被拦截\n\n§7原因：§f无法根据访问的域名解析节点\n\n§b请更换节点后再试"
		}
		e.Deny(&component.Text{
			Content: content,
		})
	}
}
