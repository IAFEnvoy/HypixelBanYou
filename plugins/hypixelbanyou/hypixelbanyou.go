package hypixelbanyou

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/robinbraemer/event"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
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

func onConnect() func(e *proxy.LoginEvent) {
	return func(e *proxy.LoginEvent) {
		println(e.Player().VirtualHost().String())
		e.Deny(&component.Text{
			"Test",
			component.Style{},
			nil,
		})
	}
}
