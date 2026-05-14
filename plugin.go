package YOUR_PLUGIN_NAME

import (
	"context"

	tg "github.com/mtgo-labs/mtgo/telegram"
)

type Config struct {
}

type Plugin struct {
	config Config
	client *tg.Client
}

func New(config Config) *Plugin {
	return &Plugin{config: config}
}

func (p *Plugin) Name() string {
	return "YOUR_PLUGIN_NAME"
}

func (p *Plugin) Start(ctx context.Context, client *tg.Client) error {
	p.client = client
	return nil
}

func (p *Plugin) Stop(ctx context.Context) error {
	return nil
}
