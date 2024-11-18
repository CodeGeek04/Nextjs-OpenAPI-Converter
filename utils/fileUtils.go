// utils/fileUtils.go
package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/types"
)

// WriteToFile writes data to a file, creating directories if needed
func WriteToFile(data []byte, filename string) error {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0755)
	CheckError(err)

	err = os.WriteFile(filename, data, 0644)
	CheckError(err)

	return nil
}

// ListDirectory prints the contents of a directory, distinguishing between files and folders
func ListDirectory(path string) ([]types.DirectoryItem, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	items := make([]types.DirectoryItem, 0, len(entries))

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, fmt.Errorf("error getting file info for %s: %w", entry.Name(), err)
		}

		itemType := types.FileType
		if entry.IsDir() {
			itemType = types.FolderType
		}

		// Check if file is hidden (works on Unix-like systems)
		isHidden := false
		if entry.Name()[0] == '.' {
			isHidden = true
		}

		item := types.DirectoryItem{
			Name:        entry.Name(),
			Path:        filepath.Join(path, entry.Name()),
			Type:        itemType,
			Size:        info.Size(),
			ModTime:     info.ModTime(),
			IsHidden:    isHidden,
			Permissions: info.Mode(),
		}

		items = append(items, item)
	}

	return items, nil
}

// ListDirectoryRecursive prints the contents of a directory and its subdirectories
func ListDirectoryRecursive(path string, indent string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	for _, entry := range entries {
		// Get entry info
		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("error getting file info: %w", err)
		}

		// Format size string
		size := ""
		if !entry.IsDir() {
			size = fmt.Sprintf(" (%d bytes)", info.Size())
		}

		// Print entry with appropriate prefix and size
		prefix := "📄 " // File prefix
		if entry.IsDir() {
			prefix = "📁 " // Directory prefix
		}
		fmt.Printf("%s%s%s%s\n", indent, prefix, entry.Name(), size)

		// Recursively list subdirectories
		if entry.IsDir() {
			err := ListDirectoryRecursive(filepath.Join(path, entry.Name()), indent+"  ")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ReadFileContents reads and returns the contents of a file as a string
func ReadFileContents(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return string(data), nil
}

// ReadFileLines reads a file and returns its contents as a slice of lines
func ReadFileLines(filename string) ([]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var lines []rune
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	for _, line := range string(data) {
		lines = append(lines, line)
	}

	return lines, nil
}
