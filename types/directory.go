package types

import (
	"fmt"
	"os"
	"time"
)

// ItemType represents the type of a directory item
type ItemType int

const (
	FileType ItemType = iota
	FolderType
)

// DirectoryItem represents a file or folder in a directory
type DirectoryItem struct {
	Name        string
	Path        string
	Type        ItemType
	Size        int64
	ModTime     time.Time
	IsHidden    bool
	Permissions os.FileMode
}

// String provides a formatted string representation of DirectoryItem
func (di DirectoryItem) String() string {
	prefix := "📄"
	if di.Type == FolderType {
		prefix = "📁"
	}

	size := ""
	if di.Type == FileType {
		size = fmt.Sprintf(" (%d bytes)", di.Size)
	}

	return fmt.Sprintf("%s %s%s", prefix, di.Name, size)
}

func (di DirectoryItem) IsFile() bool {
	return di.Type == FileType
}

// IsFolder returns true if the item is a folder
func (di DirectoryItem) IsFolder() bool {
	return di.Type == FolderType
}
