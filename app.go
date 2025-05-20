package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}

// Seleciona a pasta de entrada
func (a *App) OpenFolderDialog() (string, error) {
	if a.ctx == nil {
		return "", fmt.Errorf("context is nil")
	}
	return runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Selecione a pasta com arquivos Excel",
	})
}

// Seleciona a pasta de saída
func (a *App) SelectOutputFolder() (string, error) {
	if a.ctx == nil {
		return "", fmt.Errorf("context is nil")
	}
	return runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Selecione a pasta de saída para os PDFs",
	})
}

// Chama o processo de conversão
func (a *App) ProcessExcelFiles(inputFolder, outputFolder string) error {
	return ProcessExcelFiles(inputFolder, outputFolder)
}
