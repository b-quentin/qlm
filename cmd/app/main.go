package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func checkConfigFile() error {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return fmt.Errorf("Failed to get user home directory: %s\n", err)
	}

	filePath := filepath.Join(homeDir, ".config", "QLM", "config.yaml")

	if !fileExists(filePath) {
		return fmt.Errorf("The file %s does not exist.\n", filePath)
	}

	return nil
}

func downloadFile(url string) error {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return fmt.Errorf("Failed to get user home directory: %s\n", err)
	}

	filePath := filepath.Join(homeDir, ".config", "QLM")

	fmt.Println("Mkdir filePath: %s", filePath)
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("Failed to create directory: %s", err)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Failed to create file: %s", err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to download file: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download file: %s", resp.Status)
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to save file: %s", err)
	}

	return nil
}

func main() {
	if err := checkConfigFile(); err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Download...")

		if err := downloadFile("https://github.com/b-quentin/qlm/tree/master/assets/.config/QLM/config.yaml"); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
