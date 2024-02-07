package ping

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	. "github.com/minekube/gate-plugin-template/util"
	"github.com/minekube/gate-plugin-template/util/mini"
	"github.com/robinbraemer/event"
	"go.minekube.com/common/minecraft/color"
	c "go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

// Plugin is a ping plugin that handles ping events.
var Plugin = proxy.Plugin{
	Name: "Ping",
	Init: func(ctx context.Context, p *proxy.Proxy) error {
		log := logr.FromContextOrDiscard(ctx)
		log.Info("Ping plugin loaded")

		event.Subscribe(p.Event(), 0, onPing())

		return nil
	},
}

func onPing() func(*proxy.PingEvent) {
	line1 := mini.Gradient(
		"Speed IP for Hypixel for free!\n",
		c.Style{Bold: c.True},
		*color.Red.RGB, *color.Gold.RGB, *color.Yellow.RGB, *color.Green.RGB, *color.Blue.RGB, *color.DarkPurple.RGB,
	)
	line2 := mini.Gradient(
		fmt.Sprintf("Powered By Ficer Studio and Rick"),
		c.Style{},
		*color.White.RGB, *color.Black.RGB,
	)

	return func(e *proxy.PingEvent) {
		p := e.Ping()
		p.Description = Join(line1, line2)
		p.Players.Max = p.Players.Online + 1
	}
}
