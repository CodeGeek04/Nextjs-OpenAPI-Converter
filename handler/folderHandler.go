package handler

import (
	"strings"

	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/types"
	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/utils"
)

func generateRequestItem(request types.Request, name string) types.Item {
	return types.Item{
		Name:     name,
		Request:  &request,
		Response: []string{},
	}
}

func generateFolderItem(name string, items []types.Item) types.Item {
	return types.Item{
		Name: name,
		Item: items,
	}
}

func HandleFolder(path string, basePath string) types.Item {
	var requests []types.Request

	directoryContents, err := utils.ListDirectory(path)
	utils.CheckError(err)

	path = strings.Replace(path, "\\", "/", -1)
	// relativePath := strings.Replace(path, basePath, "", 1)
	folderName := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]

	folderItem := generateFolderItem(folderName, []types.Item{})

	for _, directoryItem := range directoryContents {
		if directoryItem.IsFolder() {
			folderItem.Item = append(folderItem.Item, HandleFolder(directoryItem.Path, basePath))
		} else {
			requests, err = HandleFile(directoryItem.Path, basePath)
			utils.CheckError(err)

			if len(requests) <= 0 {
				continue
			}

			for _, request := range requests {
				requestItem := generateRequestItem(request, directoryItem.Name)
				folderItem.Item = append(folderItem.Item, requestItem)
			}
		}
	}

	return folderItem
}
