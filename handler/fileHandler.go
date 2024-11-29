package handler

import (
	"errors"
	"strings"

	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/types"
	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/utils"
)

func generateRequest(requestType string, relativePath string) (types.Request, error) {
	switch requestType {
	case "GET":
		return types.Request{
			Method: "GET",
			Header: []types.Header{},
			URL: types.URL{
				Raw:      "http://{{eloquens_url}}:3000/" + relativePath,
				Protocol: "",
				Host:     []string{"{{eloquens_url}}"},
				Path:     strings.Split(relativePath, "/"),
				Query:    []types.Query{},
			},
		}, nil
	case "POST":
		return types.Request{
			Method: "POST",
			Header: []types.Header{},
			URL: types.URL{
				Raw:      "http://{{eloquens_url}}:3000/" + relativePath,
				Protocol: "",
				Host:     []string{"{{eloquens_url}}"},
				Port:     "",
				Path:     strings.Split(relativePath, "/"),
				Query:    []types.Query{},
			},
		}, nil
	case "PUT":
		return types.Request{
			Method: "PUT",
			Header: []types.Header{},
			URL: types.URL{
				Raw:      "http://{{eloquens_url}}:3000/" + relativePath,
				Protocol: "",
				Host:     []string{"{{eloquens_url}}"},
				Port:     "",
				Path:     strings.Split(relativePath, "/"),
				Query:    []types.Query{},
			},
		}, nil
	case "DELETE":
		return types.Request{
			Method: "DELETE",
			Header: []types.Header{},
			URL: types.URL{
				Raw:      "http://{{eloquens_url}}:3000/" + relativePath,
				Protocol: "",
				Host:     []string{"{{eloquens_url}}"},
				Port:     "",
				Path:     strings.Split(relativePath, "/"),
				Query:    []types.Query{},
			},
		}, nil
	case "PATCH":
		return types.Request{
			Method: "PATCH",
			Header: []types.Header{},
			URL: types.URL{
				Raw:      "http://{{eloquens_url}}:3000/" + relativePath,
				Protocol: "",
				Host:     []string{"{{eloquens_url}}"},
				Port:     "",
				Path:     strings.Split(relativePath, "/"),
				Query:    []types.Query{},
			},
		}, nil
	}

	return types.Request{}, errors.New("Invalid request type")
}

func checkWordInFile(fileContent string, word string) bool {
	return strings.Contains(fileContent, word)
}

func HandleFile(path string, basePath string) ([]types.Request, error) {
	// Check if path is a file or directory
	fileName := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	if !strings.Contains(fileName, "route") {
		return []types.Request{}, nil
	}

	path = strings.Replace(path, "\\", "/", -1)
	relativePath := strings.Replace(path, basePath, "", 1)

	fileContent, err := utils.ReadFileContents(path)
	utils.CheckError(err)

	var requests []types.Request

	words := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

	for _, word := range words {
		if checkWordInFile(fileContent, word) {
			request, err := generateRequest(word, relativePath)
			utils.CheckError(err)

			requests = append(requests, request)
		}
	}

	return requests, nil
}
