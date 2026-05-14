# mtgo plugin: YOUR_PLUGIN_NAME

Short description of what this plugin does.

## Install

```bash
go get github.com/mtgo-labs/plugins/YOUR_PLUGIN_NAME
```

## Getting Started

1. Copy this template to `plugins/your_plugin_name/`
2. Replace `YOUR_PLUGIN_NAME` in `go.mod`, `plugin.go`, `plugin_test.go`, and `README.md` with your plugin name (lowercase, no hyphens/underscores)
3. Update `go.mod` module path to `github.com/mtgo-labs/plugins/your_plugin_name`
4. Implement the `Plugin` interface: `Name()`, `Start()`, `Stop()`

## Usage

```go
import (
    tg "github.com/mtgo-labs/mtgo/telegram"
    yourplugin "github.com/mtgo-labs/plugins/YOUR_PLUGIN_NAME"
)

func main() {
    client, _ := tg.NewClient(apiID, apiHash, &tg.Config{
        BotToken: botToken,
    })

    client.Use(yourplugin.New(yourplugin.Config{}))

    // your handlers...
    client.Start()
}
```

## Configuration

| Field | Type | Description |
|-------|------|-------------|
|       |      |             |

## License

MIT
