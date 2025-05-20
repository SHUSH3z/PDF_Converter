package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:            "Excel → PDF",
		Width:            800,
		Height:           600,
		BackgroundColour: &options.RGBA{R: 240, G: 242, B: 245, A: 255}, // mesmo fundo do frontend
		AssetServer:      &assetserver.Options{Assets: assets},
		OnStartup:        app.OnStartup,
		Bind:             []interface{}{app},
		Windows: &windows.Options{
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			WebviewIsTransparent: true,
		},
	})
	if err != nil {
		println("Erro ao iniciar:", err.Error())
	}
}
