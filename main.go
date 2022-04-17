package main

import (
	"fmt"
	"log"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

var (
	AppName string
)

func main() {
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	if err := bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/gopher.png",
		},
		Debug: true,
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w := ws[0]
			w.OpenDevTools()
			return nil
		},
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astikit.StrPtr("#fff"),
				Center:          astikit.BoolPtr(false),
				Height:          astikit.IntPtr(720),
				Width:           astikit.IntPtr(1280),
			},
		}},
	}); err != nil {
		l.Fatal(fmt.Errorf("bootstrap failed: %w", err))
	}
}
