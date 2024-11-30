package extraction

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// Baixar arquivo WinRAR e criar a pasta
func DownloadRARFile(url, destFolder, destFileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error accessing the resource: %v", err)
	}
	defer resp.Body.Close()

	err = os.MkdirAll(destFolder, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating folder: %v", err)
	}

	rarFile := fmt.Sprintf("%s/%s", destFolder, destFileName)
	out, err := os.Create(rarFile)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return nil
}

// Extrair arquivos WinRAR
func ExtractRAR(src, destination string) error {
	err := os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return err
	}

	cmd := exec.Command("C:\\Program Files\\UnRAR\\UnRAR.exe", "x", "-y", src, destination)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to extract RAR file: %w", err)
	}

	return nil
}
