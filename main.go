package main

import (
	"encoding/json"

	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/handler"
	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/types"
	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/utils"
)

func generatePostmanCollection(item types.Item) types.PostmanCollection {
	postmanCollection := types.PostmanCollection{
		Info: types.PostmanInfo{
			Name:   "Nextjs to API Collection",
			Schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		},
		Item: []types.Item{item},
	}
	return postmanCollection
}

func indentAndWriteToFile(postmanCollection types.PostmanCollection, fileName string) {
	indentedMarshal, err := json.MarshalIndent(postmanCollection, "", "    ")
	utils.CheckError(err)

	err = utils.WriteToFile(indentedMarshal, fileName)
	utils.CheckError(err)
}

func main() {
	item := handler.HandleFolder("NextJS API Path", "NextJS Root Path")
	postmanCollection := generatePostmanCollection(item)
	indentAndWriteToFile(postmanCollection, "postman_collection.json")
}
