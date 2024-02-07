package hypixelbanyou

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/robinbraemer/event"
	component "go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"strings"
)

// Plugin is a ping plugin that handles ping events.
var Plugin = proxy.Plugin{
	Name: "HypixelBanYou",
	Init: func(ctx context.Context, p *proxy.Proxy) error {
		log := logr.FromContextOrDiscard(ctx)
		log.Info("Hypixel Ban You plugin loaded")

		event.Subscribe(p.Event(), 0, onConnect())

		return nil
	},
}

var cuteNameList = map[string]string{
	"node1": "cheating",
	"node2": "boosting",
	"node3": "ipban",
	"node4": "chat",
	"node5": "nmsl",
	"node6": "skyblock",
}

func onConnect() func(e *proxy.LoginEvent) {
	return func(e *proxy.LoginEvent) {
		domain := e.Player().VirtualHost().String()
		t := strings.Split(domain, ".")[0]
		data, exists := cuteNameList[t]
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
