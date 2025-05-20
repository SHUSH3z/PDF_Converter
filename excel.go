package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

// ProcessExcelFiles percorre a pasta de entrada e converte cada arquivo Excel em PDF
func ProcessExcelFiles(inputFolder, outputFolder string) error {
	fmt.Println("[ProcessExcelFiles] Início")
	fmt.Println("  inputFolder:", inputFolder)
	fmt.Println("  outputFolder:", outputFolder)

	files, err := os.ReadDir(inputFolder)
	if err != nil {
		return fmt.Errorf("erro ao ler pasta de entrada: %w", err)
	}

	for _, file := range files {
		name := file.Name()
		if filepath.Ext(name) == ".xlsx" || filepath.Ext(name) == ".xls" {
			inputPath := filepath.Join(inputFolder, name)
			outputPath := filepath.Join(outputFolder, name[:len(name)-len(filepath.Ext(name))]+".pdf")

			fmt.Println("[ProcessExcelFiles] Convertendo:", inputPath, "→", outputPath)

			err := convertExcelToPDF(inputPath, outputPath)
			if err != nil {
				return fmt.Errorf("erro ao converter %s: %w", name, err)
			}
		}
	}
	return nil
}

// convertExcelToPDF usa PowerShell + COM para converter um arquivo Excel em PDF
// Esta função foi movida para uma função global (não depende de App)
func convertExcelToPDF(inputPath, outputPath string) error {
	script := fmt.Sprintf(`
		$excel = New-Object -ComObject Excel.Application
		$excel.Visible = $false
		$workbook = $excel.Workbooks.Open("%s")
		$workbook.ExportAsFixedFormat(0, "%s")
		$workbook.Close($false)
		$excel.Quit()
	`, inputPath, outputPath)

	// Caminho do PowerShell
	cmdPath := "powershell.exe"

	// Cria o comando
	cmdInstance := exec.Command(cmdPath, "-WindowStyle", "Hidden", "-Command", script)

	// Configura o atributo para esconder a janela
	cmdInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	// Executa o comando
	cmdOutput, err := cmdInstance.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao chamar powershell: %w | output: %s", err, string(cmdOutput))
	}

	return nil
}
