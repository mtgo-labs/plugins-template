package YOUR_PLUGIN_NAME_test

import (
	"context"
	"testing"

	tg "github.com/mtgo-labs/mtgo/telegram"
	"YOUR_PLUGIN_NAME"
)

func TestPluginName(t *testing.T) {
	p := YOUR_PLUGIN_NAME.New(YOUR_PLUGIN_NAME.Config{})
	if p.Name() != "YOUR_PLUGIN_NAME" {
		t.Fatalf("expected YOUR_PLUGIN_NAME, got %s", p.Name())
	}
}

func TestPluginStartStop(t *testing.T) {
	p := YOUR_PLUGIN_NAME.New(YOUR_PLUGIN_NAME.Config{})
	if err := p.Start(context.Background(), &tg.Client{}); err != nil {
		t.Fatal(err)
	}
	if err := p.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
}
